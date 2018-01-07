package logger

import (
	"database/sql"
	"os"

	"github.com/sah4ez/go-invest/config"
	"github.com/sirupsen/logrus"
	pglogrus "gopkg.in/gemnasium/logrus-postgresql-hook.v1"
)

const (
	tag = "tag"
	val = "val"
)

func NewLogger(name string) *logrus.Logger {
	l := logrus.New()
	l.Out = os.Stdout
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
	return l
}

func AddDBHook(log *logrus.Logger, db *sql.DB, name string) {
	hook := pglogrus.NewHook(db, map[string]interface{}{"Name": name})
	log.Hooks.Add(hook)
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
