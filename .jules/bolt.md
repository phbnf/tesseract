## 2024-05-23 - Value Receivers Limit Lazy Init
**Learning:** `chainValidator` uses value receivers for validation methods (`validate`). This means any state mutation (like lazy map initialization) inside the method is lost after return. This necessitates performing all initialization in the constructor (`NewChainValidator`) or accepting re-computation on every call.
**Action:** When optimizing struct methods, check receiver type first. If value receiver, move expensive initialization to constructor or change to pointer receiver (if safe).
