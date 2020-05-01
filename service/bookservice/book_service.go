package bookservice

import (
	"github.com/yamakenji24/shelfer-api/db"
	"github.com/yamakenji24/shelfer-api/models"
)

func CreateBook(b *models.BookParams) (err error) {
	db := db.GetDB()
	tx := db.Begin()

	for _, book := range b.Storage {
		if err := tx.Create(&book).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func RequestAllBook() (book []models.Book, err error) {
	db := db.GetDB()
	if err := db.Find(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}
