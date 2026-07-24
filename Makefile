.PHONY: fmt lint check

fmt:
	gofmt -w ./...

lint:
	golangci-lint run ./...

check:
	@unformatted=$$(gofmt -l ./...); \
	if [ -n "$$unformatted" ]; then \
		echo "The following files are not formatted:"; \
		echo "$$unformatted"; \
		exit 1; \
	fi
