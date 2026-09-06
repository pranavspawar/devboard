# ---------- Build stage ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Download dependencies first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o devboard-backend .


# ---------- Runtime stage ----------
FROM alpine:3.22

WORKDIR /app

# Install only required runtime certificates
RUN apk upgrade --no-cache && \
    apk add --no-cache ca-certificates && \
    adduser -D -u 10001 app

COPY --from=builder /app/devboard-backend .

RUN chown app:app /app/devboard-backend

USER app

EXPOSE 8080

ENTRYPOINT ["./devboard-backend"]