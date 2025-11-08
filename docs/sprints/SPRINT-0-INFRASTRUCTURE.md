# Sprint 0: Infrastructure & Project Setup

**Sprint Goal:** Establish project foundation, development environment, and testing infrastructure

**Duration:** 1-2 days | **Story Points:** 13

---

## Sprint Backlog

| Task ID | Description | Points |
|---------|-------------|--------|
| TASK-001 | Project initialization | 3 |
| TASK-002 | Git repository setup | 2 |
| TASK-003 | Testing framework | 4 |
| TASK-004 | Development tools | 3 |
| TASK-005 | CLI skeleton | 3 |

---

## TASK-001: Project Initialization

### Functionality Description
Create Go module and project directory structure. Establish folder organization for source code, tests, and documentation.

### Test Writing (TDD - Red Phase)
Infrastructure task - manual verification:
- Verify `go.mod` exists and is valid
- Check all required directories created
- Confirm project builds without errors

### Implementation Goal (TDD - Green Phase)
Initialize project structure:
- Run `go mod init go-reloaded`
- Create directories: `tests/` 
- Create empty `main.go` with package declaration
- Verify `go build` succeeds

### Validation (TDD - Refactor Phase)
- `go.mod` file present with correct module name
- All directories created
- `go build` runs without errors
- Project structure matches plan
- Commit: `chore: initialize Go module and project structure`

### Learning Resources
- [Go modules tutorial](https://go.dev/doc/tutorial/create-module)
- [Go project layout](https://github.com/golang-standards/project-layout)

---

## TASK-002: Git Repository Setup

### Functionality Description
Initialize Git repository with proper configuration. Create `.gitignore` for Go projects. Set up branch structure and initial commit.

### Test Writing (TDD - Red Phase)
Manual verification:
- Git repository initialized
- `.gitignore` excludes build artifacts
- Initial commit contains base structure
- Remote repository connected (if applicable)

### Implementation Goal (TDD - Green Phase)
Setup version control:
- Run `git init`
- Create `.gitignore` with Go-specific exclusions (binaries, test coverage, IDE files, OS files)
- Create initial commit with existing docs
- Connect to remote repository if applicable

### Validation (TDD - Refactor Phase)
- `.git` directory exists
- `.gitignore` configured correctly
- Initial commit made
- Clean git status
- Commit: `chore: setup git repository with gitignore`

### Learning Resources
- [Git basics](https://git-scm.com/book/en/v2/Getting-Started-Git-Basics)
- [Go .gitignore templates](https://github.com/github/gitignore/blob/main/Go.gitignore)

---

## TASK-003: Testing Framework Setup

### Functionality Description
Configure Go testing environment. Create test file structure. Set up golden test directory. Prepare for TDD workflow.

### Test Writing (TDD - Red Phase)
Create sample test file to verify setup:
- Create `main_test.go` with placeholder test function
- Test should log "Testing framework is ready"
- Run `go test -v` - should pass

### Implementation Goal (TDD - Green Phase)
Setup testing infrastructure:
- Create `main_test.go` with sample test
- Create `testdata/` subdirectories: `input/` and `expected/`
- Verify test discovery works
- Run sample test successfully

### Validation (TDD - Refactor Phase)
- `go test -v` runs successfully
- Test file discovered automatically
- Test output is readable
- Golden test directories ready
- Commit: `test: setup testing framework and directories`

### Learning Resources
- [Go testing package](https://pkg.go.dev/testing)
- [Writing tests in Go](https://go.dev/doc/tutorial/add-a-test)
- [Table-driven tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)

---

## TASK-004: Development Tools Configuration

### Functionality Description
Install and configure essential Go development tools: formatter, linter, static analyzer. Set up pre-commit checks.

### Test Writing (TDD - Red Phase)
Manual verification:
- `go fmt` formats code
- `go vet` detects issues
- Optional: `golangci-lint` runs

### Implementation Goal (TDD - Green Phase)
Configure development tools:
- Verify `go fmt` works on project
- Verify `go vet` works on project
- Optional: Install and configure `golangci-lint`
- Optional: Create Makefile with common commands (test, fmt, vet)

### Validation (TDD - Refactor Phase)
- `go fmt` runs without errors
- `go vet` produces no warnings
- Linter configured (optional)
- Tools documented in README
- Commit: `chore: configure development tools and linters`

### Learning Resources
- [go fmt command](https://pkg.go.dev/cmd/gofmt)
- [go vet overview](https://pkg.go.dev/cmd/vet)
- [golangci-lint](https://golangci-lint.run/)

---

## TASK-005: CLI Skeleton Implementation

### Functionality Description
Create basic CLI that reads command-line arguments (input/output file paths). Validate arguments. Handle file I/O errors gracefully. No transformations yet.

### Test Writing (TDD - Red Phase)
Write tests in `main_test.go`:
- Test argument parsing (expect exactly 2 args)
- Test file reading (use testdata files)
- Test file writing (verify output created)
- Test error handling (missing file, wrong arg count)

### Implementation Goal (TDD - Green Phase)
Implement CLI skeleton in `main.go`:
- Parse command-line arguments (expect 2: input path, output path)
- Read input file content to string
- Write string to output file (no transformation yet - just copy)
- Return clear error messages for all failure cases

### Validation (TDD - Refactor Phase)
- All CLI tests pass
- Program runs: `go run . input.txt output.txt`
- Copies file correctly (no transformation)
- Error messages are clear
- Coverage ≥ 80% for CLI code
- Commit: `feat: add CLI skeleton with file I/O`

### Learning Resources
- [Go CLI applications](https://go.dev/doc/tutorial/getting-started)
- [os package for file I/O](https://pkg.go.dev/os)
- [Error handling in Go](https://go.dev/blog/error-handling-and-go)

---

## Sprint Success Criteria

- ✅ All 5 tasks complete
- ✅ Project structure established
- ✅ Git repository configured
- ✅ Testing framework working
- ✅ Development tools installed
- ✅ CLI skeleton functional (copies files)
- ✅ Ready to start Sprint 1

---

## Dependencies

- All tasks should be done in order (001 → 002 → 003 → 004 → 005)
- TASK-005 depends on TASK-003 (testing framework)

---

## Sprint Notes

**Why Sprint 0?**
- Establishes foundation before feature development
- Prevents rework and technical debt
- Ensures consistent development environment
- Standard practice in Agile methodology

**Common Mistakes:**
- Skipping .gitignore (bloats repository)
- Not testing the testing setup
- Forgetting to commit regularly
- Trying to add features in Sprint 0 

**Next:** Sprint 1 - Core transformation functions (tokenize, hex/bin, case)