FROM golang:1.17-alpine AS build

# Install go project dependencies
WORKDIR /go/src/fizzbuzz
ADD ./src/go.mod /go/src/fizzbuzz
RUN go mod download

# Copy project
ADD ./src /go/src/fizzbuzz

# Build app
RUN go build -o /fizzbuzz/main .

# Minimal image
FROM alpine:latest

WORKDIR /fizzbuzz

# Copy binary and .env files from builder
COPY --from=build /fizzbuzz/main /fizzbuzz/main
COPY --from=build /go/src/fizzbuzz/.env /fizzbuzz/.env

# Launch web server
ENTRYPOINT ["/./fizzbuzz/main"]
