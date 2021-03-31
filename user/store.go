package user

import (
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
	if !row.Next() {
		return User{}, err
	}

	err = row.StructScan(&user)
	if err != nil {
		return User{}, err
	}
	return *user, nil
}

func (s *store) SaveUserByID(u User) error {
	_, err := s.DB.NamedExec(`INSERT INTO users(id, data) VALUES(:id, :data)`,
		map[string]interface{}{
			"id":   u.ID,
			"data": u.Data.LastName + u.Data.FirstName + u.Data.Interests,
		})
	// fmt.Println(u)
	if err != nil {
		return err
	}
	return nil
}
