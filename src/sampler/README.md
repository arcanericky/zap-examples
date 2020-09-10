# Log Sampler

The Log Sampler attempts to reduce CPU and I/O pressure by recording only a subset of entries and dropping duplicate log entries. Log entries having the same log level and message content are considered duplicates. 

The logger maintains a separate bucket for each log entry. At each tick, the Sampler will emit the first N initial logs in each bucket and every Mth log thereafter. Sampling loggers are safe for concurrent use. In this example, we will emit the first 5 messages then one every each 100 messages thereafter.

```console
$ go run src/sampling/main.go
Log sampling to reduce the pressure on I/O and CPU by combining log entries.
This example uses the built in sampler. You probably need to wrap the whole
zapcore Sampler public methods if you need to write our own custom sampler.

We will first emit the first 5 messages then one every 100 messages
thereafter.

{"level":"info","ts":1599770059.1435816,"msg":"test at info","n":1}
{"level":"info","ts":1599770059.1437137,"msg":"test at info","n":2}
{"level":"info","ts":1599770059.1437554,"msg":"test at info","n":3}
{"level":"info","ts":1599770059.1439047,"msg":"test at info","n":4}
{"level":"info","ts":1599770059.144043,"msg":"test at info","n":5}
{"level":"info","ts":1599770059.1443965,"msg":"test at info","n":105}
```


