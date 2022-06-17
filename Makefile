unit-tests:
	ginkgo -r --race --randomize-all --randomize-suites --tags=unit

unit-tests-verbose:
	ginkgo -r -v -p --race --randomize-all --randomize-suites --tags=unit

unit-tests-report:
	mkdir -p "coverage" \
	&& go test ./tests/... -v -coverprofile=coverage/cover.out -coverpkg ./... --tags=unit \
	&& go tool cover -html=coverage/cover.out -o coverage/cover.html \
	&& go tool cover -func=coverage/cover.out -o coverage/cover.functions.html

unit-tests-cover:
	go test ./tests/... -coverpkg ./... --tags=unit


.PHONY: unit-tests,
		unit-tests-verbose,
		unit-tests-report,
		unit-tests-cover
