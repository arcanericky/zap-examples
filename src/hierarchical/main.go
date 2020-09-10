package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type level zapcore.Level
type logger struct {
	logger  *zap.Logger
	encoder zapcore.Encoder
	writer  zapcore.WriteSyncer
}

func main() {
	fmt.Print(`*** Create multiple hierarchy of loggers on raw logger ***
You probably need to wrap SetLevel to deal with more level setting/resetting.
`)

	encoderConfig := zap.NewProductionEncoderConfig()
	atomLevel := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	writer := zapcore.Lock(os.Stdout)

	core := zapcore.NewCore(encoder, writer, &atomLevel)
	coreLogger := zap.New(core)
	parentLogger := logger{coreLogger, encoder, writer}
	parentLogger.debug("Should not print")
	fmt.Printf("parent child loggers are really copying parent logger then replace settings")
	childLogger1 := parentLogger.clone()
	childLogger1.debug("This child still inherits level from parent logger. Should not print")
	childLogger2 := parentLogger.newLevel(zapcore.DebugLevel)
	childLogger2.debug("childlogger2 decends from parentLogger with Debug enabled. Debug Should print")
	childLogger3 := childLogger2.clone()
	childLogger3.debug("This child still inherits level from childLogger2. Debug Should print")
	childLogger4 := parentLogger.clone()
	childLogger4.debug("This child still inherits level from original parentLogger. Debug Should not print")
	// {"level":"debug","ts":1527539781.3316631,"msg":"childlogger2 decends from parentLogger with Debug enabled. Debug Should print"}
	// {"level":"debug","ts":1527539781.3317158,"msg":"This child still inherits level from childLogger2. Debug Should print \n"}
}

func (l *logger) newLevel(level zapcore.Level) *logger {
	newLevel := zapcore.Level(level)
	newLogger := l.logger.WithOptions(
		zap.WrapCore(
			func(zapcore.Core) zapcore.Core {
				return zapcore.NewCore(l.encoder, l.writer, newLevel)
			}))
	return &logger{newLogger, l.encoder, l.writer}
}

func (l *logger) debug(msg string) {
	l.logger.Debug(msg)
}

func (l *logger) info(msg string) {
	l.logger.Info(msg)
}

func (l *logger) error(msg string) {
	l.logger.Error(msg)
}

func (l *logger) fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l *logger) clone() *logger {
	copy := *l
	return &copy
}
