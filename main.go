package main

import (
	"database/sql"
	"flag"
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

func main() {
	var (
		migration = flag.Bool("migration", false, "Run migration scripts for PG")
	)
	flag.Parse()

	cfg = &appConfig{
		pgUser:     os.Getenv("POSTGRES_USER"),
		pgPassword: os.Getenv("POSTGRES_PASSWORD"),
		pgDb:       os.Getenv("POSTGRES_DB"),
	}
	cfg.connStr = "user=" + cfg.pgUser + " dbname=" + cfg.pgDb + " password=" + cfg.pgPassword + " sslmode=disable"
	fmt.Printf("Config: %+v\n", cfg)

	if *migration {
		migrationScripts()
		fmt.Println("Migration applied... \tOk")
	}

	db, err := sql.Open("postgres", cfg.connStr)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	err = loader.Securities(cfg.connStr)
	checkError(err)

	fmt.Println("final")
}

func migrationScripts() {
	db, err := sql.Open("postgres", cfg.connStr)
	if checkError(err) {
		fmt.Println("Cannot get sql.DB")
		return
	}
	fmt.Println("DB... \tOk")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if checkError(err) {
		fmt.Println("Cannot get instance DB")
		return
	}
	fmt.Println("New instance... \tOk")

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		cfg.pgDb,
		driver)
	if checkError(err) {
		fmt.Println("Cannot get NewInstance with migration files")
		return
	}
	fmt.Println("Migration... \tOk")

	err = m.Up()
	if checkError(err) {
		fmt.Println("Cannot update migration")
		return
	}
	fmt.Println("Migration update... \tOk")
}

func checkError(err error) bool {
	if err != nil {
		fmt.Printf("Error: \t%s\n", err.Error())
	}
	return err != nil
}
