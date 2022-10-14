# syntax=docker/dockerfile:1
## Build
FROM golang:1.19.2-alpine3.16 AS build
WORKDIR /app
COPY go.mod ./
# COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /main

## Deploy
FROM alpine:3.16.2
WORKDIR /
COPY --from=build /main /main
ENTRYPOINT ["/main"]