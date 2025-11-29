//go:build !windows

package main

import (
	"context"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)

// setProcessGroup sets the command to run in its own process group (Unix only)
func setProcessGroup(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

// killProcessGroup kills the process group (Unix only)
func killProcessGroup(cmd *exec.Cmd) {
	if cmd.Process != nil {
		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}
}

// getShellPath returns the path to the shell executable (Unix implementation)
// It first tries to find "sh" in PATH, then falls back to common shell locations
func getShellPath() string {
	// Try to find sh in PATH first
	if shellPath, err := exec.LookPath("sh"); err == nil {
		return shellPath
	}

	// Fallback to common shell locations
	commonPaths := []string{"/bin/sh", "/usr/bin/sh", "/usr/local/bin/sh"}
	for _, path := range commonPaths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// Last resort: return "sh" and let the OS handle it
	return "sh"
}

// getShellArgs returns the shell arguments for executing a command (Unix implementation)
func getShellArgs() string {
	return "-c"
}

// buildCommand creates an exec.Cmd with user/group settings (Unix implementation)
// It runs the command as the specified user and group using sudo if we're root,
// or falls back to running as current user if user lookup fails
func buildCommand(ctx context.Context, command, runAsUser, runAsGroup string) *exec.Cmd {
	// Check if we're running as root
	currentUser, err := user.Current()
	if err != nil || currentUser.Uid != "0" {
		// Not running as root or can't determine user, run command directly
		return exec.CommandContext(ctx, getShellPath(), getShellArgs(), command)
	}

	// Running as root, attempt to run as specified user/group
	targetUser, err := user.Lookup(runAsUser)
	if err != nil {
		// User not found, run command directly as current user
		return exec.CommandContext(ctx, getShellPath(), getShellArgs(), command)
	}

	targetGroup, err := user.LookupGroup(runAsGroup)
	if err != nil {
		// Group not found, run command directly as current user
		return exec.CommandContext(ctx, getShellPath(), getShellArgs(), command)
	}

	uid, err := strconv.ParseUint(targetUser.Uid, 10, 32)
	if err != nil {
		return exec.CommandContext(ctx, getShellPath(), getShellArgs(), command)
	}

	gid, err := strconv.ParseUint(targetGroup.Gid, 10, 32)
	if err != nil {
		return exec.CommandContext(ctx, getShellPath(), getShellArgs(), command)
	}

	cmd := exec.CommandContext(ctx, getShellPath(), getShellArgs(), command)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Credential: &syscall.Credential{
			Uid: uint32(uid),
			Gid: uint32(gid),
		},
		Setpgid: true,
	}

	return cmd
}
