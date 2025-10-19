# -------- Build stage --------
FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .


FROM alpine:latest
WORKDIR /app/

COPY --from=builder /build/main .
COPY --from=builder /build/templates /build/templates

CMD ["./main"]
EXPOSE 8080

