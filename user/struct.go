package user

// data is a struct of user data
type Data struct {
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Interests string `json:"interests" db:"interests"`
}

type User struct {
	ID   int
	Data string `db:"data"`
}

var Datas []Data
