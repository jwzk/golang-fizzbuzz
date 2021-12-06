# Debian based image that provide test tools requirements (gcc)
FROM golang:1.17-bullseye

# Install go project dependencies
WORKDIR /go/src/fizzbuzz
ADD ./src/go.mod /go/src/fizzbuzz
RUN go mod download

# Copy project
ADD ./src /go/src/fizzbuzz

# Build app
RUN go build -o /fizzbuzz/main .

# Launch web server
ENTRYPOINT ["/./fizzbuzz/main"]
