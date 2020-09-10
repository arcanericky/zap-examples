# Uber's zap Go Logging Library Examples

This repository provides some examples of using [Uber's zap](https://github.com/uber-go/zap) Go logging library

The individual examples can be executed with `go run`, however the `README.md` files in each example's directory also show the output alongside explanations for the output.

```console
$ go run src/simple1/main.go
...
$ go run src/customlogger/main.go
```

## Examples

* [Simplest usage using presets](./src/simple1)
* [Creating a custom logger](./src/customlogger)
* [Using the global logger](./src/globallogger)
* [Creating custom encoders for metadata fields](./src/customencoder)