package services

import (
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

}
