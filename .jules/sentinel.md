## 2025-02-09 - [DoS] Unbounded Request Body Read
**Vulnerability:** The `add-chain` and `add-pre-chain` endpoints read the entire request body into memory using `io.ReadAll` without a size limit.
**Learning:** `io.ReadAll` on `r.Body` is dangerous if the client sends a large payload. Go's `http.MaxBytesReader` must be used to enforce limits.
**Prevention:** Always wrap `r.Body` with `http.MaxBytesReader` in the handler or middleware before reading. In this case, `appHandler.ServeHTTP` was the appropriate place to apply this limit globally for the log endpoints.
