# Use the official Go image as the base image
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o go-case-converter main.go

# Use a minimal base image for the final container
FROM debian:bookworm-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/go-case-converter .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./go-case-converter"]