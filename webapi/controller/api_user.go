package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	sql := `INSERT INTO "user" ("username", "firstname", "lastname") VALUES (?,?,?) RETURNING "id"`
	result, err := entity.Engine.SQL(sql, t.Username, t.FirstName, t.LastName).Get(&id)
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
			Id:        item.Id,
			Username:  item.Username,
			FirstName: item.FirstName,
			LastName:  item.LastName,
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

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	log.Printf("filename: [%s]\n", filename)

	// 保存先ディレクトリ
	root := "./data"
	os.MkdirAll(root, os.ModePerm)

	filePath := filepath.Join(root, filename)
	log.Printf("filePath: [%s]\n", filePath)

	// ファイルへストリーム書き込み
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Failed to create file: %s, error: %v", filePath, err)
		return
	}
	defer out.Close()

	// リクエストボディをそのままコピー
	size, err := io.Copy(out, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Failed to write file: %s, error: %v", filePath, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"filename":"%s","size":%d}`, filename, size)
	log.Printf("Uploaded file: %s (%d bytes)\n", filename, size)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	log.Printf("filename: [%s]\n", filename)

	root := "./data"
	filePath := filepath.Join(root, filename)
	log.Printf("filePath: [%s]\n", filePath)

	// ファイル存在チェック
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "file not found", http.StatusNotFound)
		log.Fatalf("File not found: %s", filePath)
		return
	}

	// Content-Type をバイナリに設定
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)

	// http.ServeFile が最速
	http.ServeFile(w, r, filePath)
}
