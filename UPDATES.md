
This repository is great! It helped me learn the basics of zap immediately. I wanted to show my appreciation by updating and improving the repository.

## Use Go Modules

The examples use `env.sh` and the old `GOPATH` method. Updated the repository to use Go Modules and base it on Go 1.15. This eliminates the need for the `env.sh` script (deleted) and make executing the examples simpler with an easy `go run ...`.


## Use the `Sync()` Method

The zap package states:

```
By default, loggers are unbuffered. However, since zap's low-level APIs allow buffering, calling Sync before letting your process exit is a good habit.
```

I've sprinkled calls to the `Sync()` method at the end of the examples and before loggers are disposed of to show good practice and make all data is written.

## Don't Export Functions

Some functions in the `customerencoder` example were exported. Removed this to conform to good programming practices.

## Document the use of the `ConsoleSeparator`

The `customencoder` example's README contains information about a single call to `enc.AppendString()` because multiple calls cause spacing between the strings. Added an explanation of why this happens.