package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Shivakishore14/Oz-CTF-core/app/model"
)

func SubmitFlag(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var t interface{}
		err := decoder.Decode(&t)
		if err != nil {
			log.Println(err)
			webresponse("Not Valid", nil, nil, w)
			return
		}
		defer r.Body.Close()
		fmt.Println("%v", t)
		flagInter := t.(map[string]interface{})["flag"]
		cidInter := t.(map[string]interface{})["cid"]
		if flagInter == nil || cidInter == nil {
			webresponse("Parameter not present", nil, nil, w)
			return
		}
		flag := flagInter.(string)
		cid := cidInter.(float64)
		//fmt.Println(flag, cid)

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
		fmt.Println(gdb.Error)
		if user.ID == 0 {
			webresponse("Not Valid Login", nil, nil, w)
			return
		}
		challenge := model.Challenge{}
		db.Where("id=?", cid).First(&challenge)
		if challenge.ID == 0 {
			webresponse("Not Valid CID", nil, nil, w)
			return
		}

		if challenge.Flag == flag {
			now := fmt.Sprintf("%v", time.Now().UnixNano()/1000000)
			submission := model.Submission{UserId: user.ID, ChallengeId: challenge.ID, Points: challenge.Points, Timestamp: now}
			db.Create(&submission)
			webresponse("valid", nil, nil, w)
			return
		}
	}
	webresponse("Not Valid", nil, nil, w)
}
