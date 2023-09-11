bindir:
	@mkdir -p bin

build: bindir
	@go build -o bin/server .

run: build
	./bin/server

clean:
	rm -rf bin