 
FROM golang:1.16-alpine
# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main cmd/api/main.go

# Expose port 8989 to the outside world
EXPOSE 8989

# Command to run the executable
CMD ["./main"]
