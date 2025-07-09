help:
	@echo 'usage: make build|tests|clean|help'

prep:
	mkdir bin

build:
	go fmt cmd/theta/theta.go
	go fmt cmd/threpl/threpl.go
	go build -o bin/theta cmd/theta/theta.go
	go build -o bin/threpl cmd/threpl/threpl.go

tests:
	@go test ./pkg/core/sexp ./pkg/core/types ./pkg/core
	@go test ./pkg/site

clean:
	@rm -r bin
