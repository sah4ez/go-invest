package migrations

import (
	"database/sql"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	"github.com/sah4ez/go-invest/config"
	"github.com/sah4ez/go-invest/logger"
)

var (
	log = logger.NewLogger("migration")
)

func Up(cfg *config.AppConfig) {
	db, err := sql.Open("postgres", cfg.ConnStr)
	if checkError(err, "Cannot get sql.DB") {
		return
	}
	log.Infoln("DB... \tOk")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if checkError(err, "Cannot get instance DB") {
		return
	}
	log.Infoln("New instance... \tOk")

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		cfg.PgDb,
		driver)
	if checkError(err, "Cannot get NewInstance with migration files") {
		return
	}
	log.Infoln("Migration... \tOk")

	err = m.Up()
	if checkError(err, "Cannot update migration") {
		return
	}
	log.Infoln("Migration update... \tOk")
}

func checkError(err error, msg string) bool {
	if err != nil {
		logger.WithError(log, err).Errorln(msg)
	}
	return err != nil
}
