## 2025-05-23 - Performance Optimization in chainValidator
**Learning:** Tests bypassing constructors by modifying private struct fields directly can block optimizations that rely on constructor initialization.
**Action:** When adding pre-calculated fields (like maps) to a struct, always include fallback logic in the hot path to handle cases where the struct was manually initialized (e.g., in tests), ensuring backward compatibility without extensive test refactoring.
