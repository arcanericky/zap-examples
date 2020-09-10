# Simple Logger Presets

## Presets

Zap recommends using presets for the simplest of cases and makes three available:

- [Example](https://pkg.go.dev/go.uber.org/zap?tab=doc#NewExample)
- [Development](https://pkg.go.dev/go.uber.org/zap?tab=doc#NewDevelopment)
- [Production](https://pkg.go.dev/go.uber.org/zap?tab=doc#NewProduction)

This is the output when these presets are implemented in code.

```console
$ go run src/presets/main.go
*** Example Logger

{"level":"debug","msg":"This is a DEBUG message"}
{"level":"info","msg":"This is an INFO message"}
{"level":"info","msg":"This is an INFO message with fields","region":"us-west","id":2}
{"level":"warn","msg":"This is a WARN message"}
{"level":"error","msg":"This is an ERROR message"}
{"level":"dpanic","msg":"This is a DPANIC message"}

*** Development Logger

2020-09-10T17:07:46.303-0500	DEBUG	presets/main.go:32	This is a DEBUG message
2020-09-10T17:07:46.303-0500	INFO	presets/main.go:33	This is an INFO message
2020-09-10T17:07:46.303-0500	INFO	presets/main.go:34	This is an INFO message with fields	{"region": "us-west", "id": 2}
2020-09-10T17:07:46.303-0500	WARN	presets/main.go:35	This is a WARN message
main.main
	/home/user/Dev/zap-examples/src/presets/main.go:35
runtime.main
	/home/user/.gimme/versions/go1.15.1.linux.amd64/src/runtime/proc.go:204
2020-09-10T17:07:46.303-0500	ERROR	presets/main.go:36	This is an ERROR message
main.main
	/home/user/Dev/zap-examples/src/presets/main.go:36
runtime.main
	/home/user/.gimme/versions/go1.15.1.linux.amd64/src/runtime/proc.go:204

*** Production Logger

{"level":"info","ts":1599775666.3032868,"caller":"presets/main.go:49","msg":"This is an INFO message"}
{"level":"info","ts":1599775666.3033223,"caller":"presets/main.go:50","msg":"This is an INFO message with fields","region":"us-west","id":2}
{"level":"warn","ts":1599775666.303335,"caller":"presets/main.go:51","msg":"This is a WARN message"}
{"level":"error","ts":1599775666.3033433,"caller":"presets/main.go:52","msg":"This is an ERROR message","stacktrace":"main.main\n\t/home/user/Dev/zap-examples/src/presets/main.go:52\nruntime.main\n\t/home/user/.gimme/versions/go1.15.1.linux.amd64/src/runtime/proc.go:204"}
{"level":"dpanic","ts":1599775666.3033612,"caller":"presets/main.go:54","msg":"This is a DPANIC message","stacktrace":"main.main\n\t/home/user/Dev/zap-examples/src/presets/main.go:54\nruntime.main\n\t/home/user/.gimme/versions/go1.15.1.linux.amd64/src/runtime/proc.go:204"}

*** Sugared logger

2020-09-10T17:07:46.303-0500	INFO	presets/main.go:65	Info() uses sprint
2020-09-10T17:07:46.303-0500	INFO	presets/main.go:66	Infof() uses sprintf
2020-09-10T17:07:46.303-0500	INFO	presets/main.go:67	Infow() allows tags	{"name": "Legolas", "type": 1}

*** JSON Derived Logger

2020-09-10T17:07:46.303-0500	INFO	/home/user/Dev/zap-examples/src/presets/main.go:112	This should have an ISO8601 based time stamp	{"initFieldKey": "fieldValue"}
2020-09-10T17:07:46.303-0500	WARN	/home/user/Dev/zap-examples/src/presets/main.go:113	This is a WARN message	{"initFieldKey": "fieldValue"}
2020-09-10T17:07:46.303-0500	ERROR	/home/user/Dev/zap-examples/src/presets/main.go:114	This is an ERROR message	{"initFieldKey": "fieldValue"}
main.main
	/home/user/Dev/zap-examples/src/presets/main.go:114
runtime.main
	/home/user/.gimme/versions/go1.15.1.linux.amd64/src/runtime/proc.go:204
2020-09-10T17:07:46.303-0500	INFO	/home/user/Dev/zap-examples/src/presets/main.go:119	Failed to fetch URL.	{"initFieldKey": "fieldValue", "url": "http://example.com", "attempt": 3, "backoff": 1}
```

# Observations

- Both `Example` and `Production` loggers use the [JSON encoder](https://godoc.org/go.uber.org/zap/zapcore#NewJSONEncoder). `Development` uses the [Console](https://godoc.org/go.uber.org/zap/zapcore#NewConsoleEncoder) encoder.
- The [`logger.DPanic()`](https://pkg.go.dev/go.uber.org/zap?tab=doc#Logger.DPanic) function causes a panic in the `Development` logger but not in the `Example` and `Production` loggers.
- The `Development` logger:
    * Adds a stack trace from Warn level and up
    * Prints the package/file/line number
    * Appends extra fields as a json string
    * Level names are uppercase
    * Timestamp is in ISO8601 with seconds
- The `Production` logger:
    * Doesn't log messages at debug level
    * Adds stack trace as a json field for Error, DPanic levels, but not for Warn
    * Adds the caller as a json field
    * Timestamp is in epoch format
    * Level names are lowercase

## The Sugared Logger

The default logger expects structured tags.

```go
logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))
```

This is the fastest option for an application where performance is key. However, for a just [a small additional penalty](https://github.com/uber-go/zap#performance),
which is still slightly better than the standard library, you can use
the [Sugared Logger](https://pkg.go.dev/go.uber.org/zap?tab=doc#SugaredLogger) which uses a reflection based type detection to give you
a simpler syntax for adding tags of mixed types.

```go
slogger := logger.Sugar()
slogger.Info("Info() uses sprint")
slogger.Infof("Infof() uses %s", "sprintf")
slogger.Infow("Infow() allows tags", "name", "Legolas", "type", 1)
```

Output:

```
2020-09-10T17:07:46.303-0500	INFO	presets/main.go:65	Info() uses sprint
2020-09-10T17:07:46.303-0500	INFO	presets/main.go:66	Infof() uses sprintf
2020-09-10T17:07:46.303-0500	INFO	presets/main.go:67	Infow() allows tags	{"name": "Legolas", "type": 1}
```

You can switch from a sugared logger to a standard logger any time using the
[`.Desugar()`](https://pkg.go.dev/go.uber.org/zap?tab=doc#SugaredLogger.Desugar) method on the logger.