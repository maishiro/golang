package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"webapi/entity"
	"webapi/model"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t model.User
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var id int64 = 0
	sql := `INSERT INTO "user" ("username", "firstname", "lastname", "email", "password", "phone", "userstatus") VALUES (?,?,?,?,?,?,?) RETURNING "id"`
	result, err := entity.Engine.SQL(sql, t.Username, t.FirstName, t.LastName, t.Email, t.Password, t.Phone, t.UserStatus).Get(&id)
	if err != nil {
		log.Printf("err: [%v]\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
	log.Printf("result: [%v]\n", result)
	log.Printf("id: [%v]\n", id)
	t.Id = id

	b, _ := json.Marshal(t)
	w.Write(b)
}

func GetUserByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_name := vars["username"]
	log.Printf("username: [%s]\n", user_name)

	item := entity.User{}
	result, err := entity.Engine.Where("username = ?", user_name).Get(&item)
	if err != nil {
		log.Printf("err: [%v]\n", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	user := model.User{}
	if !result {
		log.Println("Not Found")
	} else {
		user = model.User{
			Id:         item.Id,
			Username:   item.Username,
			FirstName:  item.FirstName,
			LastName:   item.LastName,
			Email:      item.Email,
			Password:   item.Password,
			Phone:      item.Phone,
			UserStatus: item.UserStatus,
		}
	}

	b, err := json.Marshal(user)
	if err != nil {
		log.Printf("err: [%v]\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
