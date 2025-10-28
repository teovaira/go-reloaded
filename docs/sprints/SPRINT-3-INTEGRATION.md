# Sprint 3: End-to-End Testing & Documentation

**Sprint Goal:** Validate complete system with golden tests and finalize documentation

**Duration:** 2-3 days | **Story Points:** 21

---

## Sprint Backlog

| Task ID | Description | Points |
|---------|-------------|--------|
| TASK-020 | Detokenization | 3 |
| TASK-021 | End-to-end integration | 4 |
| TASK-022 | Golden file testing | 5 |
| TASK-023 | Error handling standardization | 3 |
| TASK-024 | Technical documentation | 3 |
| TASK-025 | Final quality checks | 3 |

---

## TASK-020: Detokenization

### Functionality Description
Convert token slice back to string. Join tokens with single spaces. Handle special cases where tokens are already attached (like punctuation).

### Test Writing (TDD - Red Phase)
Write tests for:
- Basic: `["hello", "world"]` → `"hello world"`
- With punctuation: `["hello,", "world!"]` → `"hello, world!"`
- With quotes: `["'hello'", "world"]` → `"'hello' world"`
- Single token: `["hello"]` → `"hello"`
- Empty: `[]` → `""`

### Implementation Goal (TDD - Green Phase)
Create detokenization function that:
- Joins tokens with single space
- Handles empty slice
- Returns final string ready for output
- No special logic needed (transformations already applied)

### Validation (TDD - Refactor Phase)
- All tests pass
- Coverage ≥ 90%
- Commit: `feat: add detokenization function`

### Learning Resources
- [Go strings.Join](https://pkg.go.dev/strings#Join)

---

## TASK-021: End-to-End Integration

### Functionality Description
Connect all components: CLI → read file → tokenize → pipeline → detokenize → write file. Complete the main function with full workflow.

### Test Writing (TDD - Red Phase)
Write integration tests:
- Create test input files in `testdata/input/`
- Test full workflow: read → process → write
- Verify output file contents match expected
- Test error cases: missing input file, write permissions
- Test with actual golden test cases

### Implementation Goal (TDD - Green Phase)
Complete main function that:
- Uses TASK-005 CLI skeleton
- Calls TASK-006 tokenization
- Calls TASK-019 pipeline
- Calls TASK-020 detokenization
- Writes result to output file
- Handles all errors gracefully

### Validation (TDD - Refactor Phase)
- All integration tests pass
- End-to-end workflow complete
- Error handling robust
- Coverage ≥ 85%
- Commit: `feat: complete end-to-end integration`

### Learning Resources
- [Integration testing in Go](https://go.dev/doc/tutorial/add-a-test)
- [Table-driven tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)

---

## TASK-022: Golden File Testing

### Functionality Description
Implement all 12 test cases from GOLDEN-TEST-SET.md. Create input/expected file pairs. Verify program output matches expected output exactly.

### Test Writing (TDD - Red Phase)
Create golden test files in `testdata/`:
- Test 1-4: Official audit examples
- Test 5-11: Tricky edge cases
- Test 12: Complex integration test
- Each test has `input_N.txt` and `expected_N.txt`

Write test function that:
- Reads each input file
- Processes through pipeline
- Compares output to expected file
- Reports differences clearly

### Implementation Goal (TDD - Green Phase)
Implement golden file test runner:
- Iterate through all test file pairs
- Run program on each input
- Compare actual vs expected output
- Generate diff on mismatch
- All 12 tests must pass

### Validation (TDD - Refactor Phase)
- All 12 golden tests pass
- Test data files committed to Git (in testdata/ directory)
- Clear failure messages
- Commit: `test: add all 12 golden file tests`

### Learning Resources
- [Go testdata conventions](https://pkg.go.dev/cmd/go#hdr-Test_packages)
- [File comparison in tests](https://pkg.go.dev/os#ReadFile)

---

## TASK-023: Error Handling Standardization

### Functionality Description
Review and improve all error messages. Add context to errors (position, value, operation). Ensure consistent error handling patterns throughout codebase.

### Test Writing (TDD - Red Phase)
Write tests for error scenarios:
- Error messages contain context
- Errors include relevant values
- All error paths tested
- No generic "error" messages

### Implementation Goal (TDD - Green Phase)
Improve error handling:
- Add context to all errors (what failed, why, where)
- Use descriptive error messages
- Ensure errors are actionable
- Review all error returns in codebase

### Validation (TDD - Refactor Phase)
- All error paths tested
- Error messages descriptive and helpful
- No generic error messages remain
- Coverage includes error cases
- Commit: `refactor: improve error messages with context`

### Learning Resources
- [Go error handling](https://go.dev/blog/error-handling-and-go)
- [Error wrapping in Go](https://go.dev/blog/go1.13-errors)

---

## TASK-024: Technical Documentation

### Functionality Description
Complete project documentation: README with usage instructions, architecture overview, function documentation with godoc comments.

### Test Writing (TDD - Red Phase)
Not applicable - documentation task.

### Implementation Goal (TDD - Green Phase)
Create comprehensive documentation:
- Update README with:
  - Installation instructions
  - Usage examples
  - Feature list
  - Testing instructions
- Add godoc comments to all exported functions
- Document transformation order and why it matters
- Add architecture diagram or explanation
- Verify all examples work

### Validation (TDD - Refactor Phase)
- README complete and accurate
- All public functions documented
- Examples tested and working
- Documentation clear for other students
- Commit: `docs: add comprehensive project documentation`

### Learning Resources
- [Effective Go documentation](https://go.dev/doc/effective_go)
- [Godoc best practices](https://go.dev/blog/godoc)

---

## TASK-025: Final Quality Checks

### Functionality Description
Comprehensive final review before project delivery. Run all quality tools, verify all requirements met, ensure code is production-ready.

### Test Writing (TDD - Red Phase)
Not applicable - verification task.

### Implementation Goal (TDD - Green Phase)
Complete final checklist:
- Run `go fmt ./...` - all code formatted
- Run `go vet ./...` - no warnings
- Run `go test -v` - all tests pass
- Run `go test -cover` - coverage ≥ 90%
- Verify all 12 golden tests pass
- Check git history is clean
- Verify README accurate
- Test program manually with various inputs
- Ensure ~200-300 lines of code (concise)

### Validation (TDD - Refactor Phase)
- All automated checks pass
- Coverage ≥ 90%
- All golden tests pass
- Code is concise and readable
- Project delivery-ready
- Commit: `chore: final quality checks and polish`

### Learning Resources
- [Go code quality checklist](https://go.dev/wiki/CodeReviewComments)
- [Go best practices](https://peter.bourgon.org/go-best-practices-2016/)

---

## Sprint Success Criteria

- ✅ All 6 tasks complete
- ✅ All 12 golden tests pass
- ✅ Code coverage ≥ 90%
- ✅ Documentation complete
- ✅ Code is ~200-300 lines
- ✅ Project production-ready

---

## Dependencies

- TASK-020 requires TASK-019 complete
- TASK-021 requires TASK-005, 006, 019, 020 complete
- TASK-022 requires TASK-021 complete
- TASK-023, 024 can be done in parallel
- TASK-025 should be done last

---

## Sprint Notes

**Common Mistakes:**
- Skipping error handling tests
- Not testing with actual files
- Forgetting to update documentation
- Not running quality tools before submission

**Success Indicators:**
- All tests green
- Code is readable and maintainable
- Other students can understand your code
- You can explain every design decision

**Project Complete!** 🎉
