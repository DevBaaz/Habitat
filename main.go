package main

import (
	"habitat/db"
	"habitat/hs"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	db.UsersDB()
	db.EstatesDB()

	// hs := hs.HabitatServer{httpAddr: ":8080", grpcAddr: ":50051"}
	// go hs.ServeGrpc()
	// hs.ServeHttp()

	e.POST("/auth/signup", hs.SignUp)
	// e.POST("/user/signup", hs.UpdateProfile)
	// e.POST("/estate/list", hs.List)
	// e.POST("/comment/imagecomment", hs.CommentWithImage)
	// e.POST("/comment/textandimagecomment", hs.CommentWithTextAndImage)
	// e.POST("/message/image", hs.MessageWithImage)
	// e.POST("/message/textandimage", hs.MessageWithTextAndImage)

	e.Use(middleware.CORS())
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(port))
}
