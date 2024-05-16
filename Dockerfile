# Use the default Golang image to compile the source code
FROM golang:1.22 as builder

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files to download the dependencies
COPY go.mod go.sum ./ 

# Download the Go module dependencies
RUN go mod download

# Copy the entire source code to the container
COPY . .

# Compile the source code into an executable
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# Create the final image that will run the API and expose port 8080
FROM debian:stable-slim

# Set the working directory
WORKDIR /root/

# Copy the compiled executable from the builder image
COPY --from=builder /app/main .

# Install necessary dependencies to run the Go binary
RUN apt-get update && apt-get install -y ca-certificates && apt-get clean

# Set the PORT environment variable
ARG PORT
ENV PORT=$PORT

# Expose the configured port
EXPOSE $PORT

# Command to run the executable
CMD ["./main"]
