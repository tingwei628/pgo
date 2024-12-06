# Build stage
FROM golang:1.23-alpine AS build

WORKDIR /build

COPY ./urlshortner/go.mod .

RUN go mod download

COPY ./urlshortner .

# -ldflags="-s -w" to strip debugging information.
RUN go build -mod=mod -v -o ./app -ldflags="-s -w" ./cmd/app

# Runtime stage
FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /
COPY --from=build /build/app .
COPY --from=build --chown=nonroot /build/internal/db.sqlite /build/internal/db.sqlite
COPY --from=build /build/internal/views /build/internal/views


ENTRYPOINT ["/app"]