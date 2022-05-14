package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	Formatter logrus.Formatter
	Data      logrus.Fields
	Level     logrus.Level
	Output    io.Writer
	Caller    bool
}

var (
	Entry *logrus.Entry
)

func New(config LoggerConfig) {
	logger := logrus.New()
	logger.SetFormatter(config.Formatter)
	logger.SetLevel(config.Level)
	logger.SetOutput(config.Output)
	logger.SetReportCaller(config.Caller)

	entry := logrus.NewEntry(logger)
	entry.Data = config.Data
	Entry = entry
}
