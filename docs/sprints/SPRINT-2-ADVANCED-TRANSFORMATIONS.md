# Sprint 2: Advanced Transformations

**Goal:** Implement linguistic rules: article correction (a→an), punctuation spacing, and quote formatting  
**Duration:** 3-4 days  
**Tasks:** 5  
**Deliverable:** All 7 transformation types complete, Golden Test Cases 2-4 pass

---

## 🎯 Sprint Objectives

By end of Sprint 2:
- ✅ Article correction (a→an) working
- ✅ Punctuation spacing correct
- ✅ Quote pairing functional
- ✅ All golden tests 1-4 passing

---

## 🧩 TASK-012: Article Correction (a → an)

**Story Points:** 4/5 (Complex - lookahead logic)  
**Time:** 3-4 hours  
**Prerequisites:** TASK-011

### Learning Objectives
- Lookahead processing in token streams
- Case-insensitive string comparison
- Linguistic rule implementation

### What to Build
Function that changes "a" to "an" before vowels (a, e, i, o, u) or 'h'.

### Test Scenarios
**Basic cases:**
- `["a", "apple"]` → `["an", "apple"]`
- `["a", "elephant"]` → `["an", "elephant"]`
- `["a", "hour"]` → `["an", "hour"]`

**Case preservation:**
- `["A", "amazing"]` → `["An", "amazing"]`

**No change:**
- `["a", "cat"]` → `["a", "cat"]` (consonant)

**Edge cases:**
- `["a"]` (no next word) → `["a"]` (unchanged)
- `["a", "apple", "and", "a", "orange"]` → `["an", "apple", "and", "an", "orange"]`

### Architecture Decision
**Question:** What about 'university' (starts with vowel but sounds like consonant)?  
**Answer:** Spec says ALL vowels and 'h' - follow spec literally.

**Implementation note:** This requires lookahead (checking next token).

### Acceptance Criteria
- [ ] Converts "a" to "an" before a, e, i, o, u, h
- [ ] Case preserved (A → An)
- [ ] "a" before consonants unchanged
- [ ] Golden Test Case 2 passes
- [ ] All test scenarios pass

### AI Guidance
**Ask:** "How do I safely check the next token in a slice without going out of bounds?"

---

## 🧩 TASK-013: Punctuation Spacing

**Story Points:** 4/5 (Complex - multiple rules)  
**Time:** 3-4 hours  
**Prerequisites:** TASK-012

### Learning Objectives
- Context-aware formatting
- Punctuation grouping rules
- String reconstruction

### What to Build
Function that fixes punctuation spacing:
- Punctuation (. , ! ? : ;) attaches to previous word
- Groups like `...` and `!?` stay together

### Test Scenarios
**Basic spacing:**
- `["Hello", ",", "world"]` → `["Hello,", "world"]`
- `["End", "of", "sentence", "."]` → `["End", "of", "sentence."]`

**Punctuation groups:**
- `["Wait", ".", ".", "."]` → `["Wait..."]`
- `["What", "!", "?"]` → `["What!?"]`

**Multiple types:**
- `["Note", ":", "important", ";", "remember"]` → `["Note:", "important;", "remember"]`

### Architecture Decision
**Question:** What if punctuation is at start of text?  
**Answer:** Leave it unchanged (edge case, spec doesn't address).

**Strategy:**
1. Scan tokens
2. When you find punctuation, collect all consecutive punctuation
3. Attach the group to previous word

### Acceptance Criteria
- [ ] Punctuation attached to previous word
- [ ] No space before punctuation
- [ ] Groups (`...`, `!?`) stay together
- [ ] Golden Test Case 3 passes
- [ ] All test scenarios pass

### AI Guidance
**Ask:** "How do I collect consecutive punctuation marks and attach them to the previous token?"

---

## 🧩 TASK-014: Quote Formatting

**Story Points:** 5/5 (Very Complex - stateful pairing)  
**Time:** 4-5 hours  
**Prerequisites:** TASK-013

### Learning Objectives
- Pairing/matching algorithms
- State machines
- Handling unpaired elements

### What to Build
Function that pairs single quotes `'...'` and formats content between them.

### Test Scenarios
**Single word:**
- `["say", "'", "hello", "'"]` → `["say", "'hello'"]`

**Multiple words:**
- `["he", "said", ":", "'", "hello", "world", "'"]` → `["he", "said:", "'hello", "world'"]`

**Spacing rules:**
- Opening quote attaches to first word: `'hello`
- Closing quote attaches to last word: `world'`
- Spaces between words preserved

**Edge cases:**
- `["'", "word", "'", "and", "'", "another", "'"]` (multiple pairs)
- `["test", "'", "word"]` (unpaired) → Leave unchanged

### Architecture Decision
**Question:** What about nested quotes?  
**Answer:** Spec doesn't require it - keep simple, just pair first available quotes.

**Strategy:**
1. Find opening quote `'`
2. Scan forward for closing quote `'`
3. If found, group content between them
4. If not found, leave opening quote unchanged

### Acceptance Criteria
- [ ] Paired quotes combined correctly
- [ ] Single word: `'word'`
- [ ] Multiple words: `'first middle last'`
- [ ] Multiple pairs work independently
- [ ] Unpaired quotes left unchanged
- [ ] Golden Test Case 4 fully passes
- [ ] All test scenarios pass

### AI Guidance
**Ask:** "What's a good algorithm for finding matching pairs in a sequence? Should I use a stack?"

---

## 🧩 TASK-015: Integration Testing

**Story Points:** 2/5 (Simple)  
**Time:** 1-2 hours  
**Prerequisites:** TASK-014

### Learning Objectives
- Integration testing vs unit testing
- Test coverage measurement
- End-to-end verification

### What to Build
Comprehensive tests for all golden test cases (1-4) and complex combinations.

### Test Scenarios
**Golden Test Cases:**
- Test Case 1: Hex and binary conversions
- Test Case 2: Article correction
- Test Case 3: Punctuation spacing
- Test Case 4: All rules combined

**Complex integration:**
```
Input: "here (cap) is a interesting text with 1A (hex) items and 11 (bin) more , all in ' a epic document (cap, 2) ' ... what do you think (up, 4) ?"
Expected: "Here is an interesting text with 26 items and 3 more, all in 'an Epic Document'... WHAT DO YOU THINK?"
```

### Architecture Decision
**Question:** What's the pipeline order?  
**Answer:** (Must be documented!)
1. Number conversions (hex, bin)
2. Case transformations (up, low, cap)
3. Article correction (a→an)
4. Punctuation spacing
5. Quote formatting

**Rationale:** Case changes must happen before article correction (affects case of 'A').

### Acceptance Criteria
- [ ] All 4 golden tests pass
- [ ] Complex integration test passes
- [ ] Test coverage >80%
- [ ] All transformation orders correct

### AI Guidance
**Ask:** "How do I measure test coverage in Go? What does >80% coverage mean?"

---

## 🧩 TASK-016: Code Review & Cleanup

**Story Points:** 1/5 (Trivial)  
**Time:** 1 hour  
**Prerequisites:** TASK-015

### Learning Objectives
- Code quality standards
- Self-review process
- Go style conventions

### What to Build
N/A - This is cleanup/review, not new features.

### Review Checklist
**Code Quality:**
- [ ] Function names are clear
- [ ] No magic numbers or strings
- [ ] Comments explain "why" not "what"
- [ ] Consistent naming conventions

**Testing:**
- [ ] All tests pass
- [ ] Test names describe what they test
- [ ] Edge cases covered

**Go Standards:**
- [ ] Run `go fmt` (auto-formats code)
- [ ] Run `go vet` (finds potential bugs)
- [ ] No compiler warnings

### Acceptance Criteria
- [ ] Code is clean and readable
- [ ] All tools (fmt, vet) pass
- [ ] Ready for Sprint 3 (final polish)

### AI Guidance
**Ask:** "Review my code structure - are my function names clear? Is there duplication I missed?"

---

## ✅ Sprint 2 Completion

Before moving to Sprint 3:

**Functional:**
- [ ] All 7 transformation types complete
- [ ] Golden Test Cases 1, 2, 3, 4 pass
- [ ] Complex integration test passes

**Technical:**
- [ ] All tests pass
- [ ] Code formatted (`go fmt`)
- [ ] No vet warnings
- [ ] Test coverage >80%

**Manual Verification:**
```
Test each golden case manually to ensure outputs match exactly!
```

---

## 🎓 What You Learned

- ✅ Lookahead logic (article correction)
- ✅ Stateful processing (quote pairing)
- ✅ Complex string manipulation
- ✅ Integration testing strategies
- ✅ Pipeline ordering importance
- ✅ Code review skills

---

**Next:** [`SPRINT-3-INTEGRATION.md`](./SPRINT-3-INTEGRATION.md) - Final polish! 🚀