package user

import (
	"encoding/json"
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

func (s *store) UpdateUserByID(id int, field, value string, user User) error {
	b := []byte(user.Data)
	var d Data
	err := json.Unmarshal(b, &d)
	if err != nil {
		return err
	}
	switch field {
	case "{first_name}":
		d.FirstName = value[1 : len(value)-1]
	case "{last_name}":
		d.LastName = value[1 : len(value)-1]
	case "{interests}":
		d.Interests = value[1 : len(value)-1]
	}

	str, err := json.Marshal(d)
	if err != nil {
		return err
	}
	fmt.Println(string(str))
	_, err = s.DB.NamedExec(`UPDATE users SET data=:data WHERE id=:id`,
		map[string]interface{}{
			"id":   id,
			"data": string(str),
		})
	if err != nil {
		return err
	}
	return nil
}
