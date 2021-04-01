package user

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type store struct {
	DB *sqlx.DB
}

func NewStore(db *sqlx.DB) *store {
	return &store{
		DB: db,
	}
}

func (s *store) GetUserByID(id int) (User, error) {
	u := User{ID: id}
	user := &User{}

	row, err := s.DB.NamedQuery(`SELECT * FROM users WHERE id=:id`, u)
	if err != nil {
		return User{}, err
	}
	if !row.Next() {
		return User{}, errors.New("failed to get author by id")
	}

	err = row.StructScan(&user)
	fmt.Println(user)
	if err != nil {
		return User{}, err
	}

	return *user, nil
}

func (s *store) SaveUserByID(u User, str string) error {

	_, err := s.DB.NamedExec(`INSERT INTO users(id, data) VALUES(:id, :data)`,
		map[string]interface{}{
			"id":   u.ID,
			"data": str,
		})
	if err != nil {
		return err
	}
	return nil
}
