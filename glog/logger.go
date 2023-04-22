package glog

/*
   @File: logger.go
   @Author: khaosles
   @Time: 2023/4/11 22:16
   @Desc:
*/

import (
	"os"
	"time"

	"github.com/khaosles/gtools/gcfg"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const ONE_DAY = time.Hour * 24

var logger *zap.SugaredLogger

func init() {

	logCfg := gcfg.GCfg.Logging

	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     encodeTime,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	})

	levelConsole := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= levelChoice(logCfg.LevelConsole)
	})

	levelFile := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= levelChoice(logCfg.LevelFile)
	})

	var cores []zapcore.Core
	if logCfg.LogInFile {
		path := logPath()
		hook, _ := rotatelogs.New(
			path+"/%Y-%m-%d.log",
			rotatelogs.WithLinkName(path),
			rotatelogs.WithMaxAge(ONE_DAY*logCfg.MaxHistory),
			rotatelogs.WithRotationTime(ONE_DAY),
			rotatelogs.WithClock(rotatelogs.Local),
		)
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(hook), levelFile))
		if logCfg.LogInConsole {
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), levelConsole))
		}
	} else {
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), levelConsole))
	}
	core := zapcore.NewTee(cores...)
	log := zap.New(core)
	if logCfg.ShowLine {
		log = log.WithOptions(zap.AddCaller())
	}
	logger = log.Sugar()
}

func logPath() string {
	path := gcfg.GCfg.Logging.Path
	if path == "" {
		wk, _ := os.Getwd()
		path = wk + "/logs"
	}
	_ = os.MkdirAll(path, os.ModePerm)
	return path
}

func encodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(gcfg.GCfg.Logging.Prefix + t.In(time.FixedZone("CTS", 8*3600)).Format("2006-01-02 15:04:05.000"))
}

func levelChoice(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	default:
		return zapcore.InfoLevel
	}
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}
