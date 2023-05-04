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

	"github.com/gorilla/mux"
)

func TestCreateUser(t *testing.T) {
	// TODO: mock DB
	entity.NewConnectDB()
	defer entity.CloseDB()

	body := `{"username":"username3","firstName":"user2","lastName":"name3","email":"mailaddress4","password":"password1234","phone":"0123456789","userStatus":0}`
	r := httptest.NewRequest(http.MethodPost, "/v2/user", bytes.NewBuffer([]byte(body)))
	w := httptest.NewRecorder()

	CreateUser(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected status OK: got %v", w.Code)
	}
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	items := model.User{}
	err = json.Unmarshal(data, &items)
	if err != nil {
		t.Errorf("expected error to be nil got %v [%v]", err, string(data))
	}
	if items.Id <= 0 {
		t.Error("Response JSON body - expected: 0 < id")
	}
}

func TestGetUserByName(t *testing.T) {
	// TODO: mock DB
	entity.NewConnectDB()
	defer entity.CloseDB()

	r := httptest.NewRequest(http.MethodGet, "/v2/user/username1", nil)
	w := httptest.NewRecorder()

	// Router, URL Parameter for UnitTest
	router := mux.NewRouter()
	router.HandleFunc("/v2/user/{username}", GetUserByName)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected status OK: got %v", w.Code)
	}
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	items := model.User{}
	err = json.Unmarshal(data, &items)
	if err != nil {
		t.Errorf("expected error to be nil got %v [%v]", err, string(data))
	}
	if items.Id <= 0 {
		t.Error("Response JSON body - expected: 0 < id")
	}
	expectedUserName := "username1"
	if items.Username != expectedUserName {
		t.Errorf("Response username - expected [%s]  got [%s]", expectedUserName, items.Username)
	}
}
