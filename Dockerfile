#### Bulding executable
FROM golang:1.18.6 as build-stage

WORKDIR $GOPATH/src/packages/web-server
COPY . .

# Install dependencies
RUN go mod download

# Build the binary
RUN GOARCH=amd64 GOOS=linux go build -o /build-output/web-server main.go

#### Building small image
FROM ubuntu:kinetic as image
ENV GOENV="production"

WORKDIR /app

# Copy executable
COPY --from=build-stage /build-output/web-server ./
COPY ./config/server.toml ./config/

EXPOSE 8080

ENTRYPOINT ["/app/web-server"]

