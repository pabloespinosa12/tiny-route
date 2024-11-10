# Use the official Golang image as the base image
FROM golang:1.23 AS base

# Set the working directory inside the container
WORKDIR /app

# Install air for hot reloading
#RUN go install github.com/air-verse/air@v1.61.1

# Install CompileDaemon for hot reloading
RUN go install github.com/githubnemo/CompileDaemon@latest

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

RUN go env -w GOFLAGS="-buildvcs=false" 

# Copy the source code into the container
COPY . .

# Expose the application port
EXPOSE 8080

# Run air for hot reloading
#WORKDIR /app/cmd/tiny-route
#CMD ["air"]
CMD ["CompileDaemon", "-polling=true", "--build=go build -o main ./cmd/tiny-route/", "--command=./main"]