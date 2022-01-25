SHELL := /bin/bash

.PHONY: run build deps-upgrade release test
# https://stackoverflow.com/questions/2214575/passing-arguments-to-make-run
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif
run:
	go run cmd/fab/fab.go $(RUN_ARGS)
build:
	goreleaser build --single-target --snapshot --rm-dist

deps-upgrade:
	go get -u -t -d -v ./...
	go mod tidy

release:
	goreleaser --rm-dist

test:
	go test -v ./...
