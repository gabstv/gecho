package gecho_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/gabstv/gecho"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type UserGet struct {
	ID   string `param:"id" query:"id"`
	Name string `query:"name"`
}

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func TestGet(t *testing.T) {
	e := echo.New()

	johnDoe := User{
		ID:        "1",
		Name:      "John Doe",
		CreatedAt: time.Date(2020, 5, 10, 13, 11, 45, 0, time.UTC),
	}
	janeDoe := User{
		ID:        "2",
		Name:      "Jane Doe",
		CreatedAt: time.Date(2018, 2, 22, 22, 0, 12, 0, time.UTC),
	}

	getUserShared := func(c echo.Context, req UserGet) (User, error) {
		switch req.ID {
		case "1":
			return johnDoe, nil
		case "2":
			return janeDoe, nil
		}
		return User{}, echo.ErrNotFound
	}

	gecho.Get[UserGet, User](e, "/user/:id", getUserShared)
	gecho.Get[UserGet, User](e, "/user", getUserShared)

	go e.Start(":7000")
	shutdctx, cf := context.WithTimeout(context.Background(), time.Second*10)
	defer cf()
	defer e.Shutdown(shutdctx)

	t.Run("john doe", func(t *testing.T) {
		resp, err := http.Get("http://localhost:7000/user/1")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		defer resp.Body.Close()
		var usr User
		assert.NoError(t, json.NewDecoder(resp.Body).Decode(&usr))
		assert.Equal(t, johnDoe, usr)
	})
	t.Run("jane doe", func(t *testing.T) {
		resp, err := http.Get("http://localhost:7000/user/2")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		defer resp.Body.Close()
		var usr User
		assert.NoError(t, json.NewDecoder(resp.Body).Decode(&usr))
		assert.Equal(t, janeDoe, usr)
	})
	t.Run("404", func(t *testing.T) {
		resp, err := http.Get("http://localhost:7000/user/3")
		assert.NoError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}
