FROM golang:latest AS builder
COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 go build -ldflags="-extldflags=-static -s -w" -o cssloop

FROM alpine:latest

COPY --from=builder /src/cssloop /app/cssloop

RUN chmod +x /app/cssloop

WORKDIR /app

CMD ["/app/cssloop"]
