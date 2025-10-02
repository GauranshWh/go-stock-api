# Stage 1: The build environment, where we compile our app
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files first to leverage Docker's layer caching.
# Dependencies will only be re-downloaded if go.mod or go.sum changes.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of our source code
COPY . .

# Build the application, creating a static, self-contained binary.
# CGO_ENABLED=0 is important for creating a static binary.
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .


# Stage 2: The final, minimal production image
FROM alpine:latest

WORKDIR /root/

# Copy only the compiled binary from the 'builder' stage.
# Nothing else from the build environment is included.
COPY --from=builder /main .

# This is the command that will run when the container starts
CMD ["./main"]