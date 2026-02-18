## 2025-05-23 - Prevent Unbounded Request Body Reads
**Vulnerability:** The `add-chain` endpoint read the entire request body into memory using `io.ReadAll` without a size limit, exposing the server to DoS attacks via memory exhaustion (OOM).
**Learning:** `http.Server` in Go does not enforce a default request body size limit. Middleware or explicit handlers must implement `http.MaxBytesReader`.
**Prevention:** Always wrap `r.Body` with `http.MaxBytesReader(w, r.Body, Limit)` before reading, especially in handlers accepting JSON payloads.
