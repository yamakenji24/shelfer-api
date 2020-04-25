package db

import (
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
	dbuser := os.Getenv("dbuser")
	dbpass := os.Getenv("dbpass")
	dbname := os.Getenv("dbname")
	db, err = gorm.Open(
		"postgres",
		"user="+dbuser+" dbname="+dbname+" password="+dbpass+" sslmode=disable",
	)
	if err != nil {
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
