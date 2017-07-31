package controller

import (
	"fmt"
	"net/http"

	"github.com/Shivakishore14/Oz-CTF-core/app/model"
)

func UserDetails(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userName, ok := session.Values["username"]
	if ok == false {
		webresponse("Not Valid", nil, nil, w)
		return
	}
	user := model.User{}
	gdb := db.Where("user_name = ?", userName).First(&user)
	if gdb.Error != nil {
		fmt.Println(gdb.Error)
		webresponse("Not Valid", nil, nil, w)
		return
	}
	user.Password = ""
	webresponse("success", nil, user, w)
}
