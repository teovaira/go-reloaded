# AI Agent Guidelines

**Project:** go-reloaded - Text transformation CLI tool  
**Current Phase:** Planning & Analysis (Week 1 - No code yet!)  
**Methodology:** TDD + Agile

---

## 📍 Where We Are

This is **Week 1: Planning Phase**. No code has been written yet. All work so far is analysis and task decomposition.

---

## 📚 Key Documents

**Start here:**
1. `docs/PROJECT-ANALYSIS.md` - Problem breakdown, architecture choice, edge cases
2. `docs/GOLDEN-TEST-SET.md` - 12 test cases that must pass
3. `docs/AGILE-ROADMAP.md` - Sprint overview and workflow

**Task details:**
- `docs/sprints/SPRINT-0-INFRASTRUCTURE.md` - Tasks 001-005
- `docs/sprints/SPRINT-1-CORE-TRANSFORMATIONS.md` - Tasks 006-014
- `docs/sprints/SPRINT-2-ADVANCED-TRANSFORMATIONS.md` - Tasks 015-019
- `docs/sprints/SPRINT-3-INTEGRATION.md` - Tasks 020-025

---

## 🎯 Project Goals

Build a CLI tool that transforms text files:
- Number conversions: `(hex)`, `(bin)`
- Case transformations: `(up)`, `(low)`, `(cap)`
- Article correction: `a` → `an`
- Punctuation spacing
- Quote pairing

**Target:** ~200-300 lines of Go code, 25 tasks, 4 sprints

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
go test -v          # Run tests
go test -cover      # Check coverage
go fmt ./...        # Format code
go vet ./...        # Check for issues
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

- **Sprint 0:** 5 tasks (setup, CLI skeleton)
- **Sprint 1:** 9 tasks (tokenize, parse, conversions, case)
- **Sprint 2:** 5 tasks (articles, punctuation, quotes, pipeline)
- **Sprint 3:** 6 tasks (integration, golden tests, docs)

**Total:** 25 tasks

---

**Note:** This file will be updated when implementation begins (Week 2).
