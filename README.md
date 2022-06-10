# up


## Fast update from remote GitHub
```sh
go clean -modcache
rm go.sum
# delete require line inside `go.mod` file
export GOPROXY=direct
go mod tidy -v
go run .
```