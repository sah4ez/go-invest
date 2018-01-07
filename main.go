package main

import (
	"database/sql"
	"flag"

	_ "github.com/mattes/migrate/source/file"
	"github.com/sah4ez/go-invest/config"
	"github.com/sah4ez/go-invest/loader"
	"github.com/sah4ez/go-invest/logger"
	"github.com/sah4ez/go-invest/migrations"
	"github.com/sirupsen/logrus"
)

var (
	db  *sql.DB
	cfg *config.AppConfig
	log *logrus.Logger
)

func main() {
	var (
		migration = flag.Bool("migration", false, "Run migration scripts for PG")
		all       = flag.Bool("all", false, "Full upgrade finance instruments")
	)
	flag.Parse()

	cfg = config.New()

	log = logger.NewLogger("main")
	logger.WithCfg(log, cfg).Infoln("Log created")

	if *migration {
		migrations.Up(cfg)
		log.Infoln("Migration applied... \tOk")
		return
	}

	db, err := sql.Open("postgres", cfg.ConnStr)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	if *all {
		err = loader.Securities(cfg.ConnStr)
		checkError(err)
	}

	log.Infoln("Final")
}

func checkError(err error) bool {
	if err != nil {
		logger.WithError(log, err)
	}
	return err != nil
}
