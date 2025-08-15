# Start from the official Golang image
FROM golang:1.20-alpine AS builder

# Install git and other dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .
# Copy configuration files
COPY --from=builder /app/.env.dev ./.env.dev
COPY --from=builder /app/docs ./docs

# Expose port
EXPOSE 8080

# Set environment variables
ENV GIN_MODE=release

# Command to run
CMD ["./main"]