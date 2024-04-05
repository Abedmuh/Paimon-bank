FROM golang:alpine3.19 AS builder

# Linux
LABEL maintainer="Abdillah ProjectSprint"
RUN apk update && apk add --no-cache git

# app
WORKDIR /app

COPY go.mod go.sum ./
RUN go get -u ./...
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

FROM alpine:latest  

WORKDIR /root/
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .
CMD ["./main"]