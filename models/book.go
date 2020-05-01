package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	BookId    string `gorm:"size:255;not null" json:"book_id"`
	Title     string `gorm:"size:255;not null" json:"title"`
	Authors   string `gorm:"size:255;not null" json:"authors"`
	Descrip   string `json:"descrip"`
	PbDate    string `gorm:"size:255" json:"pbDate"`
	InfoLink  string `json:"infoLink"`
	ImageLink string `json:"imageLink"`
}

type BookParams struct {
	Storage []Book `json:"storage"`
}
