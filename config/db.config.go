package config

import (
	"log"
	"os"

	"github.com/GaurKS/book-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Book{})
	log.Println("Database migration completed!")
	return db
}
