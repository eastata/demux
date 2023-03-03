# Support setting various labels on the final image
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Demux in a stock Go builder container
FROM golang:1.19-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git bash curl jq

# Get dependencies - will also be cached if we won't change go.mod/go.sum
COPY go.mod /demux/
COPY go.sum /demux/
RUN cd /demux && go mod download

# Install /usr/local/bin/swagger
RUN download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url'); curl -o /usr/local/bin/swagger -L'#' "$download_url"; chmod +x /usr/local/bin/swagger

ADD . /demux

# Generate OpenAPI spec to /demux/swaggerui/swagger.json
RUN /usr/local/bin/swagger generate spec -o /demux/swaggerui/swagger.json -m -w /demux/cmd/api-server/

RUN cd /demux && go build -o cli /demux/cmd/cli/main.go
RUN cd /demux && go build -o api-server /demux/cmd/api-server/main.go

# Pull Demux into a second stage deploy alpine container
FROM alpine:latest

LABEL org.opencontainers.image.source=https://github.com/eastata/demux
LABEL org.opencontainers.image.description="Demux API server with cli"
LABEL org.opencontainers.image.licenses=GPLv3

RUN apk add --no-cache ca-certificates
COPY --from=builder /demux/swaggerui /swaggerui/
COPY --from=builder /demux/cli /usr/local/bin/
COPY --from=builder /demux/api-server /usr/local/bin/

EXPOSE 8080/tcp
CMD ["api-server"]

# Add some metadata labels to help programatic image consumption
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"

