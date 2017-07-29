package controller

import (
	"net/http"

	"github.com/Shivakishore14/Oz-CTF-core/app/model"
)

func ListChallenges(w http.ResponseWriter, r *http.Request) {
	// session, err := store.Get(r, "session-name")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// _, ok := session.Values["username"]
	// if ok == false {
	// 	//http.Redirect(w, r, "/login", http.StatusUnauthorized)
	// 	webresponse("Not Logged in", nil, nil, w)
	// 	return
	// }

	challenges := make([]model.Challenge, 0, 100)
	db.Find(&challenges)
	for index, challenge := range challenges {
		challenges[index].Hints = challenge.GetHints(db)
	}
	webresponse("success", nil, challenges, w)

}
