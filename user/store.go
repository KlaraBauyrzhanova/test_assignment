package user

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
)

// store is struct of database
type store struct {
	DB *sqlx.DB
}

// NewStore creates store struct
func NewStore(db *sqlx.DB) *store {
	return &store{
		DB: db,
	}
}

// GetUserByID selects user by ID
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
	if err != nil {
		return User{}, err
	}

	return *user, nil
}

// SaveUserByID creates new user by ID
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

// UpdateUserByID updates user by ID
func (s *store) UpdateUserByID(id int, field, value string, user User) (User, error) {
	b := []byte(user.Data)
	var d Data
	err := json.Unmarshal(b, &d)
	if err != nil {
		return User{}, err
	}
	v := value[1 : len(value)-1]
	switch field {
	case "{first_name}":
		d.FirstName = v
	case "{last_name}":
		d.LastName = v
	case " {interest}":
		if d.Interests == "" {
			d.Interests = v
		} else {
			d.Interests = d.Interests + "," + v
		}
	case "-{interest}":
		if strings.Contains(d.Interests, v) {
			if d.Interests == v {
				d.Interests = ""
			} else {
				arr := strings.Split(d.Interests, ",")
				s := ""
				for i := 0; i < len(arr); i++ {
					if arr[i] == v {
						continue
					}
					if i != len(arr) {
						s += arr[i] + ","
					}
				}
				if len(s) > 0 && s[len(s)-1:] == "," {
					d.Interests = s[:len(s)-1]
				}
			}
		}
	default:
		return User{}, errors.New("no such field")
	}

	str, err := json.Marshal(d)
	if err != nil {
		return User{}, err
	}
	_, err = s.DB.NamedExec(`UPDATE users SET data=:data WHERE id=:id`,
		map[string]interface{}{
			"id":   id,
			"data": string(str),
		})
	if err != nil {
		return User{}, err
	}
	result := User{
		ID:   id,
		Data: string(str),
	}
	return result, nil
}
