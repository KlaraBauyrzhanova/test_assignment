package user

import (
	"encoding/json"
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
	e.PUT("/user/:id?field={field}&value={value}", u.updateUser)

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
	d := &Data{}
	u.ID = id
	if err := c.Bind(&d); err != nil {
		return c.JSON(400, err)
	}
	b, err := json.Marshal(d)
	if err != nil {
		return c.JSON(500, err)
	}
	str := string(b)
	u.Data = str
	err = s.store.SaveUserByID(*u, str)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(201, u)
}

func (s *Service) updateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(404, "bad id request")
	}
	field := c.QueryParam("field")
	value := c.QueryParam("value")
	fmt.Println(id, field, value)
	return nil
}
