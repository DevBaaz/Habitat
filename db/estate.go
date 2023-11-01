package db

import (
	"github.com/joho/godotenv"
	"os"

	"gorm.io/datatypes"

	"gorm.io/driver/postgres"
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
	err := godotenv.Load(".env")
	if err != nil {
		panic("Unable To Access Environment Data")
	}

	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_DSN")), &gorm.Config{})

	if err != nil {
		panic("Unable To Open Database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Unable To Migrate Database")

	}
	
	return db
}
