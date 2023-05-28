# Start from a Debian based image with Go installed
FROM golang:1.20

# Set the current working directory inside the container 
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. 
# Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the application 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/order-service

# Expose port 50051 to the outside world
EXPOSE 50051

# Command to run the executable 
CMD ["./main"] 