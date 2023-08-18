package logger

import "github.com/sirupsen/logrus"

func InitLogrus() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
}
