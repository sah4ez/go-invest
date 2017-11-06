package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"github.com/sah4ez/go-invest/loader"
)

var (
	db  *sql.DB
	cfg *appConfig
)

type appConfig struct {
	pgUser     string
	pgPassword string
	pgDb       string
	connStr    string
}

func init() {
	cfg = &appConfig{
		pgUser:     os.Getenv("POSTGRES_USER"),
		pgPassword: os.Getenv("POSTGRES_PASSWORD"),
		pgDb:       os.Getenv("POSTGRES_DB"),
	}
	cfg.connStr = "user=" + cfg.pgUser + " dbname=" + cfg.pgDb + " password=" + cfg.pgPassword + " sslmode=disable"

	db, err := sql.Open("postgres", cfg.connStr)
	fmt.Println("Ok")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	checkError(err)

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		cfg.pgDb,
		driver)
	checkError(err)

	err = m.Up()
	checkError(err)
}

func main() {
	db, err := sql.Open("postgres", cfg.connStr)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	err = loader.Securities(cfg.connStr)
	checkError(err)

	fmt.Println("final")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
