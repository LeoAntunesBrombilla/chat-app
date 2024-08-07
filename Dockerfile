# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /backend

# Copy go mod and sum files
COPY ./backend/go.mod ./
# ./backend/go.sum

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ./backend/. .

# Build the Go app
RUN go build -o main ./cmd

# Expose port 3000 to the outside
EXPOSE 3000

# Command to run the executable
CMD ["./main"]