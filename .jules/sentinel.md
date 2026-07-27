## 2026-03-23 - Inadvertent Exposure of Profiling Endpoints via pprof Import

**Vulnerability:**
The `net/http/pprof` package was imported for side-effects (`_ "net/http/pprof"`) in the main entrypoints (`cmd/tesseract/posix/main.go` and `cmd/tesseract/gcp/main.go`). This automatically registers profiling endpoints (`/debug/pprof/*`) on `http.DefaultServeMux`, leading to a CWE-200 vulnerability (exposure of sensitive system information like memory profiles, CPU profiles, and goroutine stacks).

**Learning:**
Importing `net/http/pprof` for its side-effects is dangerous in binaries that serve public traffic on or potentially fallback to `http.DefaultServeMux`. Even if the primary application handler doesn't explicitly use `DefaultServeMux`, inadvertent use or misconfiguration elsewhere can inadvertently expose these endpoints. Profiling should be handled intentionally, ideally on a separate internal port, or restricted via authentication and network policies.

**Prevention:**
- Never use the side-effect import `_ "net/http/pprof"` in production binaries.
- If profiling is required, manually register pprof handlers on a dedicated, non-public multiplexer that is only accessible internally or requires authentication.
- Utilize security scanning tools like `gosec` (which identifies this issue as G108) during CI to catch inadvertent pprof imports before they reach production.
