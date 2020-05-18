package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yamakenji24/shelfer-api/models"
)

var (
	db  *gorm.DB
	err error
)

func Connection() {
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	fmt.Println(dbuser, dbpass, dbname)
	db, err = gorm.Open(
		"postgres",
		"user="+dbuser+" dbname="+dbname+" password="+dbpass+" sslmode=disable",
	)
	if err != nil {
		fmt.Println("failed to connect db")
		panic(err)
	}
	autoMigration()
}
func GetDB() *gorm.DB {
	return db
}
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
func autoMigration() {
	db.AutoMigrate(&models.User{}, &models.Book{})
}
