BINARY_NAME=drs
VERSION ?=$(shell git describe --tags --always)

.PHONY: init-bin-dir
init-bin-dir:
	mkdir -p bin

.PHONY: build
build: clean init-bin-dir
	CGO_ENABLED=0 go build \
		-ldflags="-X 'main.Version=${VERSION}'" \
		-o bin/${BINARY_NAME} \
		github.com/kshitij10496/drs

.PHONY: clean
clean:
	rm -rf bin