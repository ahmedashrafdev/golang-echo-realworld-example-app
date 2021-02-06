# Start from base image
FROM golang:alpine

ENV GO111MODULE=on

# Set the current working directory inside the container
WORKDIR /backend

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy source from current directory to working directory
COPY . .

# Build the application
RUN go build .

RUN chmod +x ./golang-echo-realworld-example-app

# Expose necessary port
EXPOSE 8585

# Run the created binary executable after wait for mysql container to be up
CMD ["./golang-echo-realworld-example-app"]