package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leroytan/go-backend/initializers"
	"github.com/leroytan/go-backend/models"
	"gorm.io/gorm"
)

func Updatepoll(c *gin.Context) {
	//get id from param
	postid := c.Param("postid")
	pollsoptionsid := c.Param("pollsoptionsid")

	user := c.MustGet("user").(models.User)
	var post models.Post
	var pollsoptions []models.PollsOptions
	var pollsoptionsvotes models.PollsOptionsVotes

	//make sure pollsoption belong to post
	initializers.DB.Where("id = ? AND post_id = ?", pollsoptionsid, postid).First(&pollsoptions)
	if len(pollsoptions) != 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post and pollsoptions id do not match",
		})
		return
	}
	//make sure user is not the owner of the post
	initializers.DB.Find(&post, postid)
	if post.UserID == user.ID {
		c.Status(http.StatusBadRequest)
		return
	}

	//get poll options from post
	initializers.DB.Select("id").Where("post_id = ?", postid).Find(&pollsoptions)
	//create an array of ids
	ids := []uint{}
	for _, item := range pollsoptions {
		ids = append(ids, item.ID)
	}
	//if user has polled with the same option before, remove the previous vote and return
	result := initializers.DB.Where("user_id = ? AND polls_options_id = ?", user.ID, pollsoptionsid).Unscoped().Delete(&pollsoptionsvotes)
	//Update the upvote and downvote count
	initializers.DB.First(&pollsoptions, pollsoptionsid)

	if result.RowsAffected == 1 {
		if pollsoptions[0].Title == "Upvote" {
			initializers.DB.Model(&post).UpdateColumn("upvotecount", gorm.Expr("upvotecount  + ?", -1))
		} else if pollsoptions[0].Title == "Downvote" {
			initializers.DB.Model(&post).UpdateColumn("downvotecount", gorm.Expr("downvotecount  + ?", -1))
		}
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	//if user has polled with a different option before, remove the previous vote
	result2 := initializers.DB.Where("user_id = ? AND polls_options_id in ?", user.ID, ids).Unscoped().Delete(&pollsoptionsvotes)

	//if the user has not polled before, create polloptionvote with the polloption
	pollsoptionsidint, _ := strconv.Atoi(pollsoptionsid)
	polloptionvotes := models.PollsOptionsVotes{PollsOptionsID: uint(pollsoptionsidint), UserID: user.ID}
	initializers.DB.Create(&polloptionvotes)

	//Update the upvote and downvote count
	if result2.RowsAffected > 0 {
		if pollsoptions[0].Title == "Upvote" {
			initializers.DB.Model(&post).UpdateColumn("downvotecount", gorm.Expr("downvotecount  + ?", -1))
			initializers.DB.Model(&post).UpdateColumn("upvotecount", gorm.Expr("upvotecount  + ?", 1))
		} else if pollsoptions[0].Title == "Downvote" {
			initializers.DB.Model(&post).UpdateColumn("downvotecount", gorm.Expr("downvotecount  + ?", 1))
			initializers.DB.Model(&post).UpdateColumn("upvotecount", gorm.Expr("upvotecount  + ?", -1))
		}
	} else {
		if pollsoptions[0].Title == "Upvote" {
			initializers.DB.Model(&post).UpdateColumn("upvotecount", gorm.Expr("upvotecount  + ?", 1))
		} else if pollsoptions[0].Title == "Downvote" {
			initializers.DB.Model(&post).UpdateColumn("downvotecount", gorm.Expr("downvotecount  + ?", 1))
		}
	}

	//Response
	c.JSON(http.StatusOK, gin.H{})
}

/*these functions were used in development

func Getcurrentuservote(c *gin.Context) {
	//get user
	var user models.User
	user = c.MustGet("user").(models.User)
	initializers.DB.Preload("PollsOptionVotes").First(&user, user.ID)

	//Response
	c.JSON(http.StatusOK, gin.H{
		"userpollvotes": user.PollsOptionVotes,
	})

}



func Getpollsoptionsvotescount(c *gin.Context) {
	//get id from param
	pollsoptionsid := c.Param("pollsoptionsid")
	var count int64

	//Count number of pollsoptionsvotes of the pollsoptionsid
	var pollsoptionsvotes models.PollsOptionsVotes
	initializers.DB.Where("polls_options_id = ?", pollsoptionsid).Find(&pollsoptionsvotes).Count(&count)

	//Response
	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}
*/
