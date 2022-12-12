FROM ubuntu:latest

COPY db.go .
COPY main.go .
COPY website.go .
COPY hash.go .
COPY go.mod .
COPY go.sum .
COPY static /static
RUN apt-get update
RUN apt-get install -y ca-certificates
RUN apt-get install -y golang
RUN go get golang.org/x/crypto
RUN go get github.com/lib/pq
RUN go run .