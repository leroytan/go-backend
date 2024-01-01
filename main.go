package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leroytan/go-backend/controllers"
	"github.com/leroytan/go-backend/initializers"
	"github.com/leroytan/go-backend/middleware"
)

func init() {
	initializers.ConnectToDB()
	initializers.SyncDatabase()
	initializers.LoadEnvVariables()
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.POST("/api/posts", controllers.PostsCreate)
	router.PUT("/api/posts/:id", controllers.PostsUpdate)
	router.GET("/api/posts", controllers.PostsAll)
	router.GET("/api/posts/:id", controllers.PostIndex)
	router.DELETE("/api/posts/:id", controllers.PostsSoftDelete)
	router.Run(":3000") // listen and serve on 0.0.0.0:3000
}
