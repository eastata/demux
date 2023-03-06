# Demux
[![Go Reference](https://pkg.go.dev/badge/github.com/eastata/demux.svg)](https://pkg.go.dev/github.com/eastata/demux) 
[![codecov](https://codecov.io/gh/eastata/demux/branch/main/graph/badge.svg?token=9S3GK2DBP5)](https://codecov.io/gh/eastata/demux)

Runs a task in parallel

## Usage

### CLI

```shell
git clone git@github.com:eastata/demux.git
cd demux/cmd/cli
go run .
```

### API server

* [API Specification Markdowndocs](./docs/api.md)
* [API Specification in OpenAPI 2.0 json](./swaggerui/swagger.json)
* Run **Demux API server** with swagger:
    ```shell
    # export SERVER_ADDRESS=127.0.0.1
    # export SERVER_PORT=8080
    git clone git@github.com:eastata/demux.git
    cd demux/cmd/api-server
    go run .
    ```
* Swagger UI: [http://127.0.0.1:8080/swaggerui/](http://127.0.0.1:8080/swaggerui/)

## Tests
```shell
go test -v ./...
go test -bench=. ./...
go test -cover ./...
go test -coverprofile coverage.out ./...
# Check coverage
go tool cover -html=c.out
```
## Dcoker

* Build
  ```shell
  docker build -t eastata/demux .
  ```
* Run
  ```shell
  # CLI
  docker run --rm -it ghcr.io/eastata/demux cli
  # API server
  docker run --rm -it -p 8080:8080 ghcr.io/eastata/demux
  ```

## Makefile draft
```shell
# Install swagger https://github.com/go-swagger/go-swagger
download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \\n  jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')\ncurl -o /usr/local/bin/swagger -L'#' "$download_url"\nchmod +x /usr/local/bin/swagger

# Cleanup swagger-ui

SWAGGER_VERSION="4.16.1"
rm -rf ./swaggerui/*
wget https://github.com/swagger-api/swagger-ui/archive/refs/tags/v$SWAGGER_VERSION.tar.gz
mkdir ./swagger_tmp
tar -C ./swagger_tmp -xvf ./v$SWAGGER_VERSION.tar.gz swagger-ui-$SWAGGER_VERSION/dist 
mv ./swagger_tmp/swagger-ui-$SWAGGER_VERSION/dist/* ./swaggerui
rm -rf ./swagger_tmp
rm v$SWAGGER_VERSION.tar.gz

# swaggerui/swagger-initializer.js
# !install GNU sed for MacOS!
sed -i "s/https\:\/\/petstore\.swagger\.io\/v2\/swagger\.json/.\/swagger.json/" swaggerui/swagger-initializer.js

# Generate swagger yaml spec
swagger generate spec -o ./swaggerui/swagger.json -m -w ./cmd/api-server/

# Generate ./docs/api.md form spec
swagger generate markdown --skip-validation -f ./swaggerui/swagger.json --output="./docs/api.md"
```
