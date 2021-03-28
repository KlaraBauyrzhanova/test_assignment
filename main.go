package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	s := store.Store{
		DB: db,
	}

	e := echo.New()
	e.GET("/user/:id", s.GetUser)
	e.POST("/user/:id", s.SaveUser)
}
