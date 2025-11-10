# Test Commands

.PHONY: test test-verbose test-coverage test-unit test-integration clean-test

# Run all tests
test:
	go test ./...

# Run tests with verbose output
test-verbose:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run only unit tests (usecases and controllers)
test-unit:
	go test -v ./tests/usecases/...
	go test -v ./tests/controllers/...

# Run only integration tests (repositories)
test-integration:
	go test -v ./tests/repositories/...

# Clean test files
clean-test:
	rm -f coverage.out coverage.html
	rm -f test.db

# Run tests in watch mode (requires air or similar tool)
test-watch:
	air -c .air-test.toml

# Benchmark tests
test-bench:
	go test -bench=. -benchmem ./...

# Test with race detection
test-race:
	go test -race ./...