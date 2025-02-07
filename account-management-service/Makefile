TEST_PKG := ./...
COVER_FILE := ./coverage.out

.PHONY: run-tests
run-tests:
	@go test ${TEST_PKG} -coverpkg=${TEST_PKG} -coverprofile ${COVER_FILE} && go tool cover -func ${COVER_FILE}
