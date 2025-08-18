# Use the official Golang image as a base
FROM golang:1.24.1-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main/main.go

# Expose the port your application listens on (if applicable)
EXPOSE 8050

# Define the command to run your application
CMD ["./main"]

