FROM golang:1.17

ARG WORKMODE=memory

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get clean
RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go mod download
RUN go build -o shortener ./cmd/main.go

CMD ["./shortener"]
