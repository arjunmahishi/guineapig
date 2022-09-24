ifndef GOBIN
	$(error GOBIN is not set)
endif

compile:
	go build

install:
	go install

.PHONY:
	build
