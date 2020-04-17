package userservice

import (
	"github.com/yamakenji24/shelfer-api/db"
	"github.com/yamakenji24/shelfer-api/models"
)

func CreateModel(u models.User) (err error) {
	user := models.User{
		Username: u.Username,
		Password: u.Password,
	}
	db := db.GetDB()

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func FindByUser(username string) (user models.User, err error) {
	db := db.GetDB()
	if result := db.Where("username = ?", username).First(&user); result.Error != nil {
		err = result.Error
		return
	}
	return
}
