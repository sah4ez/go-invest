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
	db      *sql.DB
	connStr string
)

func init() {
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDb := os.Getenv("POSTGRES_DB")

	connStr = "user=" + pgUser + " dbname=" + pgDb + " password=" + pgPassword + " sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	checkError(err)

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	checkError(err)

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		pgDb,
		driver)
	checkError(err)

	err = m.Up()
	checkError(err)
}

func main() {
	db, err := sql.Open("postgres", connStr)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	err = loader.Securities()
	checkError(err)

	fmt.Println("final")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
