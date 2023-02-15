# demux
Runs a task in parallel

## Usage
```shell
git clone git@github.com:eastata/go-demux.git
cd go-demux/cmd/demux
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