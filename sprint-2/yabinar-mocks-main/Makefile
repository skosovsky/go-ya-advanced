.PHONY: all
all: build

.PHONY: build
build: generate-mocks lint
	go build -o cmd/foobar/foobar ./cmd/foobar

.PHONY: lint
lint: generate-mocks
	./run-linting.sh

.PHONY: generate-mocks
generate-mocks:
	go generate ./...

.PHONY: clean
clean: clean-mocks clean-bin

.PHONY: clean-mocks
clean-mocks:
	rm -rf internal/core/services/internal/mocks
	rm -rf internal/core/services/internal/mocks-reflect

.PHONY: clean-bin
clean-bin:
	rm -f cmd/foobar/foobar
