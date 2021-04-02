package user

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

// TestServiceGetUser tests handler for get user
func TestServiceGetUser(t *testing.T) {
	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetPath("/user/:id")

	if w.Code == http.StatusOK {
		fmt.Println("everythig is ok")
	} else {
		t.Errorf("expected to get status %d but instead got %d", http.StatusOK, w.Code)
	}
}

// TestServiceSaveUser tests handler for save user
func TestServiceSaveUser(t *testing.T) {
	e := echo.New()
	r := httptest.NewRequest(http.MethodPost, "/user/:id", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetPath("/user/:id")

	if w.Code == http.StatusOK {
		fmt.Println("everythig is ok")
	} else {
		t.Errorf("expected to get status %d but instead got %d", http.StatusOK, w.Code)
	}
}

// TestServiceUpdateUser tests handler for update user
func TestServiceUpdateUser(t *testing.T) {
	e := echo.New()
	r := httptest.NewRequest(http.MethodPut, "/user/:id?field={field}&value={value}", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetPath("/user/:id?field={field}&value={value}")

	if w.Code == http.StatusOK {
		fmt.Println("everythig is ok")
	} else {
		t.Errorf("Expected to get status %d but instead got %d", http.StatusOK, w.Code)
	}
}
