package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shivakishore14/Oz-CTF-core/app/console"
	"github.com/Shivakishore14/Oz-CTF-core/app/model"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
)

var database = "ctf"
var user = "test"
var password = "test"

var db *gorm.DB
var store *sessions.CookieStore

func init() {
	var err error

	store = sessions.NewCookieStore([]byte("secret-key-to-be-changed"))
	if db, err = gorm.Open("mysql", user+":"+password+"@/"+database+"?charset=utf8&parseTime=True&loc=Local"); err != nil {
		console.PrintError("Error Connecting to database")
	}

	if db.HasTable(&model.User{}) == false {
		db.CreateTable(&model.User{})
	}
	if db.HasTable(&model.Challenge{}) == false {
		db.CreateTable(&model.Challenge{})
	}
	if db.HasTable(&model.Hint{}) == false {
		db.CreateTable(&model.Hint{})
	}
	if db.HasTable(&model.Submission{}) == false {
		db.CreateTable(&model.Submission{})
	}

	fmt.Println("DEBUG : Init Contollers Done")
	//f := model.User{UserName: "max", Password: "kishore", Email: "maxeeee@gmail.com", Experiance: "11", Role: "employee", Name: "max", Status: "enabled", ProfilePic: "/img/c.jpg"}
	//a := db.Create(&f)
	//fmt.Println(a.Error)
	//db.CreateTable(model.UserLoginTimestamp{})
	//fmt.Println(reflect.TypeOf(a))
}

func webresponse(msg string, err error, data interface{}, w http.ResponseWriter) (string, error) {
	obj := model.WebResponse{}
	obj.Message = msg
	obj.Error = err
	obj.Data = data

	var resErr error
	var resTxt string

	if jsonData, e := json.Marshal(obj); e != nil {
		resErr = e
	} else {
		resTxt = string(jsonData)
	}
	fmt.Fprint(w, resTxt)
	return resTxt, resErr
}
