package controller

import "net/http"

func UserLogout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Options.MaxAge = -1
	session.Save(r, w)
	webresponse("done", nil, nil, w)
}
