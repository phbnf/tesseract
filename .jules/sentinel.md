## 2026-02-20 - Unbounded Request Body
**Vulnerability:** The `add-chain` and `add-pre-chain` endpoints used `io.ReadAll` on the request body without a limit, allowing for potential DoS via memory exhaustion.
**Learning:** `http.MaxBytesReader` must be explicitly used to wrap request bodies when reading them fully, especially for public endpoints.
**Prevention:** Enforce a `MaxBodySize` configuration in `NewLogHandler` and apply it using `http.MaxBytesReader` in all handlers that read the body.
