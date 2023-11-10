Greeter in Go
===

Currrently the default testing framework is used. A Makefile is included for running the tests conveniently. 

## Built-in test framework

Things to know:
* Uses the [testing] package from the [standard lib][std-lib]
* There are no built-in assertions, it has to be written by you
* t.Errorf will fail the test with the given string formatting

To run the tests using any of these options:
* `make test` 
* `go test ./pkg/greeter/hello_test.go` 

## Future plans

1. Adding testify framework
2. Adding ginkgo/gomega framework

[std-lib]: https://pkg.go.dev/std
[testing]: https://pkg.go.dev/testing