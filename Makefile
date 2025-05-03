build:
	go fmt cmd/theta.go
	go build -o theta cmd/theta.go

clean:
	rm theta
