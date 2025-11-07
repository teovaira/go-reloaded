# go-reloaded

A text transformation CLI tool built in Go using Test-Driven Development (TDD) and Agile methodology.

**Student:** Theodore Vairaktaris  
**Institution:** Zone01 Athens  
**Status:** Implementation & Refinement (Pipeline working)  
**License:** MIT

---

## 📋 Project Overview

**go-reloaded** processes text files and applies transformation rules:
- Number conversions: hexadecimal/binary → decimal
- Case transformations: uppercase, lowercase, capitalize
- Article correction: a → an (vowels and h)
- Punctuation spacing fixes
- Quote pairing with proper marks

**Project Scope:** ~600 lines of Go code (at this stage)

---

## 🎯 Transformation Examples

**Number Conversions:**
- `42 (hex)` → `66`
- `1010 (bin)` → `10`

**Case Transformations:**
- `hello (up)` → `HELLO`
- `WORLD (low)` → `world`
- `title (cap)` → `Title`
- `hello world (up, 2)` → `HELLO WORLD`
- `MAKE THIS lower (low, 2)` → `MAKE this lower`
- `this is nice (cap, 3)` → `This Is Nice`

**Context-Aware:**
- `a apple` → `an apple`
- `a hour` → `an hour`
- `hello , world !` → `hello, world!`
- `' hi '` → `'hi'`

---

## 🏗️ Development Approach

### Methodology
- **Agile:** 4 sprints (Sprint 0-3)
- **TDD:** Write tests first, then implement
- **Incremental:** Build feature by feature

### Architecture Pattern
**Pipeline Model:** Sequential transformations

```
Input → Tokenize → Transform Pipeline (per line) → Detokenize → Output
                        ↓
            [hex/bin] → [article] → [case] → [punct] → [quotes]
```

**Transformation Order Matters (final):**
1. Number conversions (hex/bin)
2. Article correction (a → an)
3. Case transformations (up/low/cap)
4. Punctuation spacing
5. Quote pairing

Order rationale: Token-based transforms happen first; spacing and quotes run last to avoid re-introducing spaces when tokens are joined. The pipeline processes input per line to preserve newlines.

---

## 📁 Project Structure

```
go-reloaded/
├── README.md                    # This file
├── go.mod                       # Go module file
├── main.go                      # Entry point
├── docs/
│   ├── PROJECT-ANALYSIS.md      # Requirements analysis
│   ├── GOLDEN-TEST-SET.md       # Test cases
│   ├── AGILE-ROADMAP.md         # Sprint overview
│   └── sprints/
│       ├── SPRINT-0-INFRASTRUCTURE.md
│       ├── SPRINT-1-CORE-TRANSFORMATIONS.md
│       ├── SPRINT-2-ADVANCED-TRANSFORMATIONS.md
│       └── SPRINT-3-INTEGRATION.md
└── tests/
    └── testdata/                # Golden test files
```

---

## 🚀 Getting Started (For Development)

### Prerequisites
- Go 1.21+
- Basic understanding of Go syntax
- Git for version control

### Setup Instructions

**1. Clone the repository**
```bash
git clone <your-repo-url>
cd go-reloaded
```

**2. Initialize Go module**
```bash
go mod init go-reloaded
```

**3. Run tests (once implemented)**
```bash
go test -v
```

**4. Build the program**
```bash
go build -o go-reloaded
```

**5. Run the program**
```bash
./go-reloaded input.txt output.txt
```

---

## 📚 Documentation

- **[PROJECT-ANALYSIS.md](docs/PROJECT-ANALYSIS.md)** - Detailed requirements and architecture
- **[GOLDEN-TEST-SET.md](docs/GOLDEN-TEST-SET.md)** - Complete test specifications
- **[AGILE-ROADMAP.md](docs/AGILE-ROADMAP.md)** - Sprint overview and workflow
- **Sprint Files** - Task breakdowns in `docs/sprints/`

Notes on behavior (as implemented):
- Words include both alphabetic tokens and decimal numbers; punctuation is tokenized separately.
- Invalid markers (e.g., `(up, )`, `(low, -1)`) are ignored and kept literal.
- Capitalization is Unicode-aware.
- Newlines are preserved; spaces are normalized within each line.

---

## 🧪 Testing Strategy

### Test Types
1. **Unit Tests:** Individual transformation functions
2. **Integration Tests:** Full pipeline testing
3. **Golden Tests:** Expected output files
4. **Edge Cases:** Unicode, errors, boundaries

### Running Tests
```bash
# All tests
go test -v

# With coverage
go test -cover

# Specific test
go test -run TestTokenize
```

---

## 📖 For Other Students

### Want to Review or Learn from This Project?

**To Run Locally:**
1. Clone this repository
2. Run `go mod tidy`
3. Run `go test -v` to see test results
4. Review the planning docs in `/docs`

**Found a Bug or Improvement?**
- Open an issue describing the problem
- Submit a pull request with a fix
- Reference the relevant task from sprint docs

**Learning from This Project?**
- Check the sprint files for incremental learning path
- Each task follows TDD (test → implement → refactor)
- Resources included for key Go concepts

### For Zone01 Peers
This project follows the Zone01 Agile methodology:
- **Week 1:** Planning and analysis (no code written)
- **Weeks 2-3:** Implementation with TDD
- All decisions documented in `/docs`

---

## 🔄 Development Workflow

### TDD Cycle
1. **🔴 RED:** Write failing test
2. **🟢 GREEN:** Make test pass with minimal code
3. **♻️ REFACTOR:** Improve code quality
4. **✅ COMMIT:** Save your progress

See [AGILE-ROADMAP.md](docs/AGILE-ROADMAP.md) for detailed sprint breakdown.

---

## ✅ Project Completion Criteria

The project is complete when:
- ✅ All 12 golden test cases pass
- ✅ Code coverage ≥ 90%
- ✅ All sprint tasks completed
- ✅ Documentation is complete
- ✅ Code follows Go best practices
- ✅ No critical bugs

---

## 📝 License

MIT License - Feel free to use this project for learning purposes.

---

## 🙏 Acknowledgments

- Zone01 Athens for the project specification
- Go community for excellent documentation
- Peers and mentors for code reviews

---

## 📧 Contact

**Theodore Vairaktaris**  
Zone01 Athens Student  
[GitHub Profile](https://github.com/teovaira) | [Email](mailto:t.vairaktaris@gmail.com)

---

**Status:** Completed 🚀
