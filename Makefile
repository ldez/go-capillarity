.PHONY: all

PKGS := $(shell go list ./... | grep -v '/vendor/')

default: clean checks test-unit build

clean:
	rm -f cover.out

build:
	go build

test-unit:
	go test -v -cover $(PKGS)

checks:
	golangci-lint run
