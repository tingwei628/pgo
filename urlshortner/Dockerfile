# Build stage
FROM golang:1.23-alpine AS build

WORKDIR /build

COPY ./urlshortner/go.mod .

RUN go mod download

COPY ./urlshortner .

# -ldflags="-s -w" to strip debugging information.
RUN go build -v -o ./app -ldflags="-s -w" ./cmd/app

# Runtime stage
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /build/app /app

ENTRYPOINT ["/app"]