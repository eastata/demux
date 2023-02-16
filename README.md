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

## Tests
```shell
go test -v ./...
go test -bench=. ./...
go test -cover ./...
go test -coverprofile coverage.out ./...
# Check coverage
go tool cover -html=c.out
```
