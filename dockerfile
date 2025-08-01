# Start with the official Go image for building the application
# This is our first stage, named 'builder'
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
# This is a key step to enable caching and only re-download dependencies when they change
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod tidy

# Copy the rest of the application source code
COPY . .

# Build the application
# The -o flag specifies the output name of the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -o myapp .

# Now, use a small, secure base image for the final application
# This is our second stage, which will only contain the compiled application
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the 'builder' stage
COPY --from=builder /app/myapp .

# Expose the port that the application listens on
EXPOSE 10000

# Run the application
CMD ["./myapp"]
