package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	Formatter logrus.Formatter
	Data      logrus.Fields
	Level     logrus.Level
}

var (
	Entry *logrus.Entry
)

func init() {
	config := LoggerConfig{
		Formatter: &logrus.JSONFormatter{},
		Data:      logrus.Fields{"service": "users"},
		Level:     logrus.DebugLevel,
	}
	New(config)
}

func New(config LoggerConfig) {
	logger := logrus.New()
	logger.SetFormatter(config.Formatter)
	logger.SetLevel(config.Level)

	entry := logrus.NewEntry(logger)
	entry.Data = config.Data
	Entry = entry
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		info := logrus.Fields{
			"IP":     c.ClientIP(),
			"URI":    c.Request.RequestURI,
			"method": c.Request.Method,
			"code":   c.Writer.Status(),
		}

		Entry.WithFields(info).Debugln("Access")
	}
}
