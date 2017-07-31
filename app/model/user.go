package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

//User details
type User struct {
	gorm.Model
	UserName           string `gorm:"not null;unique;index"`
	Password           string `gorm:"not null"`
	Name               string `gorm:"not null"`
	Email              string `gorm:"not null"`
	StartTime          string `gorm:"not null"`
	LastSubmissionTime string `gorm:"not null"`
}

//IsValidLogin : for checking if credentials are valid
func (user User) IsValidLogin(db *gorm.DB) (bool, string) {
	tempUser := User{}
	db.Where("user_name=?", user.UserName).First(&tempUser)

	if tempUser.Password == user.Password && user.Password != "" {
		//fmt.Println(tempUser.StartTime)
		if tempUser.StartTime == "" {
			//first login
			currentTime := fmt.Sprintf("%v", time.Now().UnixNano()/1000000)
			tempUser.StartTime = currentTime
			fmt.Println("setting time", user.UserName, currentTime)
			db.Save(&tempUser)
		}
		return true, "ok"
	}
	return false, "check credentials"
}

func (user User) RegisterLogin(db *gorm.DB, isValid bool) {

	currentTime := fmt.Sprintf("%v", time.Now().Format(time.RFC850))
	//timestamp := UserLoginTimestamp{UserName: user.UserName, Timestamp: currentTime, IsValid: isValid}
	e := db.Exec("INSERT INTO logins values(?,?,?)", user.UserName, currentTime, isValid)
	fmt.Println(e)
	//if gdb := db.Create(timestamp); gdb.Error != nil {
	//	fmt.Print(gdb.Error)
	//}

}

//CreateUser : for creating a new user
func (user User) CreateUser(db *gorm.DB) (string, error) {
	if gdb := db.Create(user); gdb.Error != nil {
		return "Check Given details", gdb.Error
	}
	return "Created user " + user.UserName, nil
}
