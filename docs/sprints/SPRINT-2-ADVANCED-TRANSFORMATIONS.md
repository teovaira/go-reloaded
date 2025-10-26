# Sprint 2: Advanced Text Processing

**Sprint Goal:** Implement context-aware transformations and pipeline integration

**Duration:** 3-4 days | **Story Points:** 24

---

## Sprint Backlog

| Task ID | Description | Points |
|---------|-------------|--------|
| TASK-015 | Apply case transformations | 3 |
| TASK-016 | Article correction (a/an) | 5 |
| TASK-017 | Punctuation spacing | 6 |
| TASK-018 | Quote pairing | 5 |
| TASK-019 | Pipeline integration | 5 |

---

## TASK-015: Apply Case Transformations

### Functionality Description
Integrate case transformations (up, low, cap) with command parsing. Process token stream, detect case commands with optional counts like `(up, 3)`, apply to previous N words.

### Test Writing (TDD - Red Phase)
Write tests for:
- Single word: `["hello", "(up)"]` → `["HELLO"]`
- With count: `["one", "two", "three", "(up, 2)"]` → `["one", "TWO", "THREE"]`
- Lowercase: `["HELLO", "WORLD", "(low, 2)"]` → `["hello", "world"]`
- Capitalize: `["hello", "world", "(cap, 2)"]` → `["Hello", "World"]`
- Count exceeds words: `["word", "(up, 10)"]` → `["WORD"]` (graceful degradation)
- No previous word: `["(up)", "word"]` → unchanged

### Implementation Goal (TDD - Green Phase)
Create case application function that:
- Iterates through token slice
- Detects case commands using TASK-007 parser
- Applies transformation to previous N tokens (default N=1)
- Uses TASK-010, 011, 012 transformation functions
- Handles count > available words gracefully
- Removes command markers from output

### Validation (TDD - Refactor Phase)
- All tests pass including edge cases
- Graceful degradation working
- Coverage ≥ 90%
- Commit: `feat: integrate case transformations with command parsing`

### Learning Resources
- [Go slice manipulation](https://go.dev/blog/slices-intro)
- [Graceful degradation pattern](https://en.wikipedia.org/wiki/Fault_tolerance)

---

## TASK-016: Article Correction (a/an)

### Functionality Description
Replace "a" with "an" before words starting with vowels (a, e, i, o, u) or 'h'. Case-insensitive detection. Handles both "a" and "A".

### Test Writing (TDD - Red Phase)
Write tests for:
- Vowel sounds: `["a", "apple"]` → `["an", "apple"]`
- Consonants: `["a", "cat"]` → `["a", "cat"]` (unchanged)
- Letter h: `["a", "hour"]` → `["an", "hour"]`, `["a", "hero"]` → `["an", "hero"]`
- Capital A: `["A", "apple"]` → `["An", "apple"]`
- End of text: `["word", "a"]` → `["word", "a"]` (unchanged)
- Multiple: `["a", "apple", "and", "a", "banana"]` → `["an", "apple", "and", "a", "banana"]`

### Implementation Goal (TDD - Green Phase)
Create article correction function that:
- Iterates through token slice
- Detects "a" or "A" tokens
- Checks if next token starts with vowel or 'h'
- Replaces with "an" or "An" (preserves case)
- Handles end of slice gracefully

### Validation (TDD - Refactor Phase)
- All tests pass including edge cases
- Case preservation working
- Coverage ≥ 90%
- Commit: `feat: add article correction with vowel and h detection`

### Learning Resources
- [Go strings.HasPrefix](https://pkg.go.dev/strings#HasPrefix)
- [Go unicode.ToLower](https://pkg.go.dev/unicode#ToLower)

---

## TASK-017: Punctuation Spacing

### Functionality Description
Fix spacing around punctuation per English typography rules. Attach punctuation to previous word, add space after. Handles: . , ! ? ; : and groups like ... !?

### Test Writing (TDD - Red Phase)
Write tests for:
- Basic: `["hello", ",", "world"]` → `["hello,", "world"]`
- Multiple marks: `["hello", ".", ".", "."]` → `["hello..."]`
- Question/exclamation: `["what", "?"]` → `["what?"]`
- Consecutive: `["word", "!", "?"]` → `["word!?"]`
- At start: `[".", "word"]` → `[".", "word"]` (unchanged)
- Mixed: `["test", ",", "wait", ".", ".", ".", "really", "?"]` → `["test,", "wait...", "really?"]`

### Implementation Goal (TDD - Green Phase)
Create punctuation fixing function that:
- Identifies punctuation tokens (. , ! ? ; :)
- Attaches to previous word token
- Merges consecutive punctuation into groups
- Handles start of text gracefully
- Returns modified token slice

### Validation (TDD - Refactor Phase)
- All tests pass
- All punctuation types handled
- Coverage ≥ 90%
- Commit: `feat: add punctuation spacing correction`

### Learning Resources
- [English punctuation rules](https://en.wikipedia.org/wiki/Sentence_spacing)
- [Go strings.Contains](https://pkg.go.dev/strings#Contains)

---

## TASK-018: Quote Pairing

### Functionality Description
Replace single quotes with proper pairing. Uses state tracking: first `'` opens, second `'` closes, third `'` opens new pair. Attach quotes to adjacent words.

### Test Writing (TDD - Red Phase)
Write tests for:
- Basic pair: `["'", "hello", "'"]` → `["'hello'"]`
- Two pairs: `["'", "a", "'", "and", "'", "b", "'"]` → `["'a'", "and", "'b'"]`
- Multi-word: `["'", "hello", "world", "'"]` → `["'hello", "world'"]`
- Odd number: `["'", "word"]` → `["'word"]` (opening only)
- No quotes: `["hello", "world"]` → unchanged
- Mid-sentence: `["he", "said", "'", "hi", "'"]` → `["he", "said", "'hi'"]`

### Implementation Goal (TDD - Green Phase)
Create quote pairing function that:
- Uses boolean toggle or state variable (open/closed)
- Alternates between opening and closing quotes
- Attaches opening quote to next word
- Attaches closing quote to previous word
- Handles odd number gracefully (leave last as opening)
- Returns modified token slice

### Validation (TDD - Refactor Phase)
- All tests pass
- State machine correct
- Coverage ≥ 90%
- Commit: `feat: add quote pairing with state tracking`

### Learning Resources
- [State machines](https://en.wikipedia.org/wiki/Finite-state_machine)
- [Go boolean toggle pattern](https://gobyexample.com/variables)

---

## TASK-019: Pipeline Integration

### Functionality Description
Connect all transformations in correct order. Create main processing pipeline that applies transformations sequentially: numbers → case → articles → punctuation → quotes.

### Test Writing (TDD - Red Phase)
Write tests for:
- Full pipeline: `["1E", "(hex)", "is", "a", "example", "(cap)"]` → `["30", "is", "an", "Example"]`
- Multiple rules: `["it", "(cap)", "was", "a", "apple", ",", "really", "!"]` → `["It", "was", "an", "apple,", "really!"]`
- All transformations: Test from GOLDEN-TEST-SET.md Test 12
- Order matters: Verify case before articles, articles before punctuation
- Empty input: `[]` → `[]`

### Implementation Goal (TDD - Green Phase)
Create pipeline orchestration function that:
- Takes token slice as input
- Applies transformations in order:
  1. Number conversions (TASK-014)
  2. Case transformations (TASK-015)
  3. Article correction (TASK-016)
  4. Punctuation spacing (TASK-017)
  5. Quote pairing (TASK-018)
- Returns final transformed token slice
- Each stage receives output of previous stage

### Validation (TDD - Refactor Phase)
- All tests pass
- Transformation order correct
- Integration with all previous tasks working
- Coverage ≥ 90%
- Commit: `feat: implement transformation pipeline orchestration`

### Learning Resources
- [Pipeline pattern](https://go.dev/blog/pipelines)
- [Function composition in Go](https://dave.cheney.net/2016/01/18/codereviews)

---

## Sprint Success Criteria

- ✅ All 5 tasks complete with passing tests
- ✅ Context-aware processing working
- ✅ Code coverage ≥ 90%
- ✅ Pipeline integration complete
- ✅ All transformations working together correctly

---

## Dependencies

- TASK-015 requires TASK-007, 010, 011, 012 complete
- TASK-016, 017, 018 can be done in parallel
- TASK-019 requires all previous tasks (014, 015, 016, 017, 018) complete

**Next:** Sprint 3 - Integration testing, documentation, final polish
