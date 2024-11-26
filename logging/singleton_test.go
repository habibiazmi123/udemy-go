package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello World")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Oke")
}
