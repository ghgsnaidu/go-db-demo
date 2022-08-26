package services

import (
	"fmt"

	"github.com/ghgsnaidu/go-db-demo/pkg/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitServices(db1 *gorm.DB) {
	db = db1
}

func GetUsers() []models.User {
	fmt.Println("getting all users")
	var users []models.User
	db.Find(&users)
	return users
}

func GetUser(id int64) (models.User, bool) {
	var user models.User
	db_ := db.Where("Id=?", id).Find(&user)
	err := true
	if db_ == nil {
		err = false
	}
	return user, err
}

func AddUser(id int64, user models.User) bool {
	_ = db.Create(&user)
	return true
}
