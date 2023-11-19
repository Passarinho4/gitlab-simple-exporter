# Stage 1: Build Stage
FROM golang:1.19 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o gitlab-simple-exporter ./cmd/gitlab-simple-exporter

# Stage 2: Final Stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/gitlab-simple-exporter .

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./gitlab-simple-exporter"]