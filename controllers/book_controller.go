package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yamakenji24/shelfer-api/models"
	"github.com/yamakenji24/shelfer-api/service/bookservice"
)

func StoreBook(c *gin.Context) error {
	book := new(models.BookParams)
	if err := c.Bind(book); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}

	err := bookservice.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, struct {
		Status string `json:"status"`
	}{
		Status: "Success Saving books",
	})
	return nil
}
