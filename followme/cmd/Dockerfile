# Stage 1: Build the Go app
FROM golang:1.22.4-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# List files to check if go.mod and go.sum are copied correctly
RUN ls -la

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main cmd/main.go

# Stage 2: Run the Go app
FROM alpine:3.20.1

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# List files to check if the binary is copied correctly
RUN ls -la

# Expose port 8080 to the outside world
EXPOSE 9981

# Command to run the executable
CMD ["./main"]
