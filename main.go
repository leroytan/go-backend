package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leroytan/go-backend/controllers"
	"github.com/leroytan/go-backend/initializers"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.POST("/api/posts", controllers.PostsCreate)
	r.PUT("/api/posts/:id", controllers.PostsUpdate)
	r.GET("/api/posts", controllers.PostsIndex)
	r.GET("/api/posts/:id", controllers.PostsShow)
	r.DELETE("/api/posts/:id", controllers.PostsSoftDelete)
	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
