package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	db "github.com/GaurKS/book-api/config"
	"github.com/GaurKS/book-api/routes"
	"github.com/GaurKS/book-api/services"
	"github.com/gin-gonic/gin"
)

func main() {
	DB := db.Init()
	h := services.New(DB)
	router := gin.Default()
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})
	r := router.Group("/api")
	{
		routes.Router(r.Group("/v1"), &h)
	}

	// get the port
	port, err := getPort()
	if err != nil {
		log.Fatal(err)
	}

	router.Run(port)
}

func getPort() (string, error) {
	// the PORT is supplied by Heroku
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
