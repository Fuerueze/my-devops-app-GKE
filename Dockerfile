# Stage 1: Build stage
#FROM golang:1.22 AS builder
FROM golang:1.21-bullseye AS builder

WORKDIR /app

# Copy the Go source code to the container
COPY . .

# Build the Go application
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Final stage
FROM debian:bullseye-slim

COPY --from=builder /app/main /main

CMD ["/main"]