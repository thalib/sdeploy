## Overview
- Implement comprehensive logging for SDeploy daemon
- Support stdout/stderr output and file-based logging
- Log deployment start/end, output, status, and errors

## Requirements
- Create logger with configurable output (stdout or file via `log_filepath`)
- Log levels: INFO, WARN, ERROR with timestamps
- Thread-safe logging for concurrent deployments
- Log format: `[TIMESTAMP] [LEVEL] [PROJECT] message`
- Rotate or append to log file based on configuration

## Acceptance
- TDD: Write `logging_test.go` tests BEFORE implementation
- Tests cover: stdout logging, file logging, log format, thread safety
- `go test ./cmd/sdeploy/... -run TestLogging` passes
- Code location: `cmd/sdeploy/logging.go`, `cmd/sdeploy/logging_test.go`
