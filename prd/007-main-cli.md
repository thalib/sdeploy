## Overview
- Implement main entry point and CLI argument parsing
- Support console and daemon execution modes
- Wire all modules together for complete application

## Requirements
- Parse CLI flags: `-c` (config path), `-d` (daemon mode)
- Console mode: run in foreground, output to stdout/stderr
- Daemon mode: run as background service, output to logger
- Load config, initialize logger, start webhook server
- Graceful shutdown on SIGINT/SIGTERM
- Binary name: `sdeploy`

## Acceptance
- TDD: Write `main_test.go` tests BEFORE implementation (if applicable)
- Tests cover: flag parsing, config loading integration, graceful shutdown
- `go build -o sdeploy ./cmd/sdeploy` produces working binary
- `./sdeploy -h` shows help, `./sdeploy` runs in console, `./sdeploy -d` runs as daemon
- Code location: `cmd/sdeploy/main.go`
