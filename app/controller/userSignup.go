package controller

import (
	"fmt"
	"net/http"

	"github.com/Shivakishore14/Oz-CTF-core/app/model"
)

func UserSignup(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Println(username, password)
	user := model.User{UserName: username, Password: password, Name: name, Email: email}

	msg, err := user.CreateUser(db)
	if err != nil {
		webresponse(msg, err, nil, w)
	}
	webresponse(msg, nil, nil, w)
}
