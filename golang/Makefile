
bin/ray-tracer: main.go
	go build -o bin/ray-tracer main.go

build: bin/ray-tracer

clean:
	rm -rf bin

run: bin/ray-tracer
	bin/ray-tracer

test:
	go test -v ./...

.PHONY: clean test run build