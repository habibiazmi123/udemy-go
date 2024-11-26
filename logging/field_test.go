package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "Azmi").Info("Hello World!")

	logger.WithField("username", "Azmi").WithField("name", "Azmi").Info("Hello Mi")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "Azmi",
		"name":     "Azmi",
	}).Info("Mantul")
}
