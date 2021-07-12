

all: bin/chapter2 bin/chapter4 bin/chapter5 bin/chapter6 bin/chapter7

bin/%: cmd/%/main.go
	go build -o $@ $^

clean:
	rm -rf bin

test:
	go test ./...

.PHONY: clean test run build