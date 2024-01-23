package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/leroytan/go-backend/initializers"
	"github.com/leroytan/go-backend/models"
)

// Middleware for any operations that requires logging in
func RequireAuth(c *gin.Context) {
	//Get the cookie off req
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//Decode tokenString
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//Check expiration of token
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//Find the user with token userid
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//Attach to req
		c.Set("user", user)
		//Continue
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}

// Middleware for any operations that require user to be in course
func CourseAuth(c *gin.Context) {
	//get data from url
	courseid, _ := strconv.Atoi(c.Param("courseid"))
	//check that user is authorized to access course
	var course models.Course

	userinfo, _ := c.Get("user")
	user := userinfo.(models.User)
	initializers.DB.Model(&models.Course{}).Preload("Users", &user).First(&course, courseid)
	//user was not authorized to access
	if len(course.Users) != 1 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//Continue
	c.Next()
}

// Middleware for any operations that require user to be an admin
func AdminAuth(c *gin.Context) {
	//check that the user is admin
	userinfo, _ := c.Get("user")
	user := userinfo.(models.User)
	if user.RoleID < 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()
}
