## 2025-05-23 - [Missing Request Body Size Limit]
**Vulnerability:** The `add-chain` endpoint allowed unlimited request body size, leading to potential DoS via memory exhaustion.
**Learning:** Go's `http.MaxBytesReader` is the standard way to mitigate this, but it requires access to `http.ResponseWriter` which might not be passed down to helper functions.
**Prevention:** Always enforce a `MaxBodySize` (e.g., 4MB) on public endpoints accepting JSON payloads.
