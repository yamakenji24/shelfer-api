package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yamakenji24/shelfer-api/auth"
	"github.com/yamakenji24/shelfer-api/controllers"
	"github.com/yamakenji24/shelfer-api/db"
)

func main() {
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Println(err)
	}
	router := gin.Default()
	//router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
	}))

	router.POST("/login", func(c *gin.Context) {
		controllers.Login(c)
	})
	router.POST("/createuser", func(c *gin.Context) {
		controllers.Create(c)
	})

	sampleGroup := router.Group("/storage", auth.CheckJWTHandler)

	sampleGroup.POST("/save", func(c *gin.Context) {
		controllers.StoreBook(c)
	})

	sampleGroup.GET("", func(c *gin.Context) {
		controllers.RequestStoredBook(c)
	})

	db.Connection()

	router.Run()

	db.Close()
}
