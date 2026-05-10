---
name: go-cli-binaries
description: Create or extend Go CLI binaries with sensible project layout, Viper-based configuration, and Bubble Tea for terminal UIs when needed.
---

# Go CLI Binaries

Use this skill when adding a new Go command-line binary or improving an existing CLI in this repository.

## Goals

- Keep CLIs small, testable, and composable.
- Prefer standard library dependencies when possible.
- Use Viper for multi-source configuration (flags, env vars, config files) when configuration complexity warrants it.
- Use Bubble Tea for interactive TUIs when a non-interactive CLI is not sufficient.

## Structure

- Place executable entrypoints under `cmd/<binary-name>/main.go`.
- Keep business logic out of `main`; place reusable logic in `pkg/` or `internal/`.
- Keep CLI parsing, wiring, and process exit handling in thin command-layer code.

## Configuration

- Prefer explicit configuration types rather than loose maps.
- Use Viper for loading configuration from:
  - command-line flags
  - environment variables
  - optional config files
- Validate configuration early and fail fast with clear errors.

## TUI Guidance

- Use Bubble Tea for interactive workflows (forms, lists, dashboards, multi-step flows).
- Keep Bubble Tea models focused and separate rendering from domain logic.
- For simple one-shot commands, do **not** add Bubble Tea unnecessarily.

## Validation Before PR Review

- Run:
  - `go tool golangci-lint run ./...`
  - `go vet ./...`
  - `go test -race ./...`
