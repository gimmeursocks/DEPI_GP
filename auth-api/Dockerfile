# Stage 1: Build Stage
FROM golang:1.23-alpine AS builder

# Create and set the working directory
WORKDIR /app

# Copy go.mod and go.sum to leverage Docker cache
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod tidy CGO_ENABLED=0 

# Copy rest of application code
COPY . .

# Build the application (No DEBUG and No C Code in GO)
RUN go build -ldflags="-s -w" -o myapp ./main.go

# Stage 2: Minimal Runtime Image
FROM scratch

# Set working directory
WORKDIR /app

# Copy built app from the builder stage
COPY --from=builder /app/myapp .

# Best practice to use a non-root user
RUN addgroup -S app && adduser -S app -G app
USER app

# Expose the app port
EXPOSE 8000

# Run the Go application
CMD ["/app/myapp"]