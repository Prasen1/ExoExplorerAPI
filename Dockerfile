FROM golang:1.18-alpine

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Run unit tests
RUN go test -v ./...

# Build the Go application
RUN go build -o /exoplanet-microservice

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["/exoplanet-microservice"]
