package controllers

import (
	"errors"
	"net/http"
	"os"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/leroytan/go-backend/initializers"
	"github.com/leroytan/go-backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup(c *gin.Context) {
	//get email/password from req body
	var body struct {
		Email    string
		Username string
		Password string
	}
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	//validate username
	usernameerr := validateUsername(body.Username)
	if usernameerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": usernameerr.Error(),
		})
		return
	}
	//validate password
	passworderr := validatePassword(body.Password)
	if passworderr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": passworderr.Error(),
		})
		return
	}
	//Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
	}
	//Create the user
	user := models.User{Email: body.Email, Username: body.Username, Password: string(hash), Posts: []models.Post{}, RoleID: 1}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
	}
	//Respond
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	// Get the email and password from req body
	var body struct {
		Email    string
		Password string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	//Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	//Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	//Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	//Send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 60*60, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}
func Signout(c *gin.Context) {
	//Authorize user already done by middleware

	//Set cookie to expiry
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	//Response
	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	var user models.User
	data, _ := c.Get("user")
	user = data.(models.User)
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"ID":       user.ID,
			"Email":    user.Email,
			"Username": user.Username,
			"Posts":    user.Posts,
			"RoleID":   user.RoleID,
		},
	})
}

// Retrieve user list with eager loading posts
func GetAllPosts(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	err := db.Model(&models.User{}).Preload("Posts").Find(&users).Error
	return users, err
}

func validateUsername(username string) error {
	//check username for only alphaNumeric character
	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return errors.New("Username can only consists of alphabets and numbers")
		}

	}
	//check username length
	if len(username) >= 4 && len(username) <= 15 {
		return nil
	} else {
		return errors.New("Username length must be 4-15")
	}
}
func validatePassword(password string) error {
	//passwords must contain alphabets, numbers, and punctuation
	var containsNumber = false
	var containsalphabets = false
	var containspunctuation = false
	for _, char := range password {
		if unicode.IsLetter(char) {
			containsalphabets = true
			continue
		}
		if unicode.IsNumber(char) {
			containsNumber = true
			continue
		}
		if unicode.IsPunct(char) {
			containspunctuation = true
			continue
		}
	}
	if !containsNumber || !containsalphabets || !containspunctuation {
		return errors.New("password must contain alphabets, numbers, and punctuation")
	}
	//check username length
	if len(password) >= 7 {
		return nil
	} else {
		return errors.New("password length must be at least 7")
	}
}

func Getuserdetails(c *gin.Context) {
	//Get id from url
	userid := c.Param("id")
	//Get user courses
	var user models.User

	err2 := initializers.DB.Model(&models.User{}).First(&user, userid).Error
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot get user",
		})
		return
	}
	//Response
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"ID":       user.ID,
			"Username": user.Username,
			"Posts":    user.Posts,
			"RoleID":   user.RoleID,
			//Doesnt return email because it is part of privacy
		}})
}
func Getallusers(c *gin.Context) {

	var users []models.User
	initializers.DB.Find(&users)
	usersarray := []models.User{}

	for _, user := range users {
		usersarray = append(usersarray, models.User{
			Model:    user.Model,
			Username: user.Username,
			RoleID:   user.RoleID,
			Courses:  user.Courses,
		})

	}
	//Response
	c.JSON(http.StatusOK, gin.H{
		"users": usersarray,
	})
}
func Getcourseusers(c *gin.Context) {
	//Get id from url
	courseid := c.Param("courseid")
	usersarray := []models.User{}
	var course models.Course
	initializers.DB.Preload("Users").Find(&course, courseid)
	for _, user := range course.Users {
		usersarray = append(usersarray, models.User{
			Model:    user.Model,
			Username: user.Username,
			RoleID:   user.RoleID,
			Courses:  user.Courses,
		})

	}
	//Response
	c.JSON(http.StatusOK, gin.H{
		"users": usersarray,
	})
}
