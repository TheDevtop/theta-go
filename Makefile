# THETA list processor
# Build script

build: prep
	go build -o bin/theta cmd/main.go

prep:
	mkdir -p bin

clean:
	rm -f bin/theta
