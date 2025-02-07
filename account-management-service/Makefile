TEST_PKG := ./...
COVERAGE_FILE := ./coverage.out
COVER_FILE := ./cover.out

.PHONY: run-tests
run-tests:
	@go test ./... -coverprofile ${COVERAGE_FILE} && cat ${COVERAGE_FILE} | grep -v 'gen' > ${COVER_FILE} && go tool cover -func ${COVER_FILE}
