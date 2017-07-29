package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/Shivakishore14/Oz-CTF-core/app/console"
	"github.com/Shivakishore14/Oz-CTF-core/app/controller"
	"github.com/gorilla/mux"
)

//LoadRoutes :for loading routing
func LoadRoutes() {

	address := ":8080"

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/login", controller.UserLogin)
	api.HandleFunc("/signup", controller.UserSignup)
	api.HandleFunc("/listChallenges", controller.ListChallenges)
	api.HandleFunc("/submitFlag", controller.SubmitFlag)
	api.HandleFunc("/logout", controller.UserLogout)

	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	console.PrintSuccess("Listening on " + address)
	log.Fatal(srv.ListenAndServe())
}
