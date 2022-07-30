package routes

import (
	"github.com/GaurKS/book-api/rate-limiter-pkg/middleware"
	"github.com/GaurKS/book-api/services"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup, h *services.Handler) {
	r.GET("/book/:isbn", middleware.RateLimit, h.GetByIsbn)
	r.POST("/create", h.CreateBook)
}
