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

type CreateBook struct {
	Isbn           int    `json:"isbn"`
	Title          string `json:"title"`
	Publisher      string `json:"publisher"`
	Published_year int    `json:"published_year"`
	Synopsis       string `json:"synopsis"`
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

func (db Handler) CreateBook(c *gin.Context) {
	var newBook CreateBook
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		fmt.Println(err.Error())
		c.Abort()
		return
	}

	book := models.Book{
		Isbn:           newBook.Isbn,
		Title:          newBook.Title,
		Publisher:      newBook.Publisher,
		Published_year: newBook.Published_year,
		Synopsis:       newBook.Synopsis,
	}
	if result := db.DB.Create(&book); result.Error != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": result.Error,
			},
		)
		fmt.Println(result.Error)
		c.Abort()
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"message": "Book created successfully",
		},
	)
}
