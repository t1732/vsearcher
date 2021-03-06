GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOFMT=$(GOCMD) fmt
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVERSION=$(shell go version)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

NAME=vsearcher
BIN=./bin
CMD=$(BIN)/$(NAME)
MIGRATE=migrate
SEED=seed

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BIN)

.PHONY: fmt
fmt:
	$(GOFMT) ./...

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: build
build:
	mkdir -p $(BIN)
	$(GOBUILD) -o $(CMD) ./cmd/$(NAME)
	$(GOBUILD) -o $(BIN)/$(MIGRATE) ./cmd/$(MIGRATE)
	$(GOBUILD) -o $(BIN)/$(SEED) ./cmd/$(SEED)

.PHONY: run
run: test build
	@$(CMD)
