# Go-Reloaded: Agile Development Roadmap

**Author:** Theodore Vairaktaris  
**Project:** go-reloaded  
**Date:** October 2025  
**Phase:** Planning & Task Decomposition (No Coding Yet!)

---

## 🎯 Purpose

This roadmap breaks down the go-reloaded project into **small, testable tasks** following Agile and TDD principles. Each task describes **WHAT** to build and test, not **HOW** to code it.

---

## 📚 Foundation Documents

This roadmap references:
- [`PROJECT-ANALYSIS.md`](./PROJECT-ANALYSIS.md) - Problem breakdown and rules
- [`GOLDEN-TEST-SET.md`](./GOLDEN-TEST-SET.md) - 12 test scenarios

**Flow:** Analysis → Planning → Implementation → Validation

---

## 🏗️ Architecture Overview

**Pattern:** Pipeline (sequential transformations)  
**Scope:** ~200-300 lines of Go code  
**Structure:** Simple (main.go + test files in root)

**Pipeline Flow:**
```
Read File → Tokenize → Transform (7 stages) → Detokenize → Write File
```

**Transformation Order:**
1. Hex/Binary conversion
2. Case changes (up, low, cap)
3. Article correction (a→an)
4. Punctuation spacing
5. Quote formatting

---

## 🗓️ Sprint Breakdown

### Sprint 0: Foundation (1-2 days, 4 tasks)
**Goal:** Get input/output working, create pipeline skeleton  
**Deliverable:** Program reads file, passes through (no transformations), writes file

### Sprint 1: Core Transformations (3-4 days, 7 tasks)
**Goal:** Implement (hex), (bin), (up), (low), (cap)  
**Deliverable:** 5 core transformations working independently

### Sprint 2: Advanced Rules (3-4 days, 5 tasks)
**Goal:** Implement (a→an), punctuation, quotes  
**Deliverable:** All 7 transformation types complete

### Sprint 3: Integration (2-3 days, 6 tasks)
**Goal:** All 12 golden tests passing, production-ready  
**Deliverable:** Audit-ready project with full documentation

**Total:** 22 tasks, ~10-12 days

---

## 🧪 TDD Workflow

Every task follows this cycle:

### 1️⃣ RED - Write Failing Test
- Define expected behavior
- Create test case
- Run test → FAIL

### 2️⃣ GREEN - Make It Pass
- Write minimum code to pass
- No extra features yet
- Run test → PASS

### 3️⃣ REFACTOR - Improve
- Clean up code
- Remove duplication
- Run test → still PASS

---

## 📋 Task Structure

Each task in sprint files includes:

| Section | Purpose |
|---------|---------|
| **Learning Objectives** | What concepts you'll learn |
| **Test Scenarios** | What cases to test (no code!) |
| **Implementation Goal** | What to build |
| **Edge Cases** | What to watch for |
| **Acceptance Criteria** | How to know you're done |
| **AI Guidance** | What questions to ask AI |
| **Resources** | Links to learn concepts |

---

## 🎓 Learning Path

### Week 1: Foundation
- Go file I/O
- Command-line args
- String tokenization
- Testing basics

### Week 2: Transformations
- Number base conversion
- String manipulation
- Pattern matching
- State management (quotes)

### Week 3: Polish
- Integration testing
- Error handling
- Code refactoring
- Documentation

---

## ✅ Definition of Done

Project is complete when:
- [ ] All 12 golden tests pass
- [ ] Program handles invalid input gracefully
- [ ] Code is ~200-300 lines (concise!)
- [ ] Tests are comprehensive
- [ ] Documentation is clear
- [ ] Ready for peer audit

---

## 🚀 How to Use This Roadmap

1. **Read this overview** (you are here!)
2. **Review PROJECT-ANALYSIS.md** to understand the problem
3. **Skim GOLDEN-TEST-SET.md** to see expected outcomes
4. **Open SPRINT-0-INFRASTRUCTURE.md**
5. **Start with TASK-001**
6. **Follow TDD cycle** for each task
7. **Ask AI agents** when stuck (see AI Guidance sections)
8. **Mark tasks complete** as you finish

---

## 🗂️ Sprint Files

| Sprint | File | Tasks | Focus |
|--------|------|-------|-------|
| 0 | [`SPRINT-0-INFRASTRUCTURE.md`](./sprints/SPRINT-0-INFRASTRUCTURE.md) | 4 | Setup, I/O, skeleton |
| 1 | [`SPRINT-1-CORE-TRANSFORMATIONS.md`](./sprints/SPRINT-1-CORE-TRANSFORMATIONS.md) | 7 | hex, bin, up, low, cap |
| 2 | [`SPRINT-2-ADVANCED-TRANSFORMATIONS.md`](./sprints/SPRINT-2-ADVANCED-TRANSFORMATIONS.md) | 5 | a→an, punctuation, quotes |
| 3 | [`SPRINT-3-INTEGRATION.md`](./sprints/SPRINT-3-INTEGRATION.md) | 6 | Testing, polish, audit prep |

---

## 🤖 AI Agent Usage

**When to use AI:**
- Explain Go concepts (not write code for you!)
- Review your test design
- Suggest edge cases to test
- Help debug when tests fail
- Explain error messages

**Example questions:**
- "What's the best way to tokenize a string in Go?"
- "How do I test file I/O without actual files?"
- "What edge cases should I consider for hex conversion?"

**Don't ask AI to:**
- Write the full implementation
- Do the project for you
- Skip the TDD process

---

## 📊 Progress Tracking

| Sprint | Status | Tests Passing |
|--------|--------|---------------|
| Sprint 0 | ⏳ | - |
| Sprint 1 | ⏳ | 0/5 |
| Sprint 2 | ⏳ | 0/4 |
| Sprint 3 | ⏳ | 0/12 |

---

## 💡 Key Principles

**Keep It Simple:**
- One function per transformation
- Clear test names
- Descriptive variable names

**Test First:**
- Write test before code
- One failing test at a time
- Tests define behavior

**Iterate:**
- Get it working first
- Then make it clean
- Don't over-engineer

---

## 🎯 Success Metrics

You'll know you succeeded when:
- ✅ All golden tests pass
- ✅ Code is readable and maintainable
- ✅ You can explain every decision
- ✅ Audit passes on first try
- ✅ You learned Go and TDD!

---

**Next Step:** Open [`SPRINT-0-INFRASTRUCTURE.md`](./sprints/SPRINT-0-INFRASTRUCTURE.md) 🚀