# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=gofmt
GOLINT=golint
VET=@echo "go vet $@"; go vet
TEST=./...
BINARY_NAME=heartbeat-manager

all: fmt vet test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v
	rm $(BINARY_NAME)

test:
	$(GOTEST) $(TEST) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

fmt:
	$(GOFMT) -w .

vet:
	$(VET)

lint:
	$(GOLINT) ./...

all: vet fmt lint test
