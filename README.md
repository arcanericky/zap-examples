# Uber's zap Go Logging Library Examples

[![Build Status](https://github.com/arcanericky/zap-examples/workflows/Build/badge.svg?branch=main)](https://github.com/arcanericky/zap-examples/actions?query=workflow%3ABuild)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

This repository provides some examples of using [Uber's zap](https://github.com/uber-go/zap) Go logging library

The individual examples can be executed with `go run`, however the `README.md` files in each example's directory also show explanations alongside the output.

```console
$ go run src/presets/main.go
...
$ go run src/customlogger/main.go
```

## Examples

* [Presets](./src/presets)
* [Custom Logger](./src/customlogger)
* [Global Logger](./src/globallogger)
* [Custom Encoders](./src/customencoder)
* [Sampler](./src/sampler)
* [Hierarchical Logging](./src/hierarchical)

## Attributions

[Forked](https://guides.github.com/activities/forking/), [pull requests](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/about-pull-requests) merged, and many refinements and changes made from the [original zap-examples project](https://github.com/sandipb/zap-examples).