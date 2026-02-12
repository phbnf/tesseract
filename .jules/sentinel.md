# Sentinel Journal

## 2026-02-12 - Unlimited Request Body Vulnerability
**Vulnerability:** The `add-chain` and `add-pre-chain` endpoints used `io.ReadAll` on the request body without any size limit, allowing for potential DoS via memory exhaustion.
**Learning:** `io.ReadAll` in Go is dangerous for untrusted input. Handlers must enforce body size limits using `http.MaxBytesReader` or middleware.
**Prevention:** Wrap request bodies with `http.MaxBytesReader` before reading, or use `http.MaxBytesHandler` at the router level.
