# 02 Build CRUD API

## Goal

Turn the basic layered scaffold into a more realistic HTTP API foundation without skipping straight to production complexity.

This stage keeps the same project shape but introduces CRUD handlers, DTO placeholders, middleware seams, metrics placeholders, and a shutdown path. The emphasis is still on finishing the missing parts yourself.

## What Student Must Implement

- `POST /items`, `PUT /items/{id}`, and `DELETE /items/{id}` service logic.
- Request validation and domain validation rules.
- Better request/response DTO mapping.
- Structured logging fields and consistent log messages.
- Request ID generation and propagation through middleware.
- Meaningful `/metrics` output.
- Error mapping from domain errors to HTTP status codes.

## Suggested Order

1. Read the DTOs in `internal/model/item_dto.go` and decide what belongs in the domain model versus transport shapes.
2. Finish service validation for create and update flows.
3. Implement create first, then update, then delete.
4. Improve middleware so request IDs and logging become useful during debugging.
5. Replace the `/metrics` placeholder last.

## Definition Of Done

- CRUD routes exist and return deliberate responses.
- Invalid input gets a client error instead of a generic 500.
- Logs include request IDs and useful fields.
- `/metrics` exposes at least one counter and one latency-related placeholder or real metric.
- The in-memory DB can later be swapped for a SQL-backed implementation without changing handlers.

## Search/Read Topics

- net/http middleware
- context.Context
- DTO vs domain model
- structured logging
- Prometheus counter vs histogram
- graceful shutdown
- database/sql
- clean error mapping
