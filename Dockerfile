FROM golang:alpine
COPY . /go/src/
WORKDIR /go/src/
RUN go install cmd/theta/theta.go
WORKDIR /go
EXPOSE 1958/tcp
CMD ["/go/bin/theta"]
LABEL org.opencontainers.image.source=https://github.com/TheDevtop/theta-go
