package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"webapi/entity"
	"webapi/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
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

	body := `{"username":"username3","firstName":"user2","lastName":"name3"}`
	r := httptest.NewRequest(http.MethodPost, "/v2/user", bytes.NewBuffer([]byte(body)))
	w := httptest.NewRecorder()

	var ExpectedID int64 = 11
	rows := sqlmock.NewRows([]string{"id"}).AddRow(ExpectedID)
	mock.ExpectQuery(`^INSERT INTO "user" (.+)$`).WillReturnRows(rows)

	CreateUser(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
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

	r := httptest.NewRequest(http.MethodGet, "/v2/user/username1", nil)
	w := httptest.NewRecorder()

	rows := sqlmock.NewRows([]string{"id", "username", "firstName", "lastName"}).
		AddRow(13, "username3", "user2", "name3")
	mock.ExpectQuery(`^SELECT "id"(.+)$`).WillReturnRows(rows)

	// Router, URL Parameter for UnitTest
	router := mux.NewRouter()
	router.HandleFunc("/v2/user/{username}", GetUserByName)
	router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	items := model.User{}
	err = json.Unmarshal(data, &items)
	assert.Nil(t, err)
	assert.Equal(t, int64(13), items.Id)
	assert.Equal(t, "username3", items.Username)
}
