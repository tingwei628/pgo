# Build stage
FROM golang:1.23-alpine AS build

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

# -trimpath to support reproduce build.
# -ldflags="-s -w" to strip debugging information.
RUN go build -o ./app -trimpath -ldflags="-s -w" .

# Runtime stage
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /build/app /app

ENTRYPOINT ["/app"]