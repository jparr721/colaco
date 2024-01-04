# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o server ./cmd/server/main.go

# Expose port 8000
EXPOSE 8000

# Set the entry point of the container
CMD ["./server"]
