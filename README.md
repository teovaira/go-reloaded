# go-reloaded

A text transformation CLI tool built in Go using Test-Driven Development and Agile methodology.

**Student:** Theodore Vairaktaris  
**Institution:** Zone01 Athens  
**Project:** First assignment - Text processing pipeline

---

## 📋 Overview

**go-reloaded** processes text files applying transformation rules:
- Number conversions (hex/binary → decimal)
- Case changes (uppercase, lowercase, capitalize)
- Article correction (a → an)
- Punctuation spacing
- Quote formatting

**Usage:**
```bash
go run . input.txt output.txt
```

---

## 🔧 Transformation Rules

### Number Conversions
- `1E (hex)` → `30`
- `10 (bin)` → `2`

### Case Transformations
- `word (up)` → `WORD`
- `WORD (low)` → `word`
- `word (cap)` → `Word`
- `words (up, 3)` → transforms 3 previous words

### Linguistic Rules
- `a apple` → `an apple`
- `Hello , world !` → `Hello, world!`
- `' hello '` → `'hello'`

See [`docs/PROJECT-ANALYSIS.md`](docs/PROJECT-ANALYSIS.md) for complete specifications.

---

## 🏗️ Project Structure

```
go-reloaded/
├── main.go                 # Main application (~200-300 lines)
├── main_test.go            # Test suite
├── go.mod
├── README.md
├── AGENTS.md               # AI agent guidelines
│
└── docs/                   # Planning & documentation
    ├── PROJECT-ANALYSIS.md
    ├── GOLDEN-TEST-SET.md
    ├── AGILE-ROADMAP.md
    └── sprints/
        ├── SPRINT-0-INFRASTRUCTURE.md
        ├── SPRINT-1-CORE-TRANSFORMATIONS.md
        ├── SPRINT-2-ADVANCED-TRANSFORMATIONS.md
        └── SPRINT-3-INTEGRATION.md
```

---

## 🚀 Quick Start

### Build & Run
```bash
# Build
go build

# Run
./go-reloaded input.txt output.txt

# Or directly
go run . input.txt output.txt
```

### Testing
```bash
# Run all tests
go test -v

# With coverage
go test -cover

# Specific test
go test -run TestGoldenCase1
```

---

## 📖 Examples

### Example 1: Number Conversions
**Input:**
```
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```
**Output:**
```
Simply add 66 and 2 and you will see the result is 68.
```

### Example 2: Multiple Rules
**Input:**
```
here (cap) is a interesting text with 1A (hex) items , all in ' a epic document (cap, 2) ' ... what do you think (up, 4) ?
```
**Output:**
```
Here is an interesting text with 26 items, all in 'an Epic Document'... WHAT DO YOU THINK?
```

---

## 🧪 Testing

**Test Cases:** 12 golden test scenarios  
**Coverage:** >85% target  
**Approach:** Test-Driven Development (TDD)

All test scenarios documented in [`docs/GOLDEN-TEST-SET.md`](docs/GOLDEN-TEST-SET.md)

---

## 📚 Development Process

This project follows **Agile methodology** with 4 sprints:

| Sprint | Focus | Tasks |
|--------|-------|-------|
| 0 | Setup & I/O | 4 |
| 1 | Core transformations | 7 |
| 2 | Advanced rules | 5 |
| 3 | Polish & audit prep | 6 |

See [`docs/AGILE-ROADMAP.md`](docs/AGILE-ROADMAP.md) for complete task breakdown.

---

## 🎯 Key Features

- ✅ Pipeline architecture (sequential transformations)
- ✅ Test-driven development
- ✅ Handles all edge cases gracefully
- ✅ ~200-300 lines of clean Go code
- ✅ Zero external dependencies
- ✅ Comprehensive documentation

---

## 🤝 For Auditors

### How to Audit

1. **Clone and build:**
   ```bash
   git clone <repo>
   cd go-reloaded
   go build
   ```

2. **Run tests:**
   ```bash
   go test -v
   ```
   Expected: All tests pass

3. **Test golden cases:**
   Use test files from `docs/GOLDEN-TEST-SET.md`

4. **Review:**
   - Code clarity and organization
   - Test coverage
   - Error handling
   - Documentation quality

---

## 📄 License

MIT License - See [LICENSE](LICENSE)

---

## 🙏 Acknowledgments

- Zone01 Athens for project specifications
- Go community for excellent documentation
- Fellow students for peer reviews

---

## 🎓 Learning Outcomes

**Technical:**
- Go programming fundamentals
- Test-Driven Development (TDD)
- Pipeline architecture pattern
- String manipulation and parsing
- Error handling best practices

**Process:**
- Agile sprint planning
- Incremental development
- Code refactoring
- Documentation standards
- AI-assisted development

---

## 📞 Contact

**Theodore Vairaktaris**  
Zone01 Athens  
go-reloaded project

---

**Built with TDD and Agile methodology** ✅