# AGENTS.md

## Setup tips
- Run `go mod tidy` to install any missing dependencies.
- Build the project with `go build -o go-reloaded`.
- Run the program using `go run . input.txt output.txt` where `input.txt` is your source file and `output.txt` is the destination.
- Use `go fmt ./...` to format all code before committing.

## Code style
- Always run `go fmt` on files before committing to maintain consistent formatting.
- Keep functions small and focused on a single task.
- Use descriptive variable names - avoid single letters except for loop counters.
- Standard library only - don't add external packages unless absolutely necessary.
- Comment the why behind complex logic, not the what.

## Testing instructions
- Run `go test ./...` to execute all tests before committing.
- Use `go test ./... -cover` to check test coverage and aim for high coverage on transformation functions.
- Add or update tests for any logic you change.
- Test edge cases like empty input, invalid markers, and boundary conditions.
- Run tests from the project root before every commit to ensure nothing breaks.

## PR instructions
- Always run `go fmt ./...` and `go test ./...` before committing.
- Use conventional commits: `feat:` for features, `fix:` for bugs, `test:` for tests, `refactor:` for restructuring, `docs:` for documentation.
- Ensure output matches expected results exactly - 100% accuracy is required.