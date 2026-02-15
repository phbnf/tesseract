## 2026-02-15 - Unbounded Request Body Reading in Helpers
**Vulnerability:** `parseBodyAsJSONChain` helper function used `io.ReadAll` on the request body without any size limit, exposing the `add-chain` endpoint to DoS attacks.
**Learning:** Helper functions abstracting request parsing often miss context (like `ResponseWriter` needed for `MaxBytesReader`) or configuration (limits), leading to default insecure behaviors (reading until EOF).
**Prevention:** Ensure all request body reading, even in helpers, enforces a maximum size limit, passing necessary dependencies (like `http.ResponseWriter`) down the call stack.
