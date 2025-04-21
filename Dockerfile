FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/server ./cmd/server/main.go

FROM alpine:latest

RUN apk add --no-cache tzdata ca-certificates

RUN mkdir -p /dev && \
    ln -sf /proc/self/fd/1 /dev/stdout && \
    ln -sf /proc/self/fd/2 /dev/stderr

COPY --from=builder /app/bin/server /server

CMD ["/server"]