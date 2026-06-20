# Sentinel Journal

## 2025-02-18 - CWE-200 /debug/vars Exposure

**Vulnerability:**
The POSIX binary imports `_ "expvar" // Registers /debug/vars, with BadgerDB metrics.` which registers `expvar` handlers to `http.DefaultServeMux`. Similarly, the GCP binary has `_ "net/http/pprof"` exposing profiling data via `http.DefaultServeMux`. Since these servers bind their API endpoints to `http.DefaultServeMux` (e.g. `http.Handle("/", otelhttp.NewHandler(logHandler, "/"))`), this makes `/debug/vars` and `/debug/pprof/*` publicly accessible.

**Learning:**
Anonymous imports like `_ "expvar"` and `_ "net/http/pprof"` automatically register routes on `http.DefaultServeMux`. If an application uses `http.DefaultServeMux` (e.g., via `http.Handle` or `http.ListenAndServe(..., nil)` or when setting `http.Server.Handler` implicitly to `nil`) for its public API, these debug endpoints are inadvertently exposed. This constitutes a CWE-200 Information Exposure.

**Prevention:**
Remove the anonymous imports of `expvar` and `net/http/pprof`. If profiling or metrics are needed, they should be explicitly registered on a separate internal multiplexer or port that is not accessible from the public internet.
