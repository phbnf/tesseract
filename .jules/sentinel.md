## 2025-05-23 - Unbounded Request Body Reading
**Vulnerability:** `io.ReadAll` was used to read the entire request body into memory without any size limit in `parseBodyAsJSONChain`.
**Learning:** In Go, `io.ReadAll` does not enforce any limit. This is a common DoS vector if used on untrusted input streams like HTTP request bodies.
**Prevention:** Always use `http.MaxBytesReader` (for HTTP handlers) or `io.LimitReader` before reading the body. Enforce a reasonable default limit (e.g., 4MB for certificate chains).
