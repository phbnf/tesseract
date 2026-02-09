## 2025-02-09 - [DoS] Unbounded Request Body Read
**Vulnerability:** The `add-chain` and `add-pre-chain` endpoints read the entire request body into memory using `io.ReadAll` without a size limit.
**Learning:** `io.ReadAll` on `r.Body` is dangerous if the client sends a large payload. Go's `http.MaxBytesReader` must be used to enforce limits.
**Prevention:** Always wrap `r.Body` with `http.MaxBytesReader` when reading the body in handlers, especially for public endpoints.
