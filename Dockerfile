FROM golang:latest AS build
RUN mkdir /code
WORKDIR /code
ENV CGO_ENABLED=0
COPY . /code/
COPY go.mod ./
RUN go mod download
RUN go mod tidy
RUN go build -o ./bin/main cmd/advertising_avito/main.go
COPY . .
EXPOSE 8000
CMD [ "./bin/main" ]