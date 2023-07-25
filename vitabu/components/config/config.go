package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var MyDB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open("sqlite3", "vitabu.db")
	if err != nil {
		panic(err)
	}

	MyDB = db
}
