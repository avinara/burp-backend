# Use the official golang image as a base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /burp-backend

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build 

# Expose the port on which the Go application will listen
EXPOSE 8080

# Command to run the Go application
CMD ["./burp-backend"]