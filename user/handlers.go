package user

import (
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	"time"
	"log"
	"encoding/json"
)

var UserDBName = "users"

// direct handlers

func UserGet(req *http.Request, w http.ResponseWriter) {

	vars := mux.Vars(req)
	userId := vars["userId"]
}

func UserSignUp(req *http.Request, w http.ResponseWriter) {

}

func UserDelete(req *http.Request, w http.ResponseWriter) {

}

func Login(req *http.Request, w http.ResponseWriter) {

}

func Logout(req *http.Request, w http.ResponseWriter) {

}

func GenerateResetPasswordToken(req *http.Request, w http.ResponseWriter) {

}

func OneTimeAction(req *http.Request, w http.ResponseWriter) {

}

// database services

func UserDBLookup(userId string) (string, bool) {
	db, _ := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/" + UserDBName)
	rows, _ := db.Query("SELECT user_id, email, last_login, registered FROM users WHERE user_id = ?", userId)
	defer db.Close()
	defer rows.Close()
	for rows.Next() {
		var (
			id		int64
			Email		string
			LastLogin	time.Location
			Registered	time.Location
		)
		err := rows.Scan(&id, &Email, &LastLogin, &Registered)
		if err != nil {
			log.Fatal(err)
		}
		return json.Marshal(User{id, Email, LastLogin, Registered}), false
	}
	return "User with ID " + userId + " was not found", true
}