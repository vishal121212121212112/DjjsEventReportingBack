# Start from the official Golang image
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .


# Build the Go application
RUN go build -o main ./cmd/main/main.go

# Expose port
EXPOSE 8080


# Command to run
CMD ["./main"]