# 01 Understand Request Flow

## Goal

Learn the basic backend request flow in the same shape as the main repo:

`handler -> service -> repository -> db -> cache`

This stage stays small on purpose. There is one route, one model, one in-memory database, and one in-memory cache. The job is not to memorize finished code. The job is to complete the missing pieces and understand why each layer exists.

## What Student Must Implement

- Tighten request validation in the handler and service.
- Finish the cache-aside behavior and decide how cache errors should behave.
- Improve repository error translation so the upper layers do not depend on storage details.
- Extend the in-memory DB and cache safely with `sync.RWMutex`.
- Decide what makes an `Item` valid for this stage.

## Suggested Order

1. Read `cmd/api/main.go` and `internal/app/container.go` to understand dependency injection.
2. Start at `internal/handler/item_handler.go` and trace one request end-to-end.
3. Finish the service TODOs before expanding the repository or DB.
4. Add a second seeded item and test cache hits/misses manually.
5. Refactor error mapping only after the flow feels clear.

## Definition Of Done

- `GET /items/{id}` returns a valid JSON item.
- Invalid IDs return a deliberate client error.
- Missing IDs return a deliberate not found error.
- Cache hits and DB hits can be distinguished while debugging.
- You can explain what each layer owns without looking at the code.

## Search/Read Topics

- dependency injection
- thin handler
- service validation
- repository pattern
- cache-aside
- sentinel errors
- sync.RWMutex
