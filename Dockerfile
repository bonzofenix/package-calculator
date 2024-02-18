# Use the official Golang image as the base image
FROM golang:1.21.0 AS builder

# Set the current working directory inside the container
WORKDIR /src

# Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go app
ENV CGO_ENABLED=0
RUN go build  -o cmd/app cmd/main.go

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the executable from the builder stage
COPY --from=builder /src/cmd/app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./app"]




