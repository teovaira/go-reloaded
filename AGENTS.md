# AI Agent Guidelines

**Project:** go-reloaded - Text transformation CLI tool  
**Current Phase:** Implementation & Refinement (Pipeline working)  
**Methodology:** TDD + Agile

---

## 📍 Where We Are

This project is **Implementation Complete**. All core functionality has been implemented, tested, and is passing golden tests. The codebase is now in refinement phase with focus on code quality improvements and comprehensive unit testing.

---

## 📚 Key Documents

**Start here:**
1. [PROJECT-ANALYSIS.md](docs/PROJECT-ANALYSIS.md) - Problem breakdown, architecture choice, edge cases
2. [GOLDEN-TEST-SET.md](docs/GOLDEN-TEST-SET.md) - 12 test cases that must pass
3. [AGILE-ROADMAP.md](docs/AGILE-ROADMAP.md) - Sprint overview and workflow
4. [README.md](README.md) - Project overview and usage

**Task details:**
- [SPRINT-0-INFRASTRUCTURE.md](docs/sprints/SPRINT-0-INFRASTRUCTURE.md) - Tasks 001-005
- [SPRINT-1-CORE-TRANSFORMATIONS.md](docs/sprints/SPRINT-1-CORE-TRANSFORMATIONS.md) - Tasks 006-014
- [SPRINT-2-ADVANCED-TRANSFORMATIONS.md](docs/sprints/SPRINT-2-ADVANCED-TRANSFORMATIONS.md) - Tasks 015-019
- [SPRINT-3-INTEGRATION.md](docs/sprints/SPRINT-3-INTEGRATION.md) - Tasks 020-025

---

## 🎯 Project Goals

Build a CLI tool that transforms text files:
- Number conversions: `(hex)`, `(bin)`
- Case transformations: `(up)`, `(low)`, `(cap)`
- Article correction: `a` → `an`
- Punctuation spacing
- Quote pairing

**Achieved:** 653 lines of Go code (initial estimate was 200–300), 25 tasks completed, 4 sprints completed

---

## 🔴🟢♻️ TDD Workflow

Every task follows: **RED** (write test) → **GREEN** (pass test) → **REFACTOR** (clean code)

---

## 🛠️ Essential Commands

**Setup (Sprint 0):**
```bash
go mod init go-reloaded
go build
go test -v
```

**Daily workflow:**
```bash
go test -v ./...        # Run all tests (all packages)
go test -cover ./...    # Check coverage (all packages)
go fmt ./...            # Format code
go vet ./...            # Check for issues
```

**Running the program:**
```bash
go run . input.txt output.txt
./go-reloaded input.txt output.txt  # After build
```

**Commit after each task:**
```bash
git add .
git commit -m "feat: TASK-XXX description"
```

---

## 🤖 How to Help

**Good questions:**
- "What test cases should I write for [feature]?"
- "What edge cases am I missing?"
- "Is this idiomatic Go?"
- "Why is my test failing?"

**Bad questions:**
- "Write the code for me"
- "Do my project"

---

## 📊 Progress

- **Sprint 0:** ✅ 5 tasks completed (setup, CLI skeleton)
- **Sprint 1:** ✅ 9 tasks completed (tokenize, parse, conversions, case)
- **Sprint 2:** ✅ 5 tasks completed (articles, punctuation, quotes, pipeline)
- **Sprint 3:** ✅ 6 tasks completed (integration, golden tests, docs)

**Total:** 25/25 tasks complete

**Current Phase:** Code quality improvements and unit test expansion

---

**Last Updated:** November 7, 2025
