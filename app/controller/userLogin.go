package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Shivakishore14/Oz-CTF-core/app/model"
)

//UserLogin : for Login functionality
func UserLogin(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t model.User
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err)
		webresponse("Not Valid", nil, nil, w)
		return
	}
	defer r.Body.Close()
	log.Println(t.UserName)

	username := t.UserName
	password := t.Password

	fmt.Println(username, password)
	user := model.User{UserName: username, Password: password}
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	check, msg := user.IsValidLogin(db)
	fmt.Println(check, msg)
	if check {
		session.Values["username"] = username
		session.Save(r, w)
	}
	webresponse(msg, nil, nil, w)
}
