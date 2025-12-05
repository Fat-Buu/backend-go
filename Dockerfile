# Stage 1: Build
FROM golang:1.25.3-alpine AS builder

# Set workdir
WORKDIR /app

# Copy go.mod and go.sum first (cache dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all source code
COPY . .

# Build binary
RUN go build -o backend-go ./cmd/main.go

# Stage 2: Run
FROM alpine:latest

# Install CA certificates (required for HTTPS)
RUN apk --no-cache add ca-certificates

# Set workdir
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/backend-go ./backend-go

# ต้อง copy .env เข้า container
COPY .env .env               

# Expose port (Fiber default)
EXPOSE 3000

# Run binary
CMD ["./backend-go"]