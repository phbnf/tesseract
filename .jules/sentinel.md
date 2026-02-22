## 2025-05-15 - Unbounded Request Body in CT Handlers
**Vulnerability:** `io.ReadAll` was used on request bodies in `parseBodyAsJSONChain` without any size limit, allowing for potential DoS via memory exhaustion.
**Learning:** Even when using helper functions for parsing, always enforce limits on untrusted input streams. `http.MaxBytesReader` is the standard way to do this in Go.
**Prevention:** Ensure `HandlerOptions` or similar config structs always include a `MaxBodySize` and enforce it in all handlers that read bodies.
