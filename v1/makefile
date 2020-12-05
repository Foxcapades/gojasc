V1_PATH := $(shell pwd)
ifneq ("$(shell basename $(V1_PATH))", "v1")
V1_PATH = $(shell pwd)/v1
endif

.PHONY: test
test:
	@if gcc -dumpmachine | grep -q x86_64; then \
		go test -v ./...; \
	else \
		GOARCH=386 go test -v ./...; \
	fi
