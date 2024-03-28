FROM golang:alpine3.19 AS builder

# Linux
LABEL maintainer="Abdillah ProjectSprint"
RUN apk update && apk add --no-cache git

# app
WORKDIR /app

COPY go.mod go.sum ./
RUN go get -u ./...
COPY . .
RUN go build -o main ./cmd

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/main .
CMD ["./main"]
