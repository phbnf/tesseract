## 2025-01-01 - Initializing Sentinel Journal

## 2025-02-09 - Fix Context Cancellation Leaks (CWE-400)
**Vulnerability:** Found multiple context cancellation leaks (CWE-400, gosec G118) where `context.WithCancel` was used without an immediate `defer cancel()` to ensure resources are released upon function exit. This occurred in `cmd/fsck/main.go`, `internal/hammer/hammer.go`, and worker goroutines in `internal/hammer/loadtest/workers.go`.
**Learning:** Even if `cancel()` is explicitly called later or passed to another struct, returning early without a `defer cancel()` leads to memory leaks where context resources (and associated goroutines) aren't properly cleaned up.
**Prevention:** Always explicitly call `defer cancel()` immediately after creating a context with `context.WithCancel` to ensure resources are freed, even if an early return happens or if explicit cancellation handles the normal execution flow.
