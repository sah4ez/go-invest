package migrations

import (
	"database/sql"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database"
	"github.com/mattes/migrate/database/postgres"
	"github.com/sah4ez/go-invest/config"
	"github.com/sah4ez/go-invest/logger"
)

var (
	log  = logger.NewLogger("migration")
	pgDb = "postgres"
)

func Up(cfg *config.AppConfig) {
	pgDb = cfg.PgDb
	db, err := sql.Open("postgres", cfg.ConnStr)
	logger.AddDBHook(log, db, "migration")
	if checkError(err, "Cannot get sql.DB") {
		return
	}
	log.Infoln("DB... \tOk")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if checkError(err, "Cannot get instance DB") {
		return
	}
	log.Infoln("New instance... \tOk")
	m := migrateFromFiles("file://./migrations", driver)
	err = m.Up()
	if checkError(err, "Cannot update migration") {
		return
	}
	log.Infoln("Migration update... \tOk")
}

func migrateFromFiles(path string, driver database.Driver) *migrate.Migrate {
	m, err := migrate.NewWithDatabaseInstance(
		path,
		pgDb,
		driver)
	if checkError(err, "Cannot get NewInstance with migration files") {
		return nil
	}
	log.Infoln("Migration... \tOk")
	return m
}

func checkError(err error, msg string) bool {
	if err != nil {
		logger.WithError(log, err).Errorln(msg)
	}
	return err != nil
}
