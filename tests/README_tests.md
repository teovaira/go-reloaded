# go-reloaded Tests

This folder contains a simple golden test runner.

How to run
- From the `tests/` directory: `bash run_tests.sh`
- The script runs the program on `../sample.txt` and writes `../result.txt`.
- It then compares `../result.txt` with `tests/result_expected.txt` using `diff -w` (whitespace-insensitive).

Notes
- The program preserves newlines and normalizes spaces within each line.
- Update `result_expected.txt` when intentional changes are made to the pipeline output.
- Use a local Go build cache under `tests/.gocache` if needed: `GOCACHE="$(pwd)/.gocache"`.
