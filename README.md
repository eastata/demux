# Demux
[![Go Reference](https://pkg.go.dev/badge/github.com/eastata/demux.svg)](https://pkg.go.dev/github.com/eastata/demux)

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
go test -coverprofile c.out ./...
# Check coverage
go tool cover -html=c.out
```
