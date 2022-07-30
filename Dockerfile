# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.16-alpine3.13 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build -o main .

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

EXPOSE 8000
CMD [ "/app/main" ]