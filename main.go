package main

import (
	"log"
	"os"

	db "github.com/GaurKS/book-api/config"
	"github.com/GaurKS/book-api/routes"
	"github.com/GaurKS/book-api/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB := db.Init()
	h := services.New(DB)
	router := gin.Default()
	r := router.Group("/api")
	{
		routes.Router(r.Group("/v1"), &h)
	}

	router.Run(os.Getenv("LOCAL_PORT"))
}
