## 2026-02-06 - Denial of Service via Large Request Body
**Vulnerability:** The `add-chain` endpoint read the entire request body into memory using `io.ReadAll` without any size limit. This allowed an attacker to send a very large request body, potentially causing an Out of Memory (OOM) crash and Denial of Service (DoS).
**Learning:** `io.ReadAll` is convenient but dangerous for untrusted input streams. In HTTP handlers, it blindly consumes the `Body` until EOF.
**Prevention:** Always enforce a maximum request size using `http.MaxBytesReader` before reading the body. This provides a hard limit on memory consumption per request.
