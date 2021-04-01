package main

import (
	"fmt"

	"Klara/user"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

const (
	port     = "5432"
	userName = "postgres"
	password = "mypassword"
	host     = "localhost"
	dbname   = "postgres"
)

func main() {
	dbStr := "postgres://" + userName + ":" + password + "@" + host + ":" + port + "/" + dbname + "?" + "sslmode=disable"
	db, err := sqlx.Connect("postgres", dbStr)
	if err != nil {
		fmt.Println("failed to connect to db")
		return
	}

	m, err := migrate.New(
		"file://migrates",
		dbStr)

	if err != nil {
		fmt.Println("failed to make migrate", err)
		return
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("failed to m.Up() in migrate")
		return
	}

	e := echo.New()

	userStore := user.NewStore(db)
	user.NewService(userStore, db, e)

	e.Start(":8000")
}
