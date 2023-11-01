package db

import (
	"gorm.io/datatypes"
	"github.com/joho/godotenv"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Username  string `gorm:"primaryKey"`
	Password  string
	Email     string
	Personal  string
	Type      string
	Picture   string
	Phone     string
	Bio       string
	Portfolio datatypes.JSON
	Messages  datatypes.JSON
}

type Messege struct {
	Time     string
	Username string
	Picture  string
	Text     string
	Media    string
}

func UsersDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Unable To Access Environment Data")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_DSN")), &gorm.Config{})

	if err != nil {
		panic("Unable To Open Database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Unable To Migrate Database")

	}

	return db
}
