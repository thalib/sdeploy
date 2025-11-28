## Overview
- Initialize the Go project structure for SDeploy webhook deployment daemon
- Set up module, directory layout, and foundational files per SPEC.md
- Establish TDD infrastructure with test scaffolding

## Requirements
- Create Go module `sdeploy` with `go.mod`
- Create directory structure: `cmd/sdeploy/` with placeholder files
- Files: `main.go`, `webhook.go`, `deploy.go`, `config.go`, `email.go`, `logging.go`
- Corresponding test files: `*_test.go` for each module
- Use Go 1.21+ with standard library where possible

## Acceptance
- `go mod tidy` runs without errors
- `go build ./cmd/sdeploy` compiles successfully
- `go test ./cmd/sdeploy/...` runs (even with placeholder tests)
- Directory structure matches SPEC.md project layout
