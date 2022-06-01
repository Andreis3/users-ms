unit-tests:
	ginkgo -r --race --randomize-all --randomize-suites --tags=unit
unit-tests-verbose:
	ginkgo -r -v -p --race --randomize-all --randomize-suites --tags=unit
unit-tests-coverage-report:
	go test ./tests/... -v -coverprofile=coverage/cover.out -coverpkg ./... --tags=unit \
	&& go tool cover -html=coverage/cover.out -o coverage/cover.html \
	&& go tool cover -func=coverage/cover.out -o coverage/cover.functions.html

.PHONY: unit-tests,
		unit-tests-verbose,
		unit-tests-coverage-report
