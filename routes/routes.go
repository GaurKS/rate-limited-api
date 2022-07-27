package routes

import (
	"github.com/GaurKS/book-api/services"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup, h *services.Handler) {
	r.GET("/book/:isbn", dummy)
}

func dummy(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
