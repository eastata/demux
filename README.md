# Demux
[![Go Reference](https://pkg.go.dev/badge/github.com/eastata/demux.svg)](https://pkg.go.dev/github.com/eastata/demux) 
[![codecov](https://codecov.io/gh/eastata/demux/branch/main/graph/badge.svg?token=9S3GK2DBP5)](https://codecov.io/gh/eastata/demux)

Runs a task in parallel

## Usage
```shell
git clone git@github.com:eastata/demux.git
cd demux/cmd/demux
go run main.go
```
## Swagger UI

[http://127.0.0.1:8080/swaggerui/](http://127.0.0.1:8080/swaggerui/)

## Tests
```shell
go test -v ./...
go test -bench=. ./...
go test -cover ./...
go test -coverprofile coverage.out ./...
# Check coverage
go tool cover -html=c.out
```

## Makefile draft
```shell
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
swagger generate spec -o ./swaggerui/swagger.json -m -w ./cmd/demux/


```