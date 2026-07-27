package handlers

import (
	logrus "github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	logger.Formatter = &logrus.TextFormatter{
		TimestampFormat:        "01-02 15:04:05",
		ForceColors:            true,
		DisableColors:          false,
		FullTimestamp:          true,
		DisableSorting:         false,
		DisableLevelTruncation: true,
	}
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

		logger.Log(level, logMsg)
	}

}
func Info(msg string) {

	logger.Log(logrus.InfoLevel, msg)
}
