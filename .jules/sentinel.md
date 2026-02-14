## 2025-02-05 - [Unbounded Request Body]
**Vulnerability:** The `add-chain` endpoint allowed unbounded request bodies, leading to potential DoS via memory exhaustion.
**Learning:** Go's `io.ReadAll` does not enforce limits by default. `http.MaxBytesReader` must be explicitly used.
**Prevention:** Always wrap `r.Body` with `http.MaxBytesReader` in HTTP handlers that read the full body.
