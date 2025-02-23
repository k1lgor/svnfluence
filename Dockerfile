# Stage 1: Build the Go binary
FROM golang:1.20 AS builder

WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the binary with static linking for smaller image
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o svnfluence cmd/main.go

# Stage 2: Create the runtime image
FROM alpine:3.18

# Install ca-certificates for HTTPS requests (needed for OpenAI API)
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/svnfluence .

# Create a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Set environment variable for OpenAI API key (to be provided at runtime)
ENV OPENAI_API_KEY=

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./svnfluence"]
