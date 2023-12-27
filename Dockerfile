
FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o url-shortener .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/url-shortener .

EXPOSE 8080

CMD ["./url-shortener"]
