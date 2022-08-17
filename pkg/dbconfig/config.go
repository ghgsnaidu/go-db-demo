package dbconfig

import (
	"fmt"

	"github.com/ghgsnaidu/go-db-demo/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db        *gorm.DB
	username  string = "root"
	pass      string = "Gireesh@2000"
	dbname    string = "dbdemo"
	mysqlpath string = username + ":" + pass + "@/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
)

func connectToDB() *gorm.DB {
	if db != nil {
		return db
	}
	db_, err := gorm.Open("mysql", mysqlpath)
	if err != nil {
		fmt.Println("error connecting to database")
	} else {
		fmt.Println("connected to db")
	}
	return db_
}

func Init() {
	db = connectToDB()
	db.AutoMigrate(&models.User{})
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
