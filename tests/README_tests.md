# go-reloaded Tests

This folder contains a simple golden test runner.

How to run
- From the `tests/` directory: `bash run_tests.sh`
- The script runs the program on `sample.txt` and `golden.txt` test files.
- It compares outputs with expected results using `diff -w` (whitespace-insensitive).

Notes
- The program preserves newlines and normalizes spaces within each line.
- Update `result_expected.txt` or `golden_expected.txt` when intentional changes are made to the pipeline output.
- The script must be run from the `tests/` directory (not from the project root).
- Additional note: If needed, a local Go build cache can be set with `GOCACHE="$(pwd)/.gocache"`.
