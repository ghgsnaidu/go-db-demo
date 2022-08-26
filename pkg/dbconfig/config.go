package dbconfig

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	username string
	pass     string
	dbname   string
}

// take input from env
// take url also from env
func NewConfig() *Config {
	return &Config{
		username: "root",
		pass:     "root",
		dbname:   "mydb",
	}
}
