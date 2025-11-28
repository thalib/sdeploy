## Overview
- Implement configuration loading and validation for SDeploy
- Support JSON config file with global and per-project settings
- Handle config file search order: `-c` flag, `/etc/sdeploy/config.json`, `./config.json`

## Requirements
- Define Go structs: `Config`, `EmailConfig`, `ProjectConfig`
- Parse JSON config with validation for required fields
- Support config search order per SPEC.md
- Validate: `listen_port`, `webhook_path` uniqueness, `webhook_secret` presence
- Return descriptive errors for invalid/missing config

## Acceptance
- TDD: Write `config_test.go` tests BEFORE implementation
- Tests cover: valid config loading, missing file, invalid JSON, missing required fields
- `go test ./cmd/sdeploy/... -run TestConfig` passes
- Code location: `cmd/sdeploy/config.go`, `cmd/sdeploy/config_test.go`
