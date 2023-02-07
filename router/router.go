package router

import (
	"github.com/Babatunde50/lms/router/api"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/api/books/:id", api.GetBook)

	return r
}
