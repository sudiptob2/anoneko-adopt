# Go Version
FROM golang:1.21-rc-alpine3.17

# Environment variables which CompileDaemon requires to run
ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    PORT=8080

# Basic setup of the container
RUN mkdir /app
COPY . /app
WORKDIR /app

RUN cp .env.example .env

RUN go build -o bin/api

EXPOSE $PORT

ENTRYPOINT ["./bin/api"]
