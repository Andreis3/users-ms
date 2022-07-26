unit-tests:
	ginkgo -r --race --tags=unit

unit-tests-verbose:
	ginkgo -r -v --race --tags=unit

unit-tests-report:
	mkdir -p "coverage" \
	&& go test ./tests/... -v -coverprofile=coverage/cover.out -coverpkg ./... --tags=unit \
	&& go tool cover -html=coverage/cover.out -o coverage/cover.html \
	&& go tool cover -func=coverage/cover.out -o coverage/cover.functions.html

unit-tests-cover:
	go test ./tests/... -coverpkg ./... --tags=unit

tag:
	./scripts/bump_version.sh

.PHONY: unit-tests,
		unit-tests-verbose,
		unit-tests-report,
		unit-tests-cover,
		tag
