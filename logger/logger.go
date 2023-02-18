/*
Package logger has a few pre-built logging options
*/
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// zapConsoleLogger creates a zap console logger
//
//	:param level: The logging level "info", "warn", or "error" default: "info"
func zapConsoleLogger(level string) *zap.Logger {
	switch level {
	case "info":
		levelOpt := zap.IncreaseLevel(zapcore.LevelOf(zapcore.InfoLevel))
		logger, _ := zap.NewProduction(levelOpt)
		return logger
	case "warn":
		levelOpt := zap.IncreaseLevel(zapcore.LevelOf(zapcore.WarnLevel))
		logger, _ := zap.NewProduction(levelOpt)
		return logger
	case "error":
		levelOpt := zap.IncreaseLevel(zapcore.LevelOf(zapcore.ErrorLevel))
		logger, _ := zap.NewProduction(levelOpt)
		return logger
	default:
		levelOpt := zap.IncreaseLevel(zapcore.LevelOf(zapcore.InfoLevel))
		logger, _ := zap.NewProduction(levelOpt)
		return logger
	}
}

// zapFileLogger implements zap logger
//
//	:param path: The path and file name of the log
//	:param level: The logging level "debug", "info", "warn" or "error" default: "info"
//	:param maxSizeMB: The maximum log file size before rotation
//	:param maxAgeDays: The maximum days before file rotation
//	:param maxBackupsTotal: The maximum backups to keep
func zapFileLogger(path string, level string, maxSizeMB int, maxAgeDays int, maxBackupsTotal int) *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	logRotate := lumberjack.Logger{
		Filename:   path,
		MaxSize:    maxSizeMB,
		MaxAge:     maxAgeDays,
		MaxBackups: maxBackupsTotal,
	}
	switch level {
	case "debug":
		defaultLogLevel := zapcore.DebugLevel
		writer := zapcore.AddSync(&logRotate)
		core := zapcore.NewTee(zapcore.NewCore(fileEncoder, writer, defaultLogLevel))
		return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	case "info":
		defaultLogLevel := zapcore.InfoLevel
		writer := zapcore.AddSync(&logRotate)
		core := zapcore.NewTee(zapcore.NewCore(fileEncoder, writer, defaultLogLevel))
		return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	case "warn":
		defaultLogLevel := zapcore.WarnLevel
		writer := zapcore.AddSync(&logRotate)
		core := zapcore.NewTee(zapcore.NewCore(fileEncoder, writer, defaultLogLevel))
		return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	case "error":
		defaultLogLevel := zapcore.ErrorLevel
		writer := zapcore.AddSync(&logRotate)
		core := zapcore.NewTee(zapcore.NewCore(fileEncoder, writer, defaultLogLevel))
		return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	default:
		defaultLogLevel := zapcore.InfoLevel
		writer := zapcore.AddSync(&logRotate)
		core := zapcore.NewTee(zapcore.NewCore(fileEncoder, writer, defaultLogLevel))
		return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	}
}

// NewZapFileLogger creates a zap file logger
//
//	:param path: The path and file name of the log
//	:param level: The logging level "debug", "info", "warn" or "error" default: "info"
//	:param maxSizeMB: The maximum log file size before rotation
//	:param maxAgeDays: The maximum days before file rotation
//	:param maxBackupsTotal: The maximum backups to keep
//
// Example
//
//	logger := NewZapFileLogger("./some-log.log", "info", 5, 30, 3)
//	logger.Info("Some log message")
func NewZapFileLogger(path string, level string, maxSizeMB int, maxAgeDays int, maxBackupsTotal int) *zap.Logger {
	return zapFileLogger(path, level, maxSizeMB, maxAgeDays, maxBackupsTotal)

}

// NewZapConsoleLogger creates a zap console logger
//
//	:param level: The logging level "info", "warn", or "error" default: "info"
//
// Example
//
//	logger := NewZapConsoleLogger("info")
//	logger.Info("Some log message")
func NewZapConsoleLogger(level string) *zap.Logger {
	return zapConsoleLogger(level)

}
