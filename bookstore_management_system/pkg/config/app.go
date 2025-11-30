package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("sqlite3", ":memory") // update db connection to point to the databse for learning purpose it is set to in memory db
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}