#### Bulding executable
FROM golang:1.18.6-alpine3.16 as build-stage

# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache 'git=~2'

WORKDIR $GOPATH/src/packages/web-server
COPY . .

# Install dependencies
RUN go mod download

# Build the binary
RUN GOARCH=amd64 GOOS=linux go build -o /build-output/web-server main.go

#### Building small image
FROM alpine:3.16.2

WORKDIR /

# Copy executable
COPY --from=build-stage /build-output/web-server /

EXPOSE 8080

ENTRYPOINT ["/web-server"]

