FROM golang:latest
RUN apt-get update && apt-get install make && \
    go get github.com/go-playground/validator && \
    go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1 && \
    go get github.com/lib/pq
RUN mkdir /code
WORKDIR /code
COPY . /code/
COPY go.mod ./
RUN go mod download && go mod tidy