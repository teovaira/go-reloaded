# Sprint 1: Core Transformations

**Goal:** Implement the 5 core transformation commands: (hex), (bin), (up), (low), (cap)  
**Duration:** 3-4 days  
**Tasks:** 7  
**Deliverable:** Basic transformations working, Golden Test Case 1 passes

---

## 🎯 Sprint Objectives

By end of Sprint 1:
- ✅ Tokenization system implemented
- ✅ Hexadecimal conversion working
- ✅ Binary conversion working
- ✅ Three case transformations working
- ✅ Code follows DRY principle

---

## 🧩 TASK-005: Tokenization

**Story Points:** 3/5 (Moderate)  
**Time:** 2-3 hours  
**Prerequisites:** TASK-004

### Learning Objectives
- String splitting strategies
- Token-based text processing
- Slice operations in Go

### What to Build
Two functions:
- `tokenize()`: Split text into words/tokens
- `detokenize()`: Join tokens back into text

### Test Scenarios
**Tokenize:**
- Input: `"hello world"` → Output: `["hello", "world"]`
- Input: `"hello    world"` (multiple spaces) → Output: `["hello", "world"]`
- Input: `"word (hex) test"` → Output: `["word", "(hex)", "test"]`
- Input: `""` → Output: `[]` (empty slice)

**Detokenize:**
- Input: `["hello", "world"]` → Output: `"hello world"`
- Input: `[]` → Output: `""`

### Architecture Decision
**Question:** How to handle punctuation?  
**Options:**
1. Keep punctuation attached to words: `"hello,"` stays as one token
2. Separate punctuation: `"hello,"` becomes `["hello", ","]`

**Recommendation:** Option 1 (keep attached) - simpler for now

### Edge Cases
- Multiple consecutive spaces
- Leading/trailing spaces
- Text with only spaces

### Acceptance Criteria
- [ ] Splits by whitespace correctly
- [ ] Handles multiple spaces
- [ ] Round-trip works: `detokenize(tokenize(x)) ≈ x`
- [ ] All test scenarios pass

### AI Guidance
**Ask:** "What's the difference between strings.Split and strings.Fields in Go? Which is better for tokenization?"

---

## 🧩 TASK-006: Hexadecimal Conversion

**Story Points:** 3/5 (Moderate)  
**Time:** 2-3 hours  
**Prerequisites:** TASK-005

### Learning Objectives
- Number base conversion (base 16 → base 10)
- Lookahead logic in token processing
- Error handling for invalid input

### What to Build
Function that processes tokens and replaces `word (hex)` with decimal equivalent.

### Test Scenarios
**Valid conversions:**
- `["1E", "(hex)", "files"]` → `["30", "files"]`
- `["FF", "(hex)", "max"]` → `["255", "max"]`
- `["ff", "(hex)"]` (lowercase) → `["255"]`
- `["A", "(hex)"]` → `["10"]`

**Invalid/Edge cases:**
- `["XYZ", "(hex)"]` (invalid hex) → `["XYZ", "(hex)"]` (unchanged)
- `["(hex)", "word"]` (no previous word) → `["(hex)", "word"]` (unchanged)
- `["1E", "(hex)", "and", "FF", "(hex)"]` → `["30", "and", "255"]` (multiple)

### Architecture Decision
**Question:** What to do with invalid hex?  
**Answer:** Leave unchanged (don't crash, don't remove)

### Acceptance Criteria
- [ ] Valid hex (0-9, A-F) converts correctly
- [ ] Case-insensitive (1E and 1e both work)
- [ ] Invalid hex left unchanged
- [ ] (hex) marker removed after conversion
- [ ] All test scenarios pass

### AI Guidance
**Ask:** "How do I parse a hexadecimal string in Go? What function converts from base 16 to base 10?"

### Resources
- [strconv.ParseInt documentation](https://pkg.go.dev/strconv#ParseInt)

---

## 🧩 TASK-007: Binary Conversion

**Story Points:** 2/5 (Simple - similar to hex)  
**Time:** 1 hour  
**Prerequisites:** TASK-006

### Learning Objectives
- Binary number system (base 2)
- Applying similar pattern to hex conversion

### What to Build
Function that processes tokens and replaces `word (bin)` with decimal equivalent.

### Test Scenarios
- `["10", "(bin)"]` → `["2"]`
- `["1010", "(bin)"]` → `["10"]`
- `["102", "(bin)"]` (invalid - has 2) → `["102", "(bin)"]` (unchanged)

### Architecture Decision
**Refactoring logic:** Hex and binary conversion have identical logic, just different base.  
**Consider:** Generic conversion function that accepts base as parameter.

### Acceptance Criteria
- [ ] Valid binary (0-1 only) converts correctly
- [ ] Invalid binary left unchanged
- [ ] Golden Test Case 1 now passes (has both hex and bin)
- [ ] All test scenarios pass

### AI Guidance
**Ask:** "I have two functions that are almost identical (hex and binary conversion). How do I refactor to avoid duplication?"

---

## 🧩 TASK-008: Uppercase Transformation

**Story Points:** 3/5 (Moderate)  
**Time:** 2-3 hours  
**Prerequisites:** TASK-007

### Learning Objectives
- Command parsing with parameters: `(up, 3)`
- Backward transformation (affects previous words)
- Handling count that exceeds available words

### What to Build
Function that processes `(up)` and `(up, N)` commands.

### Test Scenarios
**Single word:**
- `["ready", "set", "go", "(up)"]` → `["ready", "set", "GO"]`

**Multiple words:**
- `["this", "is", "exciting", "(up, 2)"]` → `["this", "IS", "EXCITING"]`

**Count exceeds available:**
- `["only", "two", "(up, 10)"]` → `["ONLY", "TWO"]` (transforms all available)

**Edge cases:**
- `["(up)", "word"]` (no previous word) → `["(up)", "word"]` (unchanged)
- `["LOUD", "(up)"]` (already uppercase) → `["LOUD"]`

### Architecture Decision
**Question:** How to parse `(up, 3)`?  
**Answer:** Use regex to extract the number, or manual string parsing

### Acceptance Criteria
- [ ] `(up)` transforms 1 previous word
- [ ] `(up, N)` transforms N previous words
- [ ] Count > available transforms all available
- [ ] Command marker removed
- [ ] All test scenarios pass

### AI Guidance
**Ask:** "How do I extract a number from a string like '(up, 5)' in Go? Should I use regex or string parsing?"

---

## 🧩 TASK-009: Lowercase Transformation

**Story Points:** 2/5 (Simple - similar to uppercase)  
**Time:** 1 hour  
**Prerequisites:** TASK-008

### Learning Objectives
- Applying same pattern as uppercase
- Recognizing code duplication

### What to Build
Function that processes `(low)` and `(low, N)` commands.

### Test Scenarios
- `["STOP", "SHOUTING", "(low)"]` → `["STOP", "shouting"]`
- `["WHY", "ARE", "WE", "YELLING", "(low, 4)"]` → `["why", "are", "we", "yelling"]`
- `["TWO", "WORDS", "(low, 10)"]` → `["two", "words"]`

### Architecture Decision
**Observation:** This is almost identical to TASK-008 (uppercase).  
**Next Sprint:** Refactor to remove duplication.

### Acceptance Criteria
- [ ] Same behavior as uppercase, but lowercases
- [ ] All test scenarios pass

---

## 🧩 TASK-010: Capitalize Transformation

**Story Points:** 3/5 (Moderate - Unicode handling)  
**Time:** 2-3 hours  
**Prerequisites:** TASK-009

### Learning Objectives
- Runes vs bytes in Go
- Unicode-safe string manipulation
- Title case vs capitalize

### What to Build
Function that processes `(cap)` and `(cap, N)` commands.  
**Behavior:** First letter uppercase, rest lowercase.

### Test Scenarios
- `["welcome", "to", "the", "brooklyn", "bridge", "(cap)"]` → `[...,"Bridge"]`
- `["the", "new", "york", "times", "(cap, 4)"]` → `["The", "New", "York", "Times"]`
- `["WORD", "(cap)"]` → `["Word"]` (first up, rest down)

### Architecture Decision
**Question:** Use `strings.Title` or manual implementation?  
**Answer:** Manual (strings.Title deprecated and behaves differently)

**Important:** Must handle Unicode correctly!  
- Bad: `word[0]` (breaks with multi-byte characters)
- Good: Convert to runes first

### Acceptance Criteria
- [ ] First letter uppercase, rest lowercase
- [ ] Works with already-uppercase words
- [ ] Unicode-safe implementation
- [ ] All test scenarios pass

### AI Guidance
**Ask:** "Why should I use runes instead of byte indexing when capitalizing words in Go?"

### Resources
- [Go Strings, Bytes, Runes](https://go.dev/blog/strings)

---

## 🧩 TASK-011: Refactor Case Transformations (DRY)

**Story Points:** 3/5 (Moderate)  
**Time:** 2 hours  
**Prerequisites:** TASK-010

### Learning Objectives
- DRY (Don't Repeat Yourself) principle
- Generic function design
- Higher-order functions in Go

### What to Build
Refactor tasks 008, 009, 010 to share common logic.

### Refactoring Strategy
**Observation:** All three case transformations:
1. Parse command `(cmd)` or `(cmd, N)`
2. Find previous N words
3. Apply transformation function
4. Remove command marker

**Generic function signature:**
```
transformCase(tokens, commandType, transformFunc)
```

### Test Scenarios
N/A - **All existing tests must still pass** after refactoring!

### Architecture Decision
**Trade-off:** DRY vs simplicity  
- **Pro:** Less code, easier to maintain
- **Con:** Slightly more complex

**Decision:** Go for it - the pattern is clear.

### Acceptance Criteria
- [ ] All existing tests pass
- [ ] Code duplication reduced
- [ ] Each case transformation uses generic helper
- [ ] Code is more maintainable

### AI Guidance
**Ask:** "I have three similar functions. How do I create a generic version that accepts a transformation function as a parameter in Go?"

---

## ✅ Sprint 1 Completion

Before moving to Sprint 2:

**Functional:**
- [ ] All 5 transformations work independently
- [ ] Golden Test Case 1 passes completely:
  ```
  Input: "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
  Output: "Simply add 66 and 2 and you will see the result is 68."
  ```

**Technical:**
- [ ] All tests pass
- [ ] Code refactored (DRY)
- [ ] No duplication

---

## 🎓 What You Learned

- ✅ Token-based text processing
- ✅ Number base conversion
- ✅ String case manipulation
- ✅ Regular expressions for pattern matching
- ✅ Unicode handling (runes)
- ✅ Code refactoring and DRY principle

---

**Next:** [`SPRINT-2-ADVANCED-TRANSFORMATIONS.md`](./SPRINT-2-ADVANCED-TRANSFORMATIONS.md) 🚀