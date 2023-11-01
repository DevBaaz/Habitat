package hs

import (
	"habitat/db"
	"habitat/internal/config"
	"habitat/internal/cryption"
	"habitat/internal/s3"
	"net/http"

	"io"
	"mime/multipart"
	"os"

	// "github.com/aws/aws-sdk-go/aws/service/s3"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  string      `json:"status"`
	Message interface{} `json:"email"`
}

type SignUpRequest struct {
	Username string
	Password string
	Personal string
	Email    string
	Type     string
	Picture  string
	Phone    string
	Bio      string
}

func SignUp(c echo.Context) error {
	DB := db.UsersDB()
	Username := c.FormValue("Username")
	Password := c.FormValue("Password")
	Personal := c.FormValue("Personal")
	Email := c.FormValue("Email")
	Type := c.FormValue("Type")
	Phone := c.FormValue("Phone")
	Bio := c.FormValue("Bio")
	retrieveduser := new(db.User)

	err := DB.Where(&db.User{Username: Username}).Find(&retrieveduser).Error
	if err != nil {
		return c.JSON(http.StatusAccepted, &Response{Status: "Fail", Message: "Unable To Access Application Data"})
	}

	if retrieveduser.Username != "" {
		return c.JSON(http.StatusAccepted, &Response{Status: "Fail", Message: "Username Already In Use, Choose Another"})
	}

	err = godotenv.Load(".env")
	if err != nil {
		return c.JSON(http.StatusAccepted, &Response{Status: "Fail", Message: "Unable To Access Environment Data"})
	}

	encryptedPassword, _ := cryption.EncryptPassword(Password, os.Getenv("MY_SECRET"))
	var image *multipart.FileHeader
	if image, err = c.FormFile("image"); err != nil {
		return c.JSON(http.StatusAccepted, &Response{Status: "Fail", Message: "Server Unable To Access Form File"})
	}

	var src multipart.File
	if src, err = image.Open(); err != nil {
		return c.JSON(http.StatusAccepted, &Response{Status: "Fail", Message: "Server Unable To Image"})
	}
	defer src.Close()

	var filePath *os.File
	if filePath, err = os.Create("/habitat/"); err != nil {
		return c.JSON(http.StatusAccepted, &Response{Status: "Fail", Message: "Error Creating File"})
	}

	if _, err = io.Copy(filePath, src); err != nil {
		return c.JSON(http.StatusAccepted, &Response{Status: "Fail", Message: "Error Creating File"})
	}

	err = s3.UploadObject("myestateapp", "/habitat/", image.Filename, s3.CreateSession(config.AWSConfig{Region: os.Getenv("REGION"), AccessKeyId: os.Getenv("ACCESS_KEY_ID"), AccessKeySecret: os.Getenv("ACCESS_KEY_SECRET")}))
	if err != nil {
		return c.JSON(http.StatusAccepted, &Response{Status: "Fail", Message: "Unable To Copy Uploaded Image"})
	}

	user := &db.User{Username: Username, Password: encryptedPassword, Personal: Personal, Email: Email, Type: Type, Phone: Phone, Bio: Bio, Picture: image.Filename}
	if err := DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusAccepted, &Response{Status: "Fail", Message: "Unable To Create User"})
	}

	userlog := new(db.User)
	if err := DB.Where("username = ?", Username).Find(&userlog).Error; err != nil {
		return c.JSON(http.StatusAccepted, &Response{Status: "Fail", Message: "Unable To ConFirm Data Creation"})
	}

	if userlog.Username != Username {
		return c.JSON(http.StatusAccepted, &Response{Status: "Fail", Message: "Unable To Create  Account"})
	}

	return c.JSON(http.StatusOK, &Response{Status: "Success", Message: "SignUp Successful"})
}
