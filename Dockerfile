# Build stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o inventory ./cmd/inventory

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/inventory .
EXPOSE 8080
CMD ["./inventory"]