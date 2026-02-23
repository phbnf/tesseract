## 2026-02-23 - Prevent DoS via Request Body Size Limit
**Vulnerability:** Unbounded `io.ReadAll` on request body in `add-chain` handlers.
**Learning:** `http.MaxBytesReader` is the standard way to limit request body size in Go, but it returns a `*http.MaxBytesError` that must be explicitly handled to return a 413 status code. Also, `http.MaxBytesReader` modifies the `ResponseWriter` to close the connection, which is a good mitigation for DoS.
**Prevention:** Always wrap `r.Body` with `http.MaxBytesReader` when reading potentially large payloads, and ensure the error is handled.
