package dbconfig

import (
	"fmt"
	"os"

	"github.com/ghgsnaidu/go-db-demo/pkg/models"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

// "root:root@tcp(172.17.0.2:8080)/mydb"
func connectToDB(config *Config) *gorm.DB {
	mysqlpath := config.username + ":" + config.pass + "@tcp(localhost:8080)/" + config.dbname
	if db != nil {
		return db
	}
	dburl := os.Getenv("DBURL")
	if dburl == "" {
		dburl = mysqlpath
	}
	db_, err := gorm.Open("mysql", dburl)
	if err != nil {
		fmt.Println("error connecting to database")
	} else {
		fmt.Println("connected to db")
	}
	return db_
}

func InitDB(config *Config) *gorm.DB {
	db = connectToDB(config)
	db.AutoMigrate(&models.User{})
	return db
}
