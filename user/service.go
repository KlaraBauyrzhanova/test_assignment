package user

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type Service struct {
	store *store
	db    *sqlx.DB
}

func NewService(store *store, db *sqlx.DB, e *echo.Echo) *echo.Echo {
	u := Service{
		store: store,
		db:    db,
	}

	e.GET("/user/:id", u.getUser)
	e.POST("/user/:id", u.saveUser)

	return e
}

func (s *Service) getUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	result, err := s.store.GetUserByID(id)
	if err != nil {
		c.JSON(500, err)
	}
	return c.JSON(200, result)
}

func (s *Service) saveUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(404, err)
	}
	u := &User{}
	u.ID = id
	if err := c.Bind(&u.Data); err != nil {
		return c.JSON(400, err)
	}
	fmt.Println(u)
	err = s.store.SaveUserByID(*u)
	if err != nil {
		fmt.Println("zdes' owibka", err)
		return c.JSON(500, err)
	}
	return c.JSON(201, u)
}
