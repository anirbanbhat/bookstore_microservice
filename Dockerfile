# Use the official Go image as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files into the container
COPY . .

# Install dependencies (Echo framework)
RUN go mod download

# Build the Go application
RUN go build -o bookstore_microservice

# Expose the port that the application listens on
EXPOSE 8080

# Run the Go application
CMD ["./bookstore_microservice"]
