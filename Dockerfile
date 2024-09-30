# Use the official Go image as the builder
FROM golang:1.21.1 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY backend/go.mod backend/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the backend code
COPY backend/ .

# Build the Go app for Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Use a smaller image for the final stage
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main ./

# Copy the .env file into the container
COPY backend/.env .env

# Expose the port
EXPOSE 8080

# Run the Go binary
CMD ["./main"]
