---
name: go-libraries-packages
description: Create or extend Go libraries/packages with clear package boundaries, testability, and correct use of internal directories.
---

# Go Libraries and Packages

Use this skill when adding reusable Go packages, modules, or internal components.

## Goals

- Design clear APIs and package boundaries.
- Keep packages focused on one responsibility.
- Prefer minimal external dependencies.

## Layout and Boundaries

- Place public, reusable packages under `pkg/` when intended for external consumption.
- Place non-exported application internals under `internal/` to enforce import boundaries.
- Avoid cyclic dependencies by keeping lower-level packages independent from command-layer code.

## API and Package Design

- Keep exported APIs small and intentional.
- Prefer constructors and options only when they materially improve clarity.
- Return rich errors and preserve context for callers.
- Keep interfaces close to where they are consumed.

## Testing

- Add unit tests for package behavior and edge cases.
- Prefer table-driven tests where they improve readability.
- Keep tests deterministic and avoid hidden global state.

## Validation Before PR Review

- Run:
  - `go tool golangci-lint run ./...`
  - `go vet ./...`
  - `go test -race ./...`
