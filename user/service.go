package user

import (
	"encoding/json"
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
		return c.String(400, "bad id request")
	}
	result, err := s.store.GetUserByID(id)
	if err != nil {
		c.String(500, "failed to get user by ID")
	}
	return c.JSON(200, result)
}

func (s *Service) saveUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(404, "bad id request")
	}
	u := &User{}
	d := &Data{}
	u.ID = id
	if err := c.Bind(&d); err != nil {
		return c.String(400, "failed to bind")
	}
	b, err := json.Marshal(d)
	if err != nil {
		return c.String(500, "failed to Marshal json")
	}
	str := string(b)
	u.Data = str
	err = s.store.SaveUserByID(*u, str)
	if err != nil {
		return c.String(500, "failed to save user")
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
	user, _ := s.store.GetUserByID(id)
	err = s.store.UpdateUserByID(id, field, value, user)
	if err != nil {
		return c.String(500, "failed to update user by ID")
	}
	return nil
}
