package store

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
