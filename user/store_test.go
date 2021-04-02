package user

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

// TestGetUserByID tests func GetUserByID
func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	u := User{}
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"data"}).AddRow(u.Data)
	mock.ExpectQuery(`SELECT (.*) FROM users WHERE`).WithArgs(1).WillReturnRows(rows)
	storeUser := NewStore(sqlxDB)

	if _, err := storeUser.GetUserByID(1); err != nil {
		t.Errorf("error was not expected while geting stats: %s", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	defer sqlxDB.Close()
}

// TestSaveUserByID tests func SaveUserByID
func TestSaveUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	u := User{ID: 1}
	str := "first_name:AAA, last_name:BBB, interests:CCC"
	mock.ExpectExec(`INSERT INTO users`).WithArgs(1, str).WillReturnResult(sqlmock.NewResult(1, 1))
	storeUser := NewStore(sqlxDB)
	if err = storeUser.SaveUserByID(u, str); err != nil {
		t.Errorf("error was not expected while inserting stats: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	defer sqlxDB.Close()
}

// TestUpdateUserByID tests func UpdateUserByID
func TestUpdateUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	str := `{"first_name":"AAA","last_name":"BBB","interests":"CCC"}`
	storeUser := NewStore(sqlxDB)
	u := User{Data: str}
	field := "{first_name}"
	value := "{AAA}"
	mock.ExpectExec(`UPDATE users`).WithArgs(str, 1).WillReturnResult(sqlmock.NewResult(0, 1))
	if _, err = storeUser.UpdateUserByID(1, field, value, u); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	defer sqlxDB.Close()
}
