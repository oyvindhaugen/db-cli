FROM ubuntu:latest
# COPY db-cli.exe .
COPY db.go .
COPY main.go .
COPY website.go .
COPY go.mod .
COPY go.sum .
COPY static /static
RUN apt-get update
RUN apt-get install -y ca-certificates
RUN apt-get install -y golang
RUN apt-get install -y vim