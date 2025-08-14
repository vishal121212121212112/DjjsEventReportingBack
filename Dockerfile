# Start from the official Golang image
FROM golang:1.21-alpine AS builder
 
WORKDIR /app
 
# Install git (for go mod) and build tools
RUN apk add --no-cache git
 
# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download
 
# Copy the source code
COPY . .
 
# Build the Go app
RUN go build -o djjs-event-reporting-back ./cmd/main.go
 
# Start a minimal runtime image
FROM alpine:latest
 
WORKDIR /app
 
# Copy the built binary from builder
COPY --from=builder /app/djjs-event-reporting-back .
 
# Expose the port (change if your app uses a different port)
EXPOSE 8050
 
# Run the binary
CMD ["./djjs-event-reporting-back"]
