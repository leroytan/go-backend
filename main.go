package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leroytan/go-backend/controllers"
	"github.com/leroytan/go-backend/initializers"
	"github.com/leroytan/go-backend/middleware"
)

func init() {
	//initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()

}

func main() {

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	user := router.Group("/api")
	{
		//users options
		user.POST("/signup", controllers.Signup)
		user.POST("/login", controllers.Login)
		user.GET("/validate", middleware.RequireAuth, controllers.Validate)
		user.POST("/signout", middleware.RequireAuth, controllers.Signout)
		//post
		user.POST("/courses/:courseid/categories/:categoryid/subcategories/:subcategoryid/posts", middleware.RequireAuth, middleware.CourseAuth, controllers.PostsCreate)
		user.PUT("/courses/:courseid/categories/:categoryid/subcategories/:subcategoryid/posts/:postid", middleware.RequireAuth, middleware.CourseAuth, controllers.PostsUpdate)
		user.GET("/courses/:courseid/categories/:categoryid/subcategories/:subcategoryid/posts", middleware.RequireAuth, middleware.CourseAuth, controllers.PostsAll)
		user.GET("/courses/:courseid/categories/:categoryid/subcategories/:subcategoryid/posts/:postid", middleware.RequireAuth, middleware.CourseAuth, controllers.PostIndex)
		user.DELETE("/courses/:courseid/categories/:categoryid/subcategories/:subcategoryid/posts/:postid", middleware.RequireAuth, middleware.CourseAuth, controllers.PostsSoftDelete)

		user.GET("/users/:id", middleware.RequireAuth, controllers.Getuserdetails)
		user.GET("/getusercourses", middleware.RequireAuth, controllers.Getusercourses)
		user.GET("/courses/:courseid/categories/", middleware.RequireAuth, middleware.CourseAuth, controllers.Getcategories)
		//likes and dislikes
		user.POST("/courses/:courseid/categories/:categoryid/subcategories/:subcategoryid/posts/:postid/pollsoptions/:pollsoptionsid", middleware.RequireAuth, middleware.CourseAuth, controllers.Updatepoll)
		//user.GET("/courses/:courseid/categories/:categoryid/subcategories/:subcategoryid/posts/:postid/pollsoptions/:pollsoptionsid", middleware.RequireAuth, middleware.CourseAuth, controllers.Getpollsoptionsvotescount)
		//user.GET("/users/:id/pollvotes", middleware.RequireAuth, controllers.Getcurrentuservote)
	}
	admin := router.Group("/api/admin", middleware.RequireAuth, middleware.AdminAuth)
	{
		//admin options
		//users
		admin.GET("/users", controllers.Getallusers)
		//admin.GET("/courses/:courseid/users/user", controllers.Getcategories)
		//courses
		admin.POST("/courses/createcourse", controllers.Createcourse)
		admin.POST("/courses/:courseid/addtocourse", controllers.Addtocourse)
		admin.POST("/courses/:courseid/removefromcourse", controllers.Removefromcourse)
		admin.GET("/courses/:courseid/getusers", controllers.Getcourseusers)

		//category
		admin.POST("/courses/:courseid/categories/createcategory", controllers.Createcategory)
		admin.GET("/courses/:courseid/categories/", controllers.Getcategories)
		//subcategory
		admin.POST("/courses/:courseid/categories/:categoryid/subcategories/createsubcategory", controllers.Createsubcategory)

	}

	router.Run(":3000") // listen and serve on 0.0.0.0:3000
}
