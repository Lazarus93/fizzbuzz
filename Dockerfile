# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="ojelaidi@gmail.com"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .


# Install the package swag for documentation
RUN go get -u github.com/swaggo/swag/cmd/swag

# Build the Go app
RUN go build -o main ./cmd/fizzbuzz

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./main"]
