# ---- Build stage ----
FROM golang:1.22 AS builder
WORKDIR /app

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Set GO env for static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./...

# ---- Runtime stage ----
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/server /app/server

# Change port if your app uses a different one
EXPOSE 8080

# If you need env vars, use: ENV VAR=value
ENTRYPOINT ["/app/server"]
