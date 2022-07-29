package services

import (
	"fmt"
	"net/http"

	"github.com/GaurKS/book-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{db}
}

func (db Handler) GetByIsbn(c *gin.Context) {
	var users models.Book
	isbn := c.Param("isbn")
	result := db.DB.Where("isbn = ?", isbn).First(&users)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"message":  "SUCCESS",
			"resource": users,
		},
	)
}
