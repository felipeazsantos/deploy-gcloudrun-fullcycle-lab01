FROM golang:1.24.3-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/cmd/server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /bin/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /bin/server .
COPY .env .
EXPOSE 8080
CMD ["./server"]
