# Stage 1: Build the application
FROM golang:1.22 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application binary, targeting cmd/main.go
RUN go build -o main ./cmd

# Stage 2: Create a minimal image with the compiled binary
FROM alpine:3.16

# Set the working directory inside the container
WORKDIR /app

# Install required libraries (if necessary)
RUN apk add --no-cache libc6-compat

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations

# Ensure the binary has execution permissions
RUN chmod +x /app/main

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
