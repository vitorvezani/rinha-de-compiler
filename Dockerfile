# syntax=docker/dockerfile:1

# Step 1
FROM golang:1.21 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build ./cmd/main.go

# Step 2
FROM node:lts-slim
WORKDIR /app
COPY --from=build /app/main .
ENTRYPOINT ["./main"]