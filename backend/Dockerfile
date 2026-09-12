# ---------- Build stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o devboard-backend .

# ---------- Runtime stage ----------
FROM alpine:3.22

WORKDIR /app

RUN apk upgrade --no-cache && \
    adduser -D -u 10001 app

COPY --from=builder /app/devboard-backend .

RUN chown app:app /app/devboard-backend

USER app

EXPOSE 8080

ENTRYPOINT ["./devboard-backend"]