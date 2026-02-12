## 2025-02-09 - [DoS] Unbounded Request Body Read
**Vulnerability:** The `add-chain` and `add-pre-chain` endpoints read the entire request body into memory using `io.ReadAll` without a size limit.
**Learning:** `io.ReadAll` on `r.Body` is dangerous if the client sends a large payload. Go's `http.MaxBytesReader` or `http.MaxBytesHandler` must be used to enforce limits.
**Prevention:** Use `http.MaxBytesHandler` at the router level (`mux.Handle`) to enforce a global limit (4MB) on all relevant endpoints. Ensure underlying handlers catch `*http.MaxBytesError` to return a 413 status code.
