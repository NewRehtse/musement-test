// Package logger abstract logrus logger package
package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Logger is the main logger that is abstracted in this package.
var Logger *logrus.Logger
var DebugEnabled = true
var LogFile = ""

func init() {
	loadConfiguration()
}

func loadConfiguration() {
	Logger = logrus.New()

	if "" != LogFile {
		logFile, err := os.OpenFile(LogFile, os.O_WRONLY|os.O_CREATE, 0664)
		if err != nil {
			logFile = nil
		}
		//mw := io.MultiWriter(os.Stdout, logFile)
		Logger.SetOutput(logFile)
	}

	Logger.SetLevel(logrus.InfoLevel)
	Logger.Formatter = &logrus.TextFormatter{DisableColors: true}
}

func ReloadConfiguration() {
	loadConfiguration()
}

func Debugf(format string, args ...interface{}) {
	if DebugEnabled {
		Logger.Debugf(format, args...)
	}
}

func Infof(format string, args ...interface{}) {
	if DebugEnabled {
		Logger.Infof(format, args...)
	}
}

func Error(args ...interface{}) {
	if DebugEnabled {
		Logger.Error(args...)
	}
}

func Errorf(format string, args ...interface{}) {
	if DebugEnabled {
		Logger.Errorf(format, args...)
	}
}

func Fatalf(format string, args ...interface{}) {
	if DebugEnabled {
		Logger.Fatalf(format, args...)
	}
}