ifndef GOBIN
	$(error GOBIN is not set)
endif

compile:
	go build ./cmd/guineapig -o guineapig

dep:
	go get -d

build:
	dep
	compile

install:
	go install ./cmd/guineapig

.PHONY:
	build