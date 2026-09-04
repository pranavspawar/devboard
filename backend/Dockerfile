FROM dhi.io/golang:1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o devboard-backend .

FROM alpine:3.21

WORKDIR /app

# Install minimal runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /app/devboard-backend ./devboard-backend

# Ensure binary is executable
RUN chmod +x ./devboard-backend

# Run as non-root user
USER nobody

EXPOSE 8080

# Healthcheck to verify container is alive
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Use exec form for proper signal handling
ENTRYPOINT ["./devboard-backend"]