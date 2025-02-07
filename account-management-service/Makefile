COVERAGE_FILE := ./coverage.out
COVER_FILE := ./cover.out

.PHONY: test
test:
	@go test ./... -coverprofile ${COVERAGE_FILE} && cat ${COVERAGE_FILE} | grep -v 'gen' > ${COVER_FILE} && go tool cover -func ${COVER_FILE}

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: test-env-up
test-env-up:
	@docker compose up --exit-code-from migrate migrate

.PHONY: test-env-down
test-env-down:
	@docker compose down -v
