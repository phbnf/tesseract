## 2026-01-21 - Hot Path Map Reallocation
**Learning:** `chainValidator.validate` was recreating maps for OID and EKU checks on every request. This is a common pattern where validation configuration is passed as slices but efficient lookup requires maps.
**Action:** Look for TODOs mentioning "pre-calc" or "builder pattern" in validation logic. Check for map creation inside frequently called methods (like `validate`) that rely on static configuration.
