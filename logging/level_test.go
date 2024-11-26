package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("This is a trace")
	logger.Debug("This is a debug")
	logger.Info("This is a info")
	logger.Warn("This is a warn")
	logger.Error("This is a error")
}

func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is a trace")
	logger.Debug("This is a debug")
	logger.Info("This is a info")
	logger.Warn("This is a warn")
	logger.Error("This is a error")
}
