# syntax=docker/dockerfile:1

# Use an official Golang runtime as a parent image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Build the Go application
RUN go build ./cmd/main.go

# Set the entry point for the container
ENTRYPOINT ["./main"]
