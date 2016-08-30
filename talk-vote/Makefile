GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test

all: build

build:
	$(GOBUILD) -v -o bin/talk-vote

install:
	$(GOINSTALL)

clean:
	$(GOCLEAN) -n -i -x
	rm -f $(GOPATH)/bin/talk-vote
	rm -rf bin/talk-vote

test:
	$(GOTEST) -v -cover ./...

.PHONY: all clean