FROM golang:1.19-alpine AS builder
# ENV PG_VERSION="9.3.4"
WORKDIR /app
COPY go.mod go.sum ./
COPY webapi ./webapi
COPY main.go .
RUN go mod download

# RUN CGO_ENABLED=0 go build -o bin/webapi
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/webapi

# FROM gcr.io/distroless/base
FROM scratch
WORKDIR /app
COPY  --from=builder /app/bin/webapi /app/bin/webapi
ENTRYPOINT ["./bin/webapi"]


# EXPOSE 80/tcp
# EXPOSE 80/udp