package controller

import (
	"fmt"
	"net/http"

	"github.com/Shivakishore14/Oz-CTF-core/app/model"
)

func ListChallenges(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userName, ok := session.Values["username"]
	if ok == false {
		//http.Redirect(w, r, "/login", http.StatusUnauthorized)
		webresponse("Not Logged in", nil, nil, w)
		return
	}
	user := model.User{}
	gdb := db.Where("user_name = ?", userName).First(&user)
	fmt.Println(gdb.Error)
	if user.ID == 0 {
		webresponse("Not Valid Login", nil, nil, w)
		return
	}
	challenges := make([]model.Challenge, 0, 100)
	db.Find(&challenges)
	for index, challenge := range challenges {
		challenges[index].Hints = challenge.GetHints(db)
		challenges[index].Solved = isSolvedByUser(user.ID, challenge.ID)
	}
	webresponse("success", nil, challenges, w)

}

func isSolvedByUser(userid uint, cid uint) bool {
	submissionObj := model.Submission{}
	gobj := db.Where("user_id = ? and challenge_id = ? ", userid, cid).Find(&submissionObj)
	if gobj.Error != nil {
		fmt.Println(gobj.Error)
		return false
	}
	return true
}
