# Go-Reloaded: Agile Development Roadmap

**Author:** Theodore Vairaktaris
**Project:** go-reloaded
**Date:** October 2025
**Phase:** Implementation & Refinement

> **📝 Implementation Note:** This roadmap was created during the planning phase. The final implementation evolved to use a more professional modular structure with `internal/` packages (543 lines vs estimated 200-300 lines), but the core transformation order and pipeline logic remained exactly as planned. See "Post-Implementation Retrospective" at the end of this document for a detailed comparison of planned vs actual implementation.

---

## 🎯 Purpose

This roadmap breaks down the go-reloaded project into **25 testable tasks** following Agile and TDD principles. Each task describes **WHAT** to build and test, not **HOW** to code it.

This is a **planning document** created during the analysis phase before any code is written.

---

## 📚 Foundation Documents

This roadmap references:
- [`PROJECT-ANALYSIS.md`](./PROJECT-ANALYSIS.md) - Problem breakdown and transformation rules
- [`GOLDEN-TEST-SET.md`](./GOLDEN-TEST-SET.md) - 12 test scenarios with expected outputs

**Document Flow:**
```
Analysis → Planning (you are here) → Implementation → Validation
```

---

## 🏗️ Architecture Overview

**Pattern:** Pipeline (sequential transformations)  
**Estimated Scope:** ~200-300 lines of Go code (rough guideline, not a requirement)  
**Structure:** Simple (main.go + main_test.go in root)

**Pipeline Flow (per line):**
```
Read File → split by lines → Tokenize → Transform (hex/bin → articles → case) → Join → Punctuation → Quotes → rejoin lines → Write File
```

**Transformation Order (final, critical):**
1. Number conversions (hex, binary)
2. Article correction (a→an)
3. Case transformations (up, low, cap)
4. Punctuation spacing (no space before; one after)
5. Quote formatting (trim inside; pair quotes)

**Why this order?**
- Article correction must happen BEFORE case transformations (so "a amazing (up, 3)" becomes "AN AMAZING" not "An AMAZING").
- Token-based transforms must happen BEFORE spacing.
- Punctuation and quotes MUST run LAST to avoid reintroducing spaces when tokens are joined.
- The pipeline runs per line to preserve newlines.

---

## 🗓️ Sprint Breakdown

### Sprint 0: Foundation (1-2 days, 5 tasks)
**Goal:** Build the skeleton - file I/O and project setup  
**Deliverable:** Program reads file, passes through unchanged, writes file

### Sprint 1: Core Transformations (4-5 days, 10 tasks)
**Goal:** Implement tokenization, command parsing, number conversions, case transformations with integration  
**Deliverable:** Core transformations working and integrated with command detection

### Sprint 2: Advanced Transformations & Pipeline Integration (3-4 days, 4 tasks)
**Goal:** Implement articles, punctuation, quotes, and connect all transformations into complete pipeline  
**Deliverable:** All transformations complete and integrated into working pipeline

### Sprint 3: End-to-End Testing & Documentation (2-3 days, 6 tasks)
**Goal:** Validate complete system with golden tests and finalize documentation  
**Deliverable:** Production-ready project with comprehensive testing

**Total:** 25 tasks, ~10-12 days

---

## 🧪 TDD Workflow

Every task follows this cycle:

### 🔴 RED - Write Failing Test
- Define expected behavior
- Create test case that describes what SHOULD happen
- Run test → FAIL (function doesn't exist yet)

### 🟢 GREEN - Make It Pass
- Write MINIMUM code to pass the test
- No extra features or optimization yet
- Run test → PASS

### ♻️ REFACTOR - Improve Quality
- Clean up code (remove duplication, improve names)
- Make it more readable
- Run test → still PASS

**Mantra:** Red → Green → Refactor → Repeat

---

## 📋 Task Structure

Each task in sprint files includes:

| Section | Purpose |
|---------|---------|
| **Story Points** | Complexity estimate (1-5) |
| **Learning Objectives** | What concepts you'll learn |
| **What to Build** | Clear description of functionality |
| **TDD Cycle** | RED-GREEN-REFACTOR steps |
| **Test Scenarios** | What cases to test (input → expected output) |
| **Edge Cases** | Boundary conditions to handle |
| **Architecture Decisions** | Design choices to make |
| **Acceptance Criteria** | How to know you're done |
| **AI Guidance** | Questions to ask AI agents |
| **Resources** | Links to learn concepts |

---

## 🎓 Learning Path

### Sprint 0-1: Foundation & Core
- Go file I/O (`os` package)
- Command-line arguments (`os.Args`)
- String tokenization (`strings` package)
- Testing basics (`testing` package)
- Error handling patterns
- Number base conversion (`strconv`)
- String manipulation (ToUpper, ToLower, runes)

### Sprint 2: Advanced Transformations
- Regular expressions (`regexp`)
- Slice operations
- Pipeline composition
- Complex string processing

### Sprint 3: Integration & Polish
- End-to-end testing (testing the complete program flow from input file to output file)
- Error handling strategies
- Code refactoring (DRY principle)
- Documentation writing
- Code review practices

**Timeline:** Flexible - typically 2-3 weeks depending on your pace and experience

---

## ✅ Definition of Done

Project is complete when:
- [ ] All 12 golden tests pass
- [ ] Program handles invalid input gracefully (no crashes)
- [ ] Test coverage >85%
- [ ] Documentation is clear
- [ ] Code is readable and maintainable
- [ ] Ready for professional code review
 - [ ] Newlines preserved; spaces normalized within lines

---

## 🚀 How to Use This Roadmap

### Step-by-Step Process

1. **Read this overview** (you are here!)
2. **Review PROJECT-ANALYSIS.md** to understand the problem deeply
3. **Skim GOLDEN-TEST-SET.md** (read quickly to get overview) to see expected outcomes
4. **Open SPRINT-0-INFRASTRUCTURE.md**
5. **Start with TASK-001**
6. **For each task:**
   - Read all sections carefully
   - Understand what to build (don't code yet!)
   - Think about test scenarios
   - Consider edge cases
   - Review architecture decisions
7. **Begin TDD cycle:**
   - 🔴 Write failing test
   - 🟢 Make it pass
   - ♻️ Refactor
8. **Validate against acceptance criteria** (run tests to confirm they pass)
9. **Move to next task**

### When You Get Stuck

1. **Review the task's "Learning Objectives"** - what concept are you missing?
2. **Check "Resources"** - read the linked documentation
3. **Use AI Guidance** - ask the suggested questions to AI agents
4. **Review similar code** - look at previous tasks for patterns
5. **Take a break** - sometimes stepping away helps!

---

## 🗂️ Sprint Files

Detailed task breakdowns:

| Sprint | File | Tasks | Focus |
|--------|------|-------|-------|
| 0 | [`SPRINT-0-INFRASTRUCTURE.md`](./sprints/SPRINT-0-INFRASTRUCTURE.md) | 5 | Setup, I/O, skeleton |
| 1 | [`SPRINT-1-CORE-TRANSFORMATIONS.md`](./sprints/SPRINT-1-CORE-TRANSFORMATIONS.md) | 10 | Tokenize, parse, hex, bin, case + integration |
| 2 | [`SPRINT-2-ADVANCED-TRANSFORMATIONS.md`](./sprints/SPRINT-2-ADVANCED-TRANSFORMATIONS.md) | 4 | Articles, punctuation, quotes, full pipeline |
| 3 | [`SPRINT-3-INTEGRATION.md`](./sprints/SPRINT-3-INTEGRATION.md) | 6 | System testing, golden tests, documentation |

---

## 🤖 AI Agent Usage

### When to Use AI Agents

**Good uses:**
- ✅ Explain Go concepts ("How does strconv.ParseInt work?")
- ✅ Review your test design ("Are these test cases comprehensive?")
- ✅ Suggest edge cases ("What edge cases should I test?")
- ✅ Help debug when tests fail ("Why is this test failing?")
- ✅ Explain error messages ("What does this compiler error mean?")
- ✅ Review your code ("Is this idiomatic Go?")

**Bad uses:**
- ❌ Write the full implementation for you
- ❌ Do the project for you
- ❌ Skip the TDD process
- ❌ Generate code you don't understand

### Example Questions

**Concept Learning:**
- "Explain hexadecimal number system and how to convert to decimal"
- "What's the difference between bytes and runes in Go?"
- "How do I tokenize a string by whitespace in Go?"

**Testing:**
- "How do I write table-driven tests in Go?"
- "What edge cases should I test for hex conversion?"
- "How do I test file I/O without actual files?"

**Debugging:**
- "My test is failing with this error: [paste error]. What does it mean?"
- "Why does ParseInt return two values in Go?"
- "How do I check if a slice index is valid before accessing it?"

**Code Review:**
- "Is this Go code idiomatic? [paste code]"
- "How can I simplify this function?"
- "Is there duplication I should refactor?"

---

## 📊 Progress Tracking

Mark tasks as you complete them:

| Sprint | Status | Tests Passing | Progress |
|--------|--------|---------------|----------|
| Sprint 0 | ⏳ Not Started | - | 0/5 tasks |
| Sprint 1 | ⏳ Not Started | - | 0/10 tasks |
| Sprint 2 | ⏳ Not Started | - | 0/4 tasks |
| Sprint 3 | ⏳ Not Started | 0/12 golden | 0/6 tasks |

**Symbols:** ⏳ Not Started | 🔄 In Progress | ✅ Complete

---

## 💡 Key Principles

### Keep It Simple
- Start with the simplest solution that works
- Don't over-engineer
- One function per transformation
- Clear naming (no abbreviations)

### Test First, Always
- Write test before code (RED)
- Make test pass (GREEN)
- Then improve (REFACTOR)
- Never skip TDD cycle

### Iterate and Improve
- Get it working first
- Then make it clean
- Then make it fast (if needed)
- "Make it work, make it right, make it fast" - Kent Beck

### Small Steps
- One failing test at a time
- Commit after each task
- Don't try to do everything at once

---

## 🎯 Success Metrics

You'll know you succeeded when:
- ✅ All 12 golden tests pass
- ✅ Code is concise, readable, and maintainable
- ✅ You can explain every design decision
- ✅ Tests give you confidence to refactor
- ✅ You learned Go and TDD!
- ✅ Code passes professional review

---

## 🔧 Tools & Commands

**Essential Go commands:**
```bash
go mod init go-reloaded    # Initialize project
go build                    # Compile program
go test                     # Run tests
go test -v                  # Verbose test output
go test -cover              # Show coverage
go fmt                      # Format code
go vet                      # Find potential bugs
```

**Recommended workflow:**
```bash
# After each task:
go test -v                  # All tests pass?
go fmt                      # Format code
go vet                      # No warnings?
git add .                   # Stage changes
git commit -m "TASK-XXX"   # Commit with task number
```

---

## 📖 Additional Context

### Why This Structure?

**4 Sprints:** Industry-standard sprint length (1-2 weeks each)  
**25 Tasks:** Small enough to complete in 1-4 hours each  
**TDD:** Industry best practice, ensures testable code  
**Pipeline Pattern:** Simple, maintainable architecture

### Real-World Application

These skills transfer directly to professional software development:
- Sprint planning → Used in every Agile team
- TDD → Standard practice at major tech companies
- Code review → Daily activity for professional developers
- Documentation → Critical for team collaboration

---

**Next Step:** Open [`SPRINT-0-INFRASTRUCTURE.md`](./sprints/SPRINT-0-INFRASTRUCTURE.md) and begin! 🚀

---

## 📝 Post-Implementation Retrospective

**Status:** ✅ Completed
**Implementation Date:** October-November 2025
**Retrospective Date:** November 2025

### Executive Summary

The project successfully met all functional requirements and passed all 12 golden test cases. During implementation, the architecture evolved from the planned simple structure to a more professional modular design, while maintaining the exact transformation order and pipeline logic specified in the planning phase.

---

### What Changed During Implementation

#### 1. Architecture Evolution

**Planned (from this roadmap):**
- Simple flat structure: `main.go` + `main_test.go` in project root
- All code in a single file or minimal files
- Estimated 200-300 lines total

**Actually Implemented:**
- Modular structure with `internal/` packages:
  - `internal/fileio/` - File I/O operations
  - `internal/pipeline/` - Main processing pipeline
  - `internal/tokenizer/` - Text tokenization
  - `internal/transform/` - All transformation functions (5 files)
- Total: 543 lines of production code
- Separation of concerns with clear package boundaries

**Why the Change:**
- Learned Go best practices during implementation
- `internal/` package pattern prevents external imports (good encapsulation)
- Each transformation in its own file improves maintainability
- Better code organization for future extensions
- More professional structure suitable for portfolio/code review

#### 2. Code Size

**Planned:** ~200-300 lines
**Actual:** 543 lines (81% increase)

**Reasons for Increase:**
- More comprehensive error handling
- Better code documentation and comments
- Unicode-aware string processing (runes vs bytes)
- Robust edge case handling
- Clearer function separation (single responsibility principle)

---

### What Stayed the Same ✅

#### Transformation Order (100% Match)
**Planned and Implemented:**
1. Number conversions (hex/bin) → `ConvertHexAndBin()`
2. Article correction (a→an) → `FixArticles()`
3. Case transformations → `ApplyCaseRules()`
4. Punctuation spacing → `ApplyPunctuationRules()`
5. Quote pairing → `FixQuotes()`

This order proved **critical** and worked perfectly as designed.

#### Pipeline Architecture
- Sequential transformation pipeline: ✅ Implemented exactly as planned
- Per-line processing to preserve newlines: ✅ Works perfectly
- Tokenization before transformations: ✅ Core design pattern maintained

#### Testing Approach
- TDD methodology: ✅ Followed throughout
- All 12 golden tests: ✅ Pass
- Edge case handling: ✅ Comprehensive

---

### Key Metrics

| Metric | Planned | Actual | Status |
|--------|---------|--------|--------|
| Lines of Code | 200-300 | 543 | ⚠️ Larger (better quality) |
| Golden Tests Passing | 12/12 | 12/12 | ✅ Perfect |
| Architecture | Flat | Modular | ⚠️ Improved |
| Transformation Order | As specified | As specified | ✅ Perfect |
| Sprint Count | 4 | 4 | ✅ Followed plan |
| Task Count | 25 | ~25 | ✅ Approximately followed |

---

### Lessons Learned

#### As a Junior Developer

1. **Initial Estimates Are Guidelines, Not Limits**
   - Estimated 200-300 lines, implemented 543
   - More code ≠ worse code (if it's cleaner and better organized)
   - Quality and maintainability matter more than hitting estimates

2. **Architecture Can Evolve During Implementation**
   - Started planning for simple structure
   - Learned about Go's `internal/` package pattern during implementation
   - Made conscious decision to adopt better practices
   - **Key insight:** It's OK to deviate from the plan when you learn better approaches

3. **Planning Still Provided Value**
   - Even though structure changed, the transformation order plan was **perfect**
   - Sprint breakdown helped organize work
   - TDD approach worked exactly as planned
   - Having a plan made it easier to know when to deviate intelligently

4. **Core Logic vs Implementation Details**
   - **Core logic** (transformation order, pipeline design): Stayed exactly as planned ✅
   - **Implementation details** (file structure, package organization): Evolved and improved ⚠️
   - This is a **healthy pattern** in software development

#### Technical Insights

1. **Go Best Practices Matter**
   - `internal/` packages prevent accidental external dependencies
   - Separation of concerns makes testing easier
   - Single responsibility principle (one file per transformation type) improves clarity

2. **Transformation Order Was Critical**
   - Articles before case transformations: **Essential** (prevents "An APPLE" instead of "AN APPLE")
   - Token transforms before spacing: **Necessary** to avoid reintroducing spaces
   - Quotes last: **Required** to handle all prior transformations correctly
   - The planning phase got this **100% correct**

3. **Unicode Complexity**
   - Runes vs bytes distinction required more code than anticipated
   - But resulted in proper international character support
   - Worth the extra complexity

---

### What Would I Do Differently?

#### If Starting Over

1. **Better Initial Estimate**
   - Account for error handling and edge cases (doubles the code)
   - Factor in documentation and comments (adds 20-30%)
   - More realistic: 400-600 lines for a robust implementation

2. **Start with Modular Structure**
   - Could have planned for `internal/` packages from the beginning
   - Would have saved refactoring time

3. **More Edge Case Analysis Upfront**
   - Empty files, invalid markers, Unicode characters
   - Would have caught these in planning phase

#### What Worked Well (Keep Doing)

1. **Transformation Order Analysis** ✅
   - Got it right the first time
   - No reordering needed during implementation

2. **TDD Approach** ✅
   - Caught bugs early
   - Gave confidence to refactor

3. **Golden Test Set** ✅
   - Perfect acceptance criteria
   - Knew exactly when the project was "done"

4. **Sprint Structure** ✅
   - Logical progression from simple to complex
   - Each sprint built on previous work

---

### Impact on Project Success

**Did the deviations hurt the project?**
**No** - They **improved** it:
- ✅ More maintainable code
- ✅ Better separation of concerns
- ✅ Easier to test
- ✅ More professional structure
- ✅ All requirements met
- ✅ All tests passing

**Was the planning phase wasted?**
**Absolutely not:**
- Transformation order plan was **perfect**
- Pipeline design worked **exactly as specified**
- Sprint structure organized the work **effectively**
- TDD approach was followed **successfully**
- The plan provided a **foundation** to build on and improve

---

### Final Verdict

**Planning Phase:** ⭐⭐⭐⭐⭐ (5/5)
- Got the critical parts (transformation order, pipeline design) **perfect**
- Provided clear roadmap and learning path
- TDD approach worked excellently

**Implementation Phase:** ⭐⭐⭐⭐⭐ (5/5)
- All functional requirements met
- All 12 golden tests pass
- Professional code structure
- Good documentation

**Agile Process:** ⭐⭐⭐⭐⭐ (5/5)
- Sprints provided structure
- Adapted when learning new best practices
- Delivered working software
- This retrospective captures the learning

---

### Recommendations for Future Projects

1. **Keep Using Agile Planning**
   - Roadmaps are valuable even if implementation evolves
   - Transformation order analysis is **critical** - spend time on this

2. **Start with Better Structure**
   - Use `internal/` packages from day one for Go projects
   - Plan for 2-3x your initial estimate for robust code

3. **Document Evolution**
   - Retrospectives like this one show professional growth
   - Explain why you made changes (learning, best practices)

4. **Trust the Plan for Core Logic**
   - Transformation order: **Plan it carefully, then stick to it**
   - Implementation details: **Adapt as you learn**

---

### Conclusion

This project demonstrates healthy software development:
- **Planned carefully** (transformation order, pipeline design)
- **Implemented professionally** (modular structure, clean code)
- **Adapted intelligently** (evolved architecture based on learning)
- **Delivered successfully** (all tests pass, requirements met)

The deviations from the roadmap represent **growth and learning**, not failure. The core architectural decisions (transformation order, pipeline pattern) were correct from the start. The implementation details (file structure, package organization) improved as I learned Go best practices during development.

**This is exactly how Agile is supposed to work.** ✅

---

*"The journey of a thousand lines begins with a single test."* — Adapted from Lao Tzu
