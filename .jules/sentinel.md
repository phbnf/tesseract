## 2026-02-14 - Request Body Size Limit
**Vulnerability:** The `add-chain` and `add-pre-chain` endpoints were vulnerable to Denial of Service (DoS) attacks because they read the entire request body into memory using `io.ReadAll` without a size limit.
**Learning:** `io.ReadAll` on `r.Body` is dangerous for public endpoints. The `http.Server`'s `MaxHeaderBytes` or `ReadHeaderTimeout` does not protect against large bodies.
**Prevention:** Use `http.MaxBytesReader` to wrap the request body before reading it, or enforce a limit at the infrastructure level (e.g., reverse proxy). In Go handlers, explicitly limit the reader.
