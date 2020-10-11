package log

import (
"os"

"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.Out = os.Stdout
	logger.Level = logrus.TraceLevel
	logger.Formatter = &logrus.JSONFormatter{}
}

// Debug msg
func Debug(args ...interface{}) {
	logger.Debugln(args)
}

// Debugf msg
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

// Error msg
func Error(args ...interface{}) {
	logger.Errorln(args)
}

// Error msg
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args)
}

// Info msg
func Info(args ...interface{}) {
	logger.Infoln(args)
}

// Info msg
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args)
}

func Fatal(args ...interface{}) {
	logger.Fatalln(args)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args);
}