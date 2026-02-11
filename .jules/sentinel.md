# Sentinel Journal

## 2026-02-11 - DoS via Unbounded Request Body
**Vulnerability:** The `add-chain` endpoint in `internal/ct/handlers.go` used `io.ReadAll` on the request body without enforcing a size limit. This allows a malicious actor to send an excessively large payload, leading to memory exhaustion (Denial of Service).
**Learning:** In Go, `io.ReadAll` reads until EOF or error. Without a `MaxBytesReader`, it will consume all available memory if the client keeps sending data. The `add-chain` handler processes certificate chains which can be large, but not infinite.
**Prevention:** Always wrap `r.Body` with `http.MaxBytesReader` when reading the entire body into memory, especially for public endpoints. Enforce a reasonable limit (e.g., 4MB for CT logs). Use `errors.As` to check for `*http.MaxBytesError` and return a 413 Payload Too Large status code.
