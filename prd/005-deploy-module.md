## Overview
- Implement deployment execution with locking and git operations
- Enforce single-instance execution per project
- Support git pull, command execution, timeout, and environment injection

## Requirements
- Per-project mutex lock to enforce single execution
- Skip deployment if lock held, log "Skipped", return 202
- If `git_update=true`, run `git pull` in `git_path` before deployment
- Execute `execute_command` in `execute_path` directory
- Support `timeout_seconds` for command execution (optional)
- Inject environment variables: `SDEPLOY_PROJECT_NAME`, `SDEPLOY_TRIGGER_SOURCE`, `SDEPLOY_GIT_BRANCH`
- Capture stdout/stderr, log output, handle errors gracefully
- Release lock after completion (success or failure)

## Acceptance
- TDD: Write `deploy_test.go` tests BEFORE implementation
- Tests cover: lock acquisition, skip on busy, git pull, command execution, timeout, env vars
- `go test ./cmd/sdeploy/... -run TestDeploy` passes
- Code location: `cmd/sdeploy/deploy.go`, `cmd/sdeploy/deploy_test.go`
