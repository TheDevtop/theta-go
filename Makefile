help:
	@echo 'usage: make build|docker|clean|help'

prep:
	mkdir bin

build:
	go fmt cmd/theta/theta.go
	go fmt cmd/threpl/threpl.go
	go build -o bin/theta cmd/theta/theta.go
	go build -o bin/threpl cmd/threpl/threpl.go

docker:build
	@docker build -f Dockerfile -t ghcr.io/thedevtop/theta-go .

clean:
	@rm -r bin
