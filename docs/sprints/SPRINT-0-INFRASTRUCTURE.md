# Sprint 0: Infrastructure & Foundation

**Goal:** Set up project structure and implement basic file I/O with CLI handling.  
**Duration:** 1-2 days  
**Tasks:** 4  
**Deliverable:** Program reads input file and writes output file (no transformations yet)

---

## 🎯 Sprint Objectives

By end of Sprint 0:
- ✅ Go module initialized
- ✅ CLI arguments validated
- ✅ File reading works
- ✅ File writing works
- ✅ Simple test exists

---

## 🧩 TASK-001: Project Initialization

**Story Points:** 1/5 (Trivial)  
**Time:** 30 minutes

### Learning Objectives
- Go module system (`go.mod`)
- Project structure for small Go projects
- `.gitignore` for Go

### What to Build
Initialize Go project with:
- `go.mod` file
- `main.go` (empty `main()` function)
- `main_test.go` (empty test file)
- `.gitignore` (Go patterns)

### Test Scenarios
N/A - Infrastructure setup

### Acceptance Criteria
- [ ] `go build` runs without errors
- [ ] `go test` shows "no tests to run"
- [ ] `.gitignore` prevents committing binaries

### AI Guidance
**Ask:** "What should I include in .gitignore for a Go CLI project?"

### Resources
- [How to Write Go Code](https://go.dev/doc/code)
- [Go Modules Reference](https://go.dev/ref/mod)

---

## 🧩 TASK-002: CLI Argument Validation

**Story Points:** 2/5 (Simple)  
**Time:** 1-2 hours  
**Prerequisites:** TASK-001

### Learning Objectives
- `os.Args` for command-line parsing
- Input validation patterns
- Exit codes (0=success, 1=error)

### What to Build
Function that validates exactly 2 arguments (input file, output file).

### Test Scenarios
**Test 1:** Valid arguments  
- Input: `["program", "in.txt", "out.txt"]`
- Expected: No error

**Test 2:** Missing output file  
- Input: `["program", "in.txt"]`
- Expected: Error with usage message

**Test 3:** No arguments  
- Input: `["program"]`
- Expected: Error with usage message

**Test 4:** Too many arguments  
- Input: `["program", "a.txt", "b.txt", "c.txt"]`
- Expected: Error with usage message

### Edge Cases
- Empty filename strings
- Very long filenames

### Acceptance Criteria
- [ ] Exactly 2 arguments required
- [ ] Usage message shown on error
- [ ] All 4 test cases pass

### AI Guidance
**Ask:** "In Go, why is os.Args[0] the program name and how do I validate argument count?"

### Resources
- [Go by Example: Command-Line Arguments](https://gobyexample.com/command-line-arguments)

---

## 🧩 TASK-003: File Reading

**Story Points:** 2/5 (Simple)  
**Time:** 1-2 hours  
**Prerequisites:** TASK-002

### Learning Objectives
- `os.ReadFile` for reading files
- Error handling patterns
- Test fixtures with `testdata/` directory

### What to Build
Function that reads file contents into a string.

### Test Scenarios
**Test 1:** Read existing file  
- Create `testdata/sample.txt` with "Hello, World!"
- Expected: Function returns "Hello, World!"

**Test 2:** Non-existent file  
- Filename: `testdata/nonexistent.txt`
- Expected: Error returned

**Test 3:** Empty filename  
- Input: `""`
- Expected: Error returned

### Edge Cases
- Empty files (valid - return empty string)
- Very large files (>1MB)
- Files without read permissions

### Acceptance Criteria
- [ ] Successfully reads text files
- [ ] Returns error for missing files
- [ ] All 3 test cases pass

### AI Guidance
**Ask:** "What's the difference between os.ReadFile and os.Open in Go? Which is simpler for reading entire files?"

### Resources
- [Go by Example: Reading Files](https://gobyexample.com/reading-files)
- [testdata convention in Go](https://dave.cheney.net/2016/05/10/test-fixtures-in-go)

---

## 🧩 TASK-004: File Writing & Pipeline Skeleton

**Story Points:** 3/5 (Moderate)  
**Time:** 2-3 hours  
**Prerequisites:** TASK-003

### Learning Objectives
- `os.WriteFile` for writing files
- Function composition (pipeline pattern)
- Integration testing

### What to Build
1. Function that writes string to file
2. `processText()` function (currently just returns input unchanged)
3. `main()` that connects: read → process → write

### Test Scenarios
**Test 1:** Write content to file  
- Input: "Test output"
- Expected: File created with exact content

**Test 2:** Overwrite existing file  
- Create file, write new content
- Expected: File replaced with new content

**Test 3:** End-to-end pipeline  
- Create input.txt: "Hello"
- Run: `main` with input.txt and output.txt
- Expected: output.txt contains "Hello" (unchanged for now)

### Edge Cases
- Writing to protected directory (permission denied)
- Very long file paths
- Empty string content (valid)

### Acceptance Criteria
- [ ] File writing works correctly
- [ ] `processText()` passes through input unchanged (stub)
- [ ] Full pipeline works: read → process → write
- [ ] All 3 test cases pass

### AI Guidance
**Ask:** "What file permissions should I use with os.WriteFile for a text file? What does 0644 mean?"

### Resources
- [Go by Example: Writing Files](https://gobyexample.com/writing-files)
- [Unix File Permissions](https://chmod-calculator.com/)

---

## ✅ Sprint 0 Completion

Before moving to Sprint 1:

**Functional:**
- [ ] Program reads input file
- [ ] Program writes output file
- [ ] CLI validation works

**Technical:**
- [ ] All tests pass
- [ ] `go build` succeeds
- [ ] No compiler warnings

**Manual Test:**
```
echo "Test" > input.txt
go run . input.txt output.txt
cat output.txt  # Should show: Test
```

---

## 🎓 What You Learned

- ✅ Go project setup
- ✅ Command-line argument parsing
- ✅ File I/O operations
- ✅ Basic testing in Go
- ✅ Pipeline architecture concept

---

**Next:** [`SPRINT-1-CORE-TRANSFORMATIONS.md`](./SPRINT-1-CORE-TRANSFORMATIONS.md) 🚀