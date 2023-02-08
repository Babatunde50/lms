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

	r.GET("/api/books", api.GetBooks)
	r.GET("/api/books/:isbn", api.GetBook)
	r.POST("/api/books", api.AddBook)

	r.POST("/api/auth/signup", api.SignUp)
	r.POST("/api/auth/login", api.Login)

	return r
}

// member signup
// member login
// borrow book
// return book
// reserve book
// review book
// recommend book
