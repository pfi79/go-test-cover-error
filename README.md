# go 1.22 errors with go test -cover

## PASS

```shell
go test ./...
```

```shell
go test -v ./... 
```

```shell
go test -v -cover github.com/pfi79/go-test-cover-error/b
```

## FAIL

```shell
go test -v -cover /... 
```

```shell
go test -v -cover github.com/pfi79/go-test-cover-error/a github.com/pfi79/go-test-cover-error/b 
```

```shell
go test -v -cover github.com/pfi79/go-test-cover-error/b github.com/pfi79/go-test-cover-error/a
```

## Problem 

An extra covervars.go file appears in the github.com/pfi79/go-test-cover-error/a package. Because of this this package is different in plugin and in test.