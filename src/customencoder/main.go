package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan  2 15:04:05"))
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func main() {
	cfg := zap.Config{
		Encoding:    "console",
		OutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
	}

	fmt.Print("*** Using standard ISO8601 time encoder\n\n")

	// avoiding copying of atomic values
	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)

	logger, _ := cfg.Build()
	logger.Info("This should have an ISO8601 based time stamp")
	logger.Sync()

	fmt.Print("\n*** Using a custom time encoder\n\n")

	cfg.EncoderConfig.EncodeTime = syslogTimeEncoder

	logger, _ = cfg.Build()
	logger.Info("This should have a syslog style time stamp")
	logger.Sync()

	fmt.Print("\n*** Using a custom level encoder\n\n")

	cfg.EncoderConfig.EncodeLevel = customLevelEncoder

	logger, _ = cfg.Build()
	logger.Info("This should have a interesting level name")
	logger.Sync()
}
