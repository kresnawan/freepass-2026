FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./cmd/api


FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 3000
CMD ["./main"]