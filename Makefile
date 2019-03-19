.PHONY: help
help:
	@echo "Usage: make <command>"

.PHONY: build
build:
	@echo "start build..."
	go build -o gen_leetcode
	@echo "build success!"