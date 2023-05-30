# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="bussiness.me@lucifergotmad.id"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/user

# Expose port 50051 to the outside world
EXPOSE 50051

# Command to run the executable
CMD ["./main"]
