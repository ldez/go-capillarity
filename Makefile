.PHONY: all

PKGS := $(shell go list ./... | grep -v '/vendor/')

default: clean lint checks test-unit build

clean:
	rm -f cover.out

build:
	go build

test-unit:
	go test -v -cover $(PKGS)

lint:
	golint -set_exit_status $(PKGS)

checks:
	staticcheck $(PKGS)
	gosimple $(PKGS)