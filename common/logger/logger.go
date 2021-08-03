package logger

import (
	"io/fs"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// The Logger structure
type Logger struct {
	logger BaseLogger
	config *Config
}

// Debug method
func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

// Debugf method
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// Info method
func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

// Infof method
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Debugf method
func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

// Errorf method
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// Fatal method
func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Fatalf method
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

// NewLogger returns a new logger object
func NewLogger() *Logger {
	log := logrus.New()
	config := NewConfig()
	perm := 0600
	file, err := os.OpenFile(
		config.FilePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		fs.FileMode(perm),
	)
	if err != nil {
		panic(errors.Wrap(err, "logger"))
	}
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.Out = file
	if config.IsDebug {
		log.SetLevel(logrus.DebugLevel)
	}
	return &Logger{logger: log, config: config}
}
