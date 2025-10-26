# Sprint 1: Core Text Transformation Functions

**Sprint Goal:** Implement fundamental text processing (tokenization, number conversion, case transformations)

**Duration:** 3-4 days | **Story Points:** 28

---

## Sprint Backlog

| Task ID | Description | Points |
|---------|-------------|--------|
| TASK-006 | Tokenize input text | 3 |
| TASK-007 | Command detection & parsing | 4 |
| TASK-008 | Hex to decimal | 3 |
| TASK-009 | Binary to decimal | 3 |
| TASK-010 | Uppercase transformation | 3 |
| TASK-011 | Lowercase transformation | 3 |
| TASK-012 | Capitalize words | 4 |
| TASK-013 | Refactor case functions | 3 |
| TASK-014 | Apply number conversions | 2 |

---

## TASK-006: Tokenize Input Text

### Functionality Description
Split input string into processable tokens (words and commands). Handles all whitespace types (space, tab, newline). Preserves punctuation and command markers like `(hex)`, `(cap)`.

### Test Writing (TDD - Red Phase)
Write tests that verify:
- Basic word splitting: `"hello world"` → `["hello", "world"]`
- Multiple/mixed whitespace preserved as single space
- Command preservation: `"word (cap)"` → `["word", "(cap)"]`
- Punctuation as separate tokens: `"hello,"` → `["hello", ","]`
- Empty input handling
- Leading/trailing whitespace removal

### Implementation Goal (TDD - Green Phase)
Create tokenization function that:
- Splits text into words, commands, and punctuation
- Returns slice of tokens in order
- Preserves command markers like `(hex)`, `(cap, 2)`
- Uses Go's `strings` package

### Validation (TDD - Refactor Phase)
- All tests pass
- Coverage ≥ 90%
- Clean code (no duplication)
- Commit: `feat: add tokenization function`

### Learning Resources
- [Go strings package](https://pkg.go.dev/strings)
- [strings.Fields documentation](https://pkg.go.dev/strings#Fields)

---

## TASK-007: Command Detection & Parsing

### Functionality Description
Identify and parse command markers in token stream. Detect `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)`, and count variants like `(up, 3)`. Extract command type and optional count parameter.

### Test Writing (TDD - Red Phase)
Write tests for:
- Simple commands: `"(hex)"` → type: hex, count: 1
- Count commands: `"(up, 3)"` → type: up, count: 3
- Invalid commands: `"(invalid)"`, `"(up, -1)"` → not recognized
- Edge cases: `"(cap,5)"` (no space), `"( hex )"` (extra spaces)
- Non-commands: `"hello"` → not a command

### Implementation Goal (TDD - Green Phase)
Create command parser that:
- Detects if token is a command (starts with `(` and ends with `)`)
- Extracts command type (hex, bin, up, low, cap)
- Extracts optional count parameter
- Returns command struct or error for invalid commands
- Handles whitespace variations

### Validation (TDD - Refactor Phase)
- All tests pass including edge cases
- Invalid commands handled gracefully
- Coverage ≥ 90%
- Commit: `feat: add command detection and parsing`

### Learning Resources
- [Go strings.Split](https://pkg.go.dev/strings#Split)
- [Go strconv.Atoi](https://pkg.go.dev/strconv#Atoi)
- [Go strings.TrimSpace](https://pkg.go.dev/strings#TrimSpace)

---

## TASK-008: Hex to Decimal Conversion

### Functionality Description
Convert hexadecimal strings to decimal integers. Input: `"1E"`, Output: `30`. Case-insensitive (1E = 1e). No "0x" prefix required.

### Test Writing (TDD - Red Phase)
Write tests for:
- Valid hex: `"1E"` → `30`, `"FF"` → `255`, `"A"` → `10`
- Case insensitivity: `"1e"` and `"1E"` both work
- Invalid input: `"ZZ"`, `"XYZ"`, `""` should return errors
- Edge cases: `"0"`, `"00A"` (leading zeros)

### Implementation Goal (TDD - Green Phase)
Create hex conversion function that:
- Accepts hex string, returns integer and error
- Uses `strconv.ParseInt` with base 16
- Returns descriptive errors for invalid input
- Handles both uppercase and lowercase letters

### Validation (TDD - Refactor Phase)
- All tests pass including error cases
- Error messages are clear
- Coverage ≥ 90%
- Commit: `feat: add hex to decimal conversion`

### Learning Resources
- [strconv.ParseInt documentation](https://pkg.go.dev/strconv#ParseInt)
- [Hexadecimal number system](https://www.mathsisfun.com/hexadecimals.html)

---

## TASK-009: Binary to Decimal Conversion

### Functionality Description
Convert binary strings to decimal integers. Input: `"101"`, Output: `5`. Only accepts '0' and '1' characters. No "0b" prefix required.

### Test Writing (TDD - Red Phase)
Write tests for:
- Valid binary: `"10"` → `2`, `"101"` → `5`, `"1010"` → `10`, `"11111111"` → `255`
- Invalid input: `"102"`, `"abc"`, `"12"` should return errors
- Edge cases: `"0"`, `"1"`, `"0010"` (leading zeros)

### Implementation Goal (TDD - Green Phase)
Create binary conversion function that:
- Accepts binary string, returns integer and error
- Uses `strconv.ParseInt` with base 2
- Validates input contains only 0 and 1
- Returns descriptive errors for invalid input

### Validation (TDD - Refactor Phase)
- All tests pass including error cases
- Clear error messages
- Coverage ≥ 90%
- Commit: `feat: add binary to decimal conversion`

### Learning Resources
- [strconv.ParseInt with base 2](https://pkg.go.dev/strconv#ParseInt)
- [Binary number system](https://www.mathsisfun.com/binary-number-system.html)

---

## TASK-010: Uppercase Transformation

### Functionality Description
Convert text to uppercase. Input: `"hello"`, Output: `"HELLO"`. Preserves numbers, punctuation, whitespace. Handles Unicode correctly (café → CAFÉ).

### Test Writing (TDD - Red Phase)
Write tests for:
- Basic: `"hello"` → `"HELLO"`
- Already upper: `"HELLO"` → `"HELLO"` (idempotent)
- Mixed case: `"Hello"` → `"HELLO"`
- With numbers/punctuation: `"hello123!"` → `"HELLO123!"`
- Unicode: `"café"` → `"CAFÉ"`
- Empty string handling

### Implementation Goal (TDD - Green Phase)
Create uppercase function that:
- Uses Go's `strings.ToUpper()` for Unicode support
- Preserves non-alphabetic characters
- Returns new string (immutable)
- Works on single words

### Validation (TDD - Refactor Phase)
- All tests pass including Unicode
- Coverage ≥ 90%
- Commit: `feat: add uppercase transformation`

### Learning Resources
- [strings.ToUpper documentation](https://pkg.go.dev/strings#ToUpper)
- [Unicode in Go](https://go.dev/blog/strings)

---

## TASK-011: Lowercase Transformation

### Functionality Description
Convert text to lowercase. Input: `"HELLO"`, Output: `"hello"`. Preserves numbers, punctuation, whitespace. Handles Unicode correctly (CAFÉ → café).

### Test Writing (TDD - Red Phase)
Write tests for:
- Basic: `"HELLO"` → `"hello"`
- Already lower: `"hello"` → `"hello"` (idempotent)
- Mixed case: `"HeLLo"` → `"hello"`
- With numbers/punctuation: `"HELLO123!"` → `"hello123!"`
- Unicode: `"CAFÉ"` → `"café"`
- Empty string handling

### Implementation Goal (TDD - Green Phase)
Create lowercase function that:
- Uses Go's `strings.ToLower()` for Unicode support
- Preserves non-alphabetic characters
- Returns new string (immutable)
- Works on single words

### Validation (TDD - Refactor Phase)
- All tests pass including Unicode
- Coverage ≥ 90%
- Commit: `feat: add lowercase transformation`

### Learning Resources
- [strings.ToLower documentation](https://pkg.go.dev/strings#ToLower)
- [Unicode case mapping](https://unicode.org/reports/tr21/)

---

## TASK-012: Capitalize Words

### Functionality Description
Capitalize first letter of each word. Input: `"hello"`, Output: `"Hello"`. Requires manual rune manipulation for Unicode. First letter uppercase, rest lowercase.

### Test Writing (TDD - Red Phase)
Write tests for:
- Basic: `"hello"` → `"Hello"`
- All caps input: `"HELLO"` → `"Hello"` (first upper, rest lower)
- All lower: `"hello"` → `"Hello"`
- Mixed: `"hELLo"` → `"Hello"`
- Unicode: `"café"` → `"Café"`
- Empty and single character strings

### Implementation Goal (TDD - Green Phase)
Create capitalize function that:
- Converts string to runes for Unicode support
- Makes first rune uppercase using `unicode.ToUpper()`
- Makes remaining runes lowercase using `unicode.ToLower()`
- Handles multi-byte characters correctly
- Works on single words

### Validation (TDD - Refactor Phase)
- All tests pass including Unicode
- Rune manipulation correct
- Coverage ≥ 90%
- Commit: `feat: add word capitalization`

### Learning Resources
- [Go runes and characters](https://go.dev/blog/strings)
- [unicode.ToUpper documentation](https://pkg.go.dev/unicode#ToUpper)

---

## TASK-013: Refactor Case Functions

### Functionality Description
Eliminate code duplication across uppercase, lowercase, and capitalize functions. Extract common validation and helper functions.

### Test Writing (TDD - Red Phase)
No new tests needed - all existing tests must continue passing.

### Implementation Goal (TDD - Green Phase)
Refactor case transformation functions:
- Extract shared validation logic (empty string checks)
- Identify and eliminate code duplication
- Improve function names and clarity
- No behavior changes (all existing tests still pass)

### Validation (TDD - Refactor Phase)
- All original tests still pass
- Code duplication reduced ≥ 50%
- Coverage maintained at ≥ 90%
- Commit: `refactor: extract common case transformation logic`

### Learning Resources
- [Go refactoring patterns](https://dave.cheney.net/2019/07/09/clear-is-better-than-clever)
- [DRY principle](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself)

---

## TASK-014: Apply Number Conversions

### Functionality Description
Integrate hex and binary conversion with command parsing. Process token stream, detect `(hex)` and `(bin)` commands, apply conversions to previous word, replace with decimal result.

### Test Writing (TDD - Red Phase)
Write tests for:
- Basic hex: `["1E", "(hex)", "items"]` → `["30", "items"]`
- Basic binary: `["10", "(bin)", "years"]` → `["2", "years"]`
- Multiple conversions: `["A", "(hex)", "and", "11", "(bin)"]` → `["10", "and", "3"]`
- Invalid input: `["XYZ", "(hex)"]` → unchanged (graceful degradation)
- No previous word: `["(hex)", "word"]` → unchanged

### Implementation Goal (TDD - Green Phase)
Create conversion application function that:
- Iterates through token slice
- Detects conversion commands using TASK-007 parser
- Applies conversion to previous token
- Replaces both word and command with result
- Handles errors gracefully (skip invalid conversions)

### Validation (TDD - Refactor Phase)
- All tests pass including error cases
- Graceful degradation for invalid input
- Coverage ≥ 90%
- Commit: `feat: integrate number conversions with command parsing`

### Learning Resources
- [Go slice manipulation](https://go.dev/blog/slices-intro)
- [Go error handling patterns](https://go.dev/blog/error-handling-and-go)

---

## Sprint Success Criteria

- ✅ All 9 tasks complete with passing tests
- ✅ Code coverage ≥ 90%
- ✅ All functions handle Unicode correctly
- ✅ Number conversions integrated with pipeline
- ✅ Clean git history with meaningful commits

---

## Dependencies

- TASK-006 must be done first (tokenization)
- TASK-007 must be done second (command parsing)
- TASK-008, 009, 010, 011, 012 can be done in parallel
- TASK-013 requires TASK-010, 011, 012 complete
- TASK-014 requires TASK-007, 008, 009 complete

**Next:** Sprint 2 - Advanced text processing (articles, punctuation, quotes)
