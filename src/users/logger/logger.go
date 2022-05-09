package logger

import (
	"io"

	"github.com/gin-gonic/gin"
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

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		info := logrus.Fields{
			"ip":     c.ClientIP(),
			"uri":    c.Request.RequestURI,
			"method": c.Request.Method,
			"status": c.Writer.Status(),
		}

		Entry.WithFields(info).Debugln("Access")
	}
}
