package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leroytan/go-backend/initializers"
	"github.com/leroytan/go-backend/models"
)

func Createcourse(c *gin.Context) {
	//get the user creating the course
	userinfo, _ := c.Get("user")
	user := userinfo.(models.User)
	//get data from request body
	var body struct {
		Title       string
		Description string
	}
	c.BindJSON(&body)
	//Create the course
	course := models.Course{Title: body.Title, Description: body.Description}
	initializers.DB.Create(&course)
	//Add the first user to the course
	initializers.DB.Model(&course).Association("Users").Append(&user)
	//Response
	c.JSON(http.StatusOK, gin.H{})
}
func Getusercourses(c *gin.Context) {
	//Get user courses
	var user models.User
	userinfo, err1 := c.Get("user")
	user = userinfo.(models.User)
	if !err1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot get user",
		})
		return
	}

	err2 := initializers.DB.Model(&models.User{}).Preload("Courses").First(&user).Error
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot get courses",
		})
		return
	}
	//Response
	c.JSON(http.StatusOK, gin.H{
		"courses": user.Courses,
	})
}

func Addtocourse(c *gin.Context) {
	//get data from request body
	var body struct {
		UserIDs  []uint
		CourseID uint
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot recognise body",
		})
		return
	}
	//Find the users that we are adding
	var users []models.User
	initializers.DB.Where("id IN ?", body.UserIDs).Find(&users)
	//Find the course that we want to add to
	var course models.Course
	initializers.DB.First(&course, body.CourseID)
	//Add specificed user to specified course
	initializers.DB.Model(&course).Association("Users").Append(&users)
	//Response
	c.JSON(http.StatusOK, gin.H{})
}

func Removefromcourse(c *gin.Context) {
	//get data from request body
	var body struct {
		UserIDs  []uint
		CourseID uint
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot recognise body",
		})
		return
	}
	//Find the users that we are removing
	var users []models.User
	initializers.DB.Where("id IN ?", body.UserIDs).Find(&users)

	var course models.Course
	initializers.DB.First(&course, body.CourseID)
	//Delete user from course
	initializers.DB.Model(&course).Association("Users").Delete(&users)
	//Response
	c.JSON(http.StatusOK, gin.H{})
}
func Createcategory(c *gin.Context) {
	//get id from url
	courseidstr := c.Param("courseid")
	courseid, _ := strconv.Atoi(courseidstr)
	//get data from request body
	var body struct {
		Title       string
		Description string
	}
	c.BindJSON(&body)
	//Create the course
	category := models.Category{Title: body.Title, Description: body.Description, Subcategories: nil, CourseID: uint(courseid)}
	initializers.DB.Create(&category)

	//Response
	c.JSON(http.StatusOK, gin.H{})
}
func Getcategories(c *gin.Context) {
	//get id from url
	courseid := c.Param("courseid")
	var course models.Course

	userinfo, _ := c.Get("user")
	user := userinfo.(models.User)
	initializers.DB.Model(&models.Course{}).Preload("Categories.Subcategories").Preload("Users", &user).First(&course, courseid)

	//Response
	c.JSON(http.StatusOK, gin.H{
		"categories": course.Categories,
	})
}

func Createsubcategory(c *gin.Context) {
	//get id from url
	categoryid, _ := strconv.Atoi(c.Param("categoryid"))
	//get data from request body
	var body struct {
		Title       string
		Description string
	}
	c.BindJSON(&body)

	//Create the course
	subcategory := models.Subcategory{Title: body.Title, Description: body.Description, CategoryID: uint(categoryid)}
	initializers.DB.Create(&subcategory)

	//Response
	c.JSON(http.StatusOK, gin.H{})
}
