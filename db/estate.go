package db

import (
	"log"

	"gorm.io/datatypes"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Estate struct {
	Id          string `gorm:"primaryKey"`
	Username    string
	Description string
	Type        string
	Price       string
	DownPayment string
	Sittingroom string
	Bedroom     string
	Hybrid      string
	Bathroom    string
	Toilet      string
	Active      string
	Likes       string
	Media       datatypes.JSON
	Comments    datatypes.JSON
}

type Comments struct {
	Id       []string
	Username string
	Picture  string
	Likes    string
	Text     string
	Media    string
}

func EstatesDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("tmp/general.db"), &gorm.Config{})
	if err == nil {
		err = db.AutoMigrate()
	} else {
		log.Println("error creating databse")
	}
	return db
}
