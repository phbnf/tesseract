# Sentinel Journal

This journal tracks critical security learnings and patterns specific to this codebase.

## 2026-02-13 - [MaxBytesHandler Pattern]
**Vulnerability:** Go's `http.MaxBytesHandler` wraps the `ResponseWriter` but `io.ReadAll` returns `*http.MaxBytesError` which needs to be caught. Also, `http.MaxBytesHandler` does not automatically return 413 if the handler writes to the response, it only prevents reading more. However, `http.MaxBytesReader` (which I used) is different.
**Learning:** When using `http.MaxBytesReader`, `io.ReadAll` returns an error. The handler must catch this error and return 413 explicitly. The `http.MaxBytesReader` documentation says it returns `*http.MaxBytesError` on read.
**Prevention:** Always wrap `r.Body` with `http.MaxBytesReader` and check for `*http.MaxBytesError` when reading.
