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

func main() {

	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsSoftDelete)
	r.Run() // listen and serve on 0.0.0.0:3000
}
