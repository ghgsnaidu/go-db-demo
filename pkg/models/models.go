package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	Id    int64  `gorm:"primary_key;auto_increment:false" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
