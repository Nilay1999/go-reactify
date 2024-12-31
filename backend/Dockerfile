# Use an official Go image for building the application
FROM golang:1.23 AS builder

ARG DATABASE_URL
ARG DATABASE_PASS
ARG PORT
ARG STAGE

ENV  DATABASE_URL=${DATABASE_URL}
ENV  DATABASE_PASS=${DATABASE_PASS}
ENV  PORT=8080
ENV  STAGE=production

# Set environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create and set the working directory
WORKDIR /app

# Copy go.mod and go.sum to the container and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal image for the final production build
FROM alpine:latest

# Set working directory in the minimal image
WORKDIR /root/

# Copy the built application from the builder
COPY --from=builder /app/main .

# Expose the application port (update the port as per your application)
EXPOSE 8080

# Run the application
CMD ["./main"]
