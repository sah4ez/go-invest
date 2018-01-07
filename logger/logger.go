package logger

import (
	"log/syslog"
	"os"

	"github.com/sah4ez/go-invest/config"
	"github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

const (
	tag = "tag"
	val = "val"
)

func NewLogger(name string) *logrus.Logger {
	l := logrus.New()
	l.Out = os.Stdout
	logrus.SetFormatter(&logrus.TextFormatter{})
	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		l.SetLevel(logrus.DebugLevel)
	case "info":
		l.SetLevel(logrus.InfoLevel)
	case "warn":
		l.SetLevel(logrus.WarnLevel)
	case "error":
		l.SetLevel(logrus.ErrorLevel)
	case "fatal":
		l.SetLevel(logrus.FatalLevel)
	default:
		l.SetLevel(logrus.PanicLevel)
	}
	hook, err := lSyslog.NewSyslogHook("", "", syslog.LOG_INFO, "")
	if err != nil {
		l.Error("Unable to connect to local syslog daemon")
	} else {
		l.AddHook(hook)
	}
	return l
}

func WithCfg(l *logrus.Logger, cfg *config.AppConfig) *logrus.Entry {
	return l.WithFields(
		logrus.Fields{
			tag: "cfg",
			val: cfg,
		},
	)
}

func WithError(l *logrus.Logger, err error) *logrus.Entry {
	return l.WithFields(
		logrus.Fields{
			tag: "error",
			val: err.Error(),
		},
	)
}
