package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"webapi/entity"
	"webapi/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	entity.NewConnectDBWithDB(db)
	defer entity.CloseDB()

	var ExpectedID int64 = 11
	rows := sqlmock.NewRows([]string{"id"}).AddRow(ExpectedID)
	mock.ExpectQuery(`^INSERT INTO "user" (.+)$`).WillReturnRows(rows)

	app := fiber.New()
	app.Post("/v2/user", CreateUser)

	body := `{"username":"username3","firstName":"user2","lastName":"name3"}`
	req := httptest.NewRequest("POST", "/v2/user/", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.Nil(t, err)

	data, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	items := model.User{}
	err = json.Unmarshal(data, &items)
	assert.Nil(t, err)
	assert.Equal(t, ExpectedID, items.Id)
}

func TestGetUserByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	entity.NewConnectDBWithDB(db)
	defer entity.CloseDB()

	rows := sqlmock.NewRows([]string{"id", "username", "firstName", "lastName"}).
		AddRow(13, "username3", "user2", "name3")
	mock.ExpectQuery(`^SELECT "id"(.+)$`).WillReturnRows(rows)

	app := fiber.New()
	app.Get("/v2/user/:username", GetUserByName)

	req := httptest.NewRequest("GET", "/v2/user/username1", nil)

	resp, err := app.Test(req, -1)
	assert.Nil(t, err)

	data, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	items := model.User{}
	err = json.Unmarshal(data, &items)
	assert.Nil(t, err)
	assert.Equal(t, int64(13), items.Id)
	assert.Equal(t, "username3", items.Username)
}
