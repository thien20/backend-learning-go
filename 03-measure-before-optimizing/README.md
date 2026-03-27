# 03 Measure Before Optimizing

## Goal

Learn to measure first and optimize second.

This stage intentionally includes one slow code path and one inefficient repository path. The point is not to guess at fixes. The point is to capture a baseline, profile the process, and then improve the right thing for the right reason.

## What Student Must Implement

- Real `/metrics` output through `internal/metrics`.
- Real pprof route wiring in `internal/handler/profiler.go`.
- A before/after measurement process in `benchmarks/baseline.md` and `benchmarks/after.md`.
- Fix the intentionally slow response-building path.
- Fix the intentionally inefficient repository path.
- Decide which optimizations belong in code versus storage design.

## Suggested Order

1. Run the server and execute `load/smoke.js`.
2. Capture a baseline in `benchmarks/baseline.md`.
3. Wire pprof and inspect CPU and heap profiles.
4. Investigate the slow response-building path in the service layer.
5. Investigate the N+1 style repository path.
6. Re-run the load scripts and document the delta in `benchmarks/after.md`.

## Definition Of Done

- You can explain the difference between throughput and concurrency for this API.
- You can show before/after measurements for at least one meaningful change.
- The obvious bottlenecks are removed or intentionally justified.
- `/metrics` and pprof are useful enough to guide the next round of work.
- You did not optimize blindly.

## Search/Read Topics

- p50/p95/p99
- throughput vs concurrency
- go tool pprof
- CPU profile vs heap profile
- N+1 query problem
- EXPLAIN ANALYZE
- connection pool basics
- benchmark before/after
