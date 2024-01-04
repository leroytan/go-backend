package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leroytan/go-backend/initializers"
	"github.com/leroytan/go-backend/models"
)

// creates a new post after authorization from middleware
func PostsCreate(c *gin.Context) {
	//get data from req body and userid from jwt
	var user models.User
	var body struct {
		Title   string
		Content string
	}
	user = c.MustGet("user").(models.User)
	c.BindJSON(&body)
	//create post
	post := models.Post{Title: body.Title, Content: body.Content, UserID: user.ID}

	result := initializers.DB.Create(&post)
	//response
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// gets all the post
func PostsAll(c *gin.Context) {
	//Retrieve all the posts from database
	var posts []models.Post
	initializers.DB.Find(&posts)

	//respond with them
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

// gets a specific post
func PostIndex(c *gin.Context) {
	//get id from url
	postid := c.Param("id")
	//get the specific post from database
	var post models.Post
	result := initializers.DB.Find(&post, postid)

	//response
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// updates a post
func PostsUpdate(c *gin.Context) {
	//get id from url
	postid := c.Param("id")

	//get data from req body
	var user models.User
	var body struct {
		Title   string
		Content string
	}
	user = c.MustGet("user").(models.User)
	c.BindJSON(&body)

	//Retrieve the post we are updating from database
	var post models.Post
	initializers.DB.Find(&post, postid)

	//Check if user is post creator
	if post.UserID != user.ID {
		c.Status(http.StatusBadRequest)
		return
	}

	//update it
	result := initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Content: body.Content})

	//response
	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// softdeletes a specific post
func PostsSoftDelete(c *gin.Context) {
	//get id from url
	id := c.Param("id")
	//delete the posts
	initializers.DB.Delete(&models.Post{}, id)
	//response
	c.Status(http.StatusOK)
}
