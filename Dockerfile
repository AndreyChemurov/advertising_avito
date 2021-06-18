FROM golang:latest
RUN apt-get update && apt-get install make
RUN go get github.com/go-playground/validator
RUN go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1
RUN go get github.com/lib/pq
RUN mkdir /code
WORKDIR /code
COPY . /code/
COPY go.mod ./
RUN go mod download
RUN go mod tidy
EXPOSE 8000