# syntax=docker/dockerfile:1

FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

COPY ./cmd/server ./cmd/server

# Build
# RUN go build -o ./cmd/server/server ./cmd/server

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 3000

# Run
# CMD ["/bin/sh", "-c", "ls -la & ls -la /app && ls -la /app/cmd/server && /app/cmd/server/main"]

# Run go run main.go
CMD ["go", "run", "./cmd/server/main.go"]