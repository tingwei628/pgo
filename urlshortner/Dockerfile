# Build stage
FROM golang:1.23 AS build

COPY . /app
WORKDIR /app

RUN go mod download

RUN go build -o main .

# Runtime stage
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /app/main /app/main

WORKDIR /app
ENTRYPOINT ["./main"]