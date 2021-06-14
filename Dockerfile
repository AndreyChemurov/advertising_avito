FROM golang:latest
RUN apt-get update && apt-get install make
RUN go get github.com/go-playground/validator
RUN mkdir /code
WORKDIR /code
COPY . /code/
COPY go.mod ./
RUN go mod download
RUN go mod tidy
EXPOSE 8000