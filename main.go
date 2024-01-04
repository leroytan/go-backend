package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leroytan/go-backend/controllers"
	"github.com/leroytan/go-backend/initializers"
	"github.com/leroytan/go-backend/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()

}

func main() {

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.POST("/api/posts", middleware.RequireAuth, controllers.PostsCreate)
	router.PUT("/api/posts/:id", middleware.RequireAuth, controllers.PostsUpdate)
	router.GET("/api/posts", controllers.PostsAll)
	router.GET("/api/posts/:id", controllers.PostIndex)
	router.DELETE("/api/posts/:id", middleware.RequireAuth, controllers.PostsSoftDelete)
	router.Run(":3000") // listen and serve on 0.0.0.0:3000
}
