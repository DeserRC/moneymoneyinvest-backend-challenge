package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Logger *zap.Logger

func InitLogger(folder string) error {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		if err := os.MkdirAll(folder, 0777); err != nil {
			return err
		}
	}

	current := time.Now()
	format := current.Format("2006-01-02")

	path := folder + "/" + format + ".log"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	fileWriter := zapcore.AddSync(file)
	consoleWriter := zapcore.AddSync(os.Stdout)

	level := zapcore.InfoLevel

	fileCore := zapcore.NewCore(fileEncoder, fileWriter, level)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleWriter, level)

	tee := zapcore.NewTee(fileCore, consoleCore)

	caller := zap.AddCaller()
	stacktrace := zap.AddStacktrace(zapcore.ErrorLevel)

	Logger = zap.New(tee, caller, stacktrace)
	defer Logger.Sync()

	return nil
}
