package store

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

var schema = `
		CREATE TABLE "users" (
			"id" INT PRIMARY KEY,
			"data" VARCHAR
		);`

// data is a struct of user data
type Data struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Interests string `json:"interests" db:"interests"`
}

var Datas []Data

type Store struct {
	DB *sqlx.DB
}

func (s *Store) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}

	u := &Data{ID: id}
	user := &Data{}

	row, err := s.DB.NamedQuery(`SELECT * FROM users WHERE id=:id`, &u)
	fmt.Println(row)
	return c.JSON(200, user)
}

func (s *Store) SaveUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}

	u := &Data{ID: id}
}
