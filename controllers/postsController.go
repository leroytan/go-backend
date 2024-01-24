package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leroytan/go-backend/initializers"
	"github.com/leroytan/go-backend/models"
)

// creates a new post after authorization from middleware
func PostsCreate(c *gin.Context) {
	//get data from url
	subcategoryid, _ := strconv.Atoi(c.Param("subcategoryid"))
	courseid := c.Param("courseid")
	categoryid := c.Param("categoryid")

	userinfo, _ := c.Get("user")
	user := userinfo.(models.User)

	//get data from req body
	var body struct {
		Title        string
		Content      string
		ParentpostID uint
	}
	c.BindJSON(&body)

	//make sure that courseid, categoryid and subcategoryid are correct
	var courses []models.Course
	initializers.DB.Preload("Categories", "id = ?", categoryid).Preload("Categories.Subcategories", "id = ?", subcategoryid).First(&courses, courseid)
	if len(courses) == 0 || len(courses[0].Categories) == 0 || len(courses[0].Categories[0].Subcategories) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "courseid, categoryid or subcategoryid is wrong",
		})
		return
	}

	//create post
	var post models.Post
	if body.ParentpostID == 0 {
		post = models.Post{Title: body.Title, Content: body.Content, SubcategoryID: uint(subcategoryid), UserID: user.ID}
	} else {
		post = models.Post{Title: body.Title, Content: body.Content, ParentpostID: &body.ParentpostID, SubcategoryID: uint(subcategoryid), UserID: user.ID}
	}

	result := initializers.DB.Create(&post)
	//create polloptions
	polloptions := []*models.PollsOptions{
		{Title: "Upvote", PostID: post.ID, PollsOptionsVotes: []models.PollsOptionsVotes{}},
		{Title: "Downvote", PostID: post.ID, PollsOptionsVotes: []models.PollsOptionsVotes{}},
	}
	initializers.DB.Create(&polloptions)
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

// gets 15 post
func PostsAll(c *gin.Context) {
	//get id from url
	subcategoryid := c.Param("subcategoryid")
	courseid := c.Param("courseid")
	categoryid := c.Param("categoryid")

	//make sure that courseid, categoryid and subcategoryid are correct
	var courses []models.Course
	initializers.DB.Preload("Categories", "id = ?", categoryid).Preload("Categories.Subcategories", "id = ?", subcategoryid).First(&courses, courseid)
	if len(courses) == 0 || len(courses[0].Categories) == 0 || len(courses[0].Categories[0].Subcategories) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "courseid, categoryid or subcategoryid is wrong",
		})
		return
	}

	//Retrieve 15 posts from database
	var posts []models.Post
	initializers.DB.Limit(15).Where("parentpost_id IS NULL AND subcategory_id = ?", subcategoryid).Find(&posts)

	//respond with them
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

// gets a specific post
func PostIndex(c *gin.Context) {
	//get id from url
	postid := c.Param("postid")
	//get the specific post from database
	var post models.Post
	result := initializers.DB.Preload("Comments").Preload("PollsOptions").Find(&post, postid)

	//gets the current user's own votes

	user := c.MustGet("user").(models.User)
	initializers.DB.Preload("PollsOptionVotes").First(&user, user.ID)

	//Get post owner details
	var postowner models.User

	err2 := initializers.DB.First(&postowner, post.UserID).Error
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot get user",
			"user":  post.UserID,
		})
		return
	}
	//Get the comments
	var comments []models.Post
	initializers.DB.Where("parentpost_id = ?", postid).Find(&comments)
	//Find the users who interacted with the post
	var users []models.User
	userids := []uint{postowner.ID}
	for _, comment := range comments {
		userids = append(userids, comment.UserID)
	}
	initializers.DB.Select("id", "username", "role_id").Where("id IN ?", userids).Find(&users)

	//Response
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": gin.H{
			"ID":            post.ID,
			"CreatedAt":     post.CreatedAt,
			"UpdatedAt":     post.UpdatedAt,
			"DeletedAt":     post.DeletedAt,
			"Title":         post.Title,
			"Content":       post.Content,
			"ParentpostID":  post.ParentpostID,
			"Parentpost":    post.Parentpost,
			"UserID":        post.UserID,
			"Username":      postowner.Username,
			"SubcategoryID": post.SubcategoryID,
			"PollsOptions":  post.PollsOptions,
			"Upvotecount":   post.Upvotecount,
			"Downvotecount": post.Downvotecount,
			"Comments":      comments,
		},
		"userpollsvotes": user.PollsOptionVotes,
		"users":          users,
	})
}

// updates a post
func PostsUpdate(c *gin.Context) {
	//get id from url
	postid := c.Param("postid")

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
	id := c.Param("postid")
	//delete the posts
	initializers.DB.Delete(&models.Post{}, id)
	//response
	c.Status(http.StatusOK)
}
