## 2026-02-26 - Unbounded Request Body
**Vulnerability:** Several HTTP handlers used `io.ReadAll(r.Body)` without enforcing a maximum size, exposing the service to DoS attacks via memory exhaustion.
**Learning:** Default Go HTTP server does not enforce request body limits. Handlers must explicitly use `http.MaxBytesReader` or similar mechanisms.
**Prevention:** Always wrap `r.Body` with `http.MaxBytesReader(w, r.Body, maxBytes)` before reading. Added `MaxBodySize` to `HandlerOptions` to enforce this centrally.
