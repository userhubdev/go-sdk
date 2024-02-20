.PHONY: setup
setup:
	@sh -c "echo '#!/bin/sh\n\nmake lint test' > .git/hooks/pre-commit"
	@chmod 755 .git/hooks/pre-commit

.PHONY: fmt
fmt:
	@goimports -w .

.PHONY: lint
lint:
	@test -z "$$( gofmt -l . )" || (echo "Needs formatting:" && gofmt -l . && exit 1)

.PHONY: test
test:
	@go test -count=1 ./...
