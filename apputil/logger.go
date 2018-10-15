package apputil

import (
	"log"

	"github.com/sirupsen/logrus"
)

var logr = logrus.New()
var logctx *logrus.Entry
var LogWriter = logr.Writer()

func InitLog() {
	logr.Formatter = &logrus.TextFormatter{}
	logr.SetLevel(logrus.DebugLevel)

	logctx = logr.WithFields(logrus.Fields{
		"app": "Pizza Order System",
	})

	log.SetOutput(logr.Writer())
}

//Debug is used to log Debug level Information
func Debug(msg string, data ...interface{}) {
	logctx.WithFields(logrus.Fields{
		"Message": msg,
	}).Debug(data)
}

//Error is used to log Error Info
func Error(msg string, data ...interface{}) {
	logctx.WithFields(logrus.Fields{
		"Message": msg,
	}).Error(data)
}

//Info is used to log Info level logs
func Info(msg string, data ...interface{}) {
	logctx.WithFields(logrus.Fields{
		"Message": msg,
	}).Info(data)
}
