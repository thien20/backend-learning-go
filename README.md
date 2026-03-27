# Backend Learning Go

This repo is a staged Go backend learning path.

Each top-level folder is a standalone mini-repo with its own `go.mod`. Start with one folder, finish the work in that stage, then move to the next one. The stages are designed to teach backend thinking in order: request flow, CRUD API design, then measurement and optimization.

## Recommended Path

1. Start with `01-understand-request-flow`.
2. Move to `02-build-crud-api` after you can explain the handler -> service -> repository -> db -> cache flow.
3. Move to `03-measure-before-optimizing` only after the CRUD stage feels clear.

## Stage Guide

- `01-understand-request-flow`
  Learn the layering and trace one request end to end.
  Default server port: `:8080`
- `02-build-crud-api`
  Extend the same layout into a more realistic CRUD API shell.
  Default server port: `:8081`
- `03-measure-before-optimizing`
  Measure first, then improve the most obvious bottlenecks.
  Default server port: `:8082`

## How To Use

1. Pick a stage folder.
2. Read that folder's `README.md` first.
3. Run the tests for that stage:

```bash
go test ./...
```

4. Start the API from inside that stage folder:

```bash
go run ./cmd/api
```

5. Make changes only inside the current stage folder.

## Example Start

```bash
cd 01-understand-request-flow
go test ./...
go run ./cmd/api
```

## Notes For New Users

- Do not treat the repo root as one Go module. Each stage is independent.
- Do not skip straight to the last stage. The later folders assume you understand the earlier ones.
- If you are unsure what to implement next, open the stage `README.md` and follow its goal, suggested order, and definition of done.
