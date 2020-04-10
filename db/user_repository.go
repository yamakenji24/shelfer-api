package db

import "github.com/yamakenji24/shelfer-api/models"

type UserRepository struct {
	SqlHandler
}

func (repository *UserRepository) Store(u models.User) (id int, err error) {
	user := models.User{
		Username: u.Username,
		Password: u.Password,
	}
	if result := repository.Create(&user); result.Error != nil {
		err = result.Error
		return
	}
	id = int(user.ID)
	return
}

func (repository *UserRepository) Update(u models.User) (id int, err error) {
	if result := repository.Save(&u); result.Error != nil {
		err = result.Error
	}
	id = int(u.ID)
	return
}

func (repository *UserRepository) FindByUser(username string) (user models.User, err error) {
	if result := repository.Where("username = ?", username).First(&user); result.Error != nil {
		err = result.Error
		return
	}
	return
}
