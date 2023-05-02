package controller

import (
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	objJSON := make(map[string]interface{})
	objJSON["time"] = "2021-07-17 16:37:05"
	objJSON["item1"] = "value1"
	objJSON["item2"] = 1
	objJSON["item3"] = 1.234
	b, _ := json.Marshal(objJSON)
	w.Write(b)
}
