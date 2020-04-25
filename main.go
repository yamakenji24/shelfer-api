package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yamakenji24/shelfer-api/controllers"
	"github.com/yamakenji24/shelfer-api/db"
)

func main() {
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Println(err)
	}
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/login", func(c *gin.Context) {
		controllers.Login(c)
	})
	router.POST("/storage", func(c *gin.Context) {
		controllers.StoreBook(c)
	})

	db.Connection()

	router.Run()

	db.Close()
}
