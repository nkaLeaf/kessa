package handlers

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.TextFormatter{
		TimestampFormat:        "01-02 15:04:05",
		ForceColors:            true,
		DisableColors:          false,
		FullTimestamp:          true,
		DisableSorting:         false,
		DisableLevelTruncation: true,
	}
	log.SetLevel(logrus.DebugLevel)
}

func Logger(msg string, level logrus.Level, err error) {
	var logMsg string
	if err != nil {
		switch level {
		case logrus.FatalLevel:
			logMsg = msg + err.Error()
		case logrus.WarnLevel:
			logMsg = msg + err.Error()
		case logrus.PanicLevel:
			logMsg = msg + err.Error()
		case logrus.DebugLevel:
			logMsg = msg + err.Error()
		}
		log.Log(level, logMsg)
		if level == logrus.FatalLevel {
			os.Exit(1)
		}
	}

}
func Info(msg string) {
	log.Log(logrus.InfoLevel, msg)
}
