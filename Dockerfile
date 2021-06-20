FROM golang:latest
RUN go get github.com/go-playground/validator && \
    go get github.com/lib/pq
RUN mkdir /code
WORKDIR /code
COPY . /code/
COPY go.mod ./
RUN go mod download && go mod tidy