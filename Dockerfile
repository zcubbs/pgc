# Build stage for the Go application
FROM golang:1.21 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Final stage to set up PostgreSQL and the Go binary
FROM postgres:14

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=secret

WORKDIR /app

# Copy from build stage
COPY --from=builder /app/main /app/main
COPY --from=builder /app/config.yaml /app/config.yaml

# Give execution permissions
RUN chmod +x /app/main

# Custom entry point to ensure PostgreSQL is ready before starting the application
COPY ./docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

# Set custom entrypoint
ENTRYPOINT ["/docker-entrypoint.sh"]
