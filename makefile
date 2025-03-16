.PHONY: test run-api build-api clean

API_EXECUTABLE=cmd/api/main.go
API_BINARY=api
COVERAGE_THRESHOLD=80
IGNORED_PACKAGE="constant|model|service"
COVERAGE_FILE="coverage.out"
FILTERED_COVERAGE_FILE="coverage.filtered.out"

run-api:
	@echo "Starting the API..."
	@go run $(API_EXECUTABLE)

build-api:
	@echo "Building the API..."
	@go build -o $(API_BINARY) $(API_EXECUTABLE)

clean:
	@echo "Cleaning up..."
	@rm -f $(API_BINARY)

test:
	@echo "Running test..."
	@rm -f $(COVERAGE_FILE) $(FILTERED_COVERAGE_FILE)
	@echo "mode: atomic" > $(COVERAGE_FILE)

	@for pkg in $$(go list ./internal/... | grep -vE $(IGNORED_PACKAGE)); do \
		go test -coverprofile=profile.out -covermode=atomic $$pkg || exit 1; \
		if [ -f profile.out ]; then \
			cat profile.out >> $(COVERAGE_FILE); \
			rm profile.out; \
		fi; \
	done

	@echo "mode: atomic" > $(FILTERED_COVERAGE_FILE)
	@grep -v -E $(IGNORED_PACKAGE) coverage.out | grep -v "mode:" $(COVERAGE_FILE) >> $(FILTERED_COVERAGE_FILE)

	@total_coverage=$$(go tool cover -func=$(FILTERED_COVERAGE_FILE) | grep '^total:' | awk '{print int($$3)}' | sed 's/%//g'); \
	echo "Coverage threshold: ${COVERAGE_THRESHOLD}%"; \
	echo "Total coverage: $$total_coverage%"; \
	if [ $$total_coverage -lt $(COVERAGE_THRESHOLD) ]; then \
		echo "Warning! Coverage below threshold!"; \
		rm -f $(COVERAGE_FILE) $(FILTERED_COVERAGE_FILE); \
		exit 1; \
	fi; \
	
	@rm -f $(COVERAGE_FILE) $(FILTERED_COVERAGE_FILE)