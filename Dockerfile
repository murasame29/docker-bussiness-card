FROM golang:1.21.3 AS builder
WORKDIR /build

COPY . .

RUN go build -o main .

FROM alpine:latest AS runner
WORKDIR /app

COPY --from=builder /build/main .

CMD ["./main"]