package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leroytan/go-backend/initializers"
	"github.com/leroytan/go-backend/models"
)

func PostsCreate(c *gin.Context) {
	//get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	//create post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	//response
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//get id from url
	id := c.Param("id")
	//get the posts
	var post models.Post
	result := initializers.DB.Find(&post, id)

	//response
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//get id from url
	id := c.Param("id")
	//get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	//find the post we are updating
	var post models.Post
	initializers.DB.Find(&post, id)

	//update it
	result := initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	//response
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsSoftDelete(c *gin.Context) {
	//get id from url
	id := c.Param("id")
	//delete the posts
	initializers.DB.Delete(&models.Post{}, id)
	//response
	c.Status(200)
}
