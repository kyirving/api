FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o server ./cmd/main.go

FROM alpine:3.21

RUN adduser -D appuser

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/config/config.yaml ./config/config.yaml

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --retries=3 \
    CMD wget -qO- http://localhost:8080/ping || exit 1

USER appuser

CMD ["./server"]
