# Build stage
FROM golang:1.23 AS build

WORKDIR /app
COPY go.* .
RUN go mod download

COPY . .
RUN go build -o urlshortner .

# Runtime stage
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /app/urlshortner /app/urlshortner

ENTRYPOINT ["/app/urlshortner"]
