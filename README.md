# Uber's zap Go Logging Library Examples

[![Build Status](https://github.com/arcanericky/zap-examples/workflows/Build/badge.svg?branch=main)](https://github.com/arcanericky/zap-examples/actions?query=workflow%3ABuild)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

This repository provides some examples of using [Uber's zap](https://github.com/uber-go/zap) Go logging library

The individual examples can be executed with `go run`, however the `README.md` files in each example's directory also show explanations alongside the output.

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
* [Using the Sampler](./src/sampler)
* [Hierarchical Logging](./src/hierarchical)