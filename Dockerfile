# Use a small and efficient base image
FROM golang:1.23-alpine AS builder

# Set necessary environment variables for Go modules
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy only the go.mod and go.sum files first
# This leverages Docker's layer caching for faster builds
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Use a smaller, non-root user for security
# (Create a user and group with limited permissions)
RUN addgroup -S appuser && adduser -S -G appuser appuser
USER appuser

# Use a distroless image for the final stage
# This reduces image size and attack surface
FROM gcr.io/distroless/base-debian11

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port (change if needed)
EXPOSE 8080

# Command to run the application
CMD ["./main"]
