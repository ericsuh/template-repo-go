# Agent Instructions

## Dependencies

- Prefer the Go standard library first.
- Keep third-party dependencies to a minimum.
- When third-party dependencies are needed, prefer highly vetted and widely used packages, ideally maintained by organizations with strong security practices (for example Google, Amazon, or Microsoft), when feasible for the use case.

## Quality Gates Before PR Review

- Always run lint and test checks before requesting PR review.
- At minimum, run:
  - `go tool golangci-lint run ./...`
  - `go vet ./...`
  - `go test -race ./...`
