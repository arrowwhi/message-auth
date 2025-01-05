# Stage 1: Builder
FROM golang:1.23.3 AS builder

# Arguments for building
ARG GO_SVC_PATHS
ARG SERVICE_NAME
ARG VERSION

# Working directory inside the container
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

# Copy all files to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -mod=mod -a -installsuffix cgo \
    -o app ./cmd/auth_service/main.go

# Stage 2: Production-ready image
FROM alpine:3.19

# Working directory inside the container
WORKDIR /app

# Copy the built application and documentation
COPY --from=builder /src/app .
#COPY --from=builder /src/internal/docs ./docs

# Add a non-root user and switch to it
RUN addgroup -S app && adduser -S app -G app
USER app

# Expose a port (optional, set the port your app listens to, e.g., 8080)
EXPOSE 8080
#HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 CMD curl -f http://localhost:8080/health || exit 1


# Start the application
CMD ["/app/app"]
