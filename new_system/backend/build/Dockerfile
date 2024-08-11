# Stage 1: Build the Go app
FROM golang:1.22.6-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY . .

WORKDIR /app/backend

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Build the Go app
RUN go build -o main cmd/main.go

# Stage 2: Run the Go app
FROM alpine:3.20.1

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/backend/main .

# Expose port 8080 to the outside world
EXPOSE 9981

# Command to run the executable
CMD ["./main"]