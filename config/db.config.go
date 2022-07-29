package config

import (
	"log"
	"os"

	"github.com/GaurKS/book-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Book{Isbn: 1118, Title: "The International Enc", Publisher: "Wiley Blackwell", Published_year: 2017, Synopsis: "The International Encyclopedia"})
	log.Println("Database migration completed!")
	return db
}
