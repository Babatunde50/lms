package router

import (
	"github.com/Babatunde50/lms/pkg/token"
	"github.com/Babatunde50/lms/router/api"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/api/auth/signup", api.SignUp)
	r.POST("/api/auth/login", api.Login)

	authRoutes := r.Group("/api").Use(authMiddleware(token.JWTTokenMaker))

	authRoutes.GET("/books", api.GetBooks)
	authRoutes.GET("/books/:isbn", api.GetBook)
	authRoutes.POST("/books", api.AddBook)

	return r
}
