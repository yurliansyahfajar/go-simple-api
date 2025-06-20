package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yurliansyahfajar/go-simple-api/models"
)

func signup(c *gin.Context) {
	var user models.User

	// read from request.body
	err := c.ShouldBindBodyWithJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(500, gin.H{"message": "Failed create user, please try again."})
		return
	}

	c.JSON(201, gin.H{
		"message": "user created!",
		"event":   user,
	})

}

func login(c *gin.Context) {
	var user models.User

	//read from request.body
	err := c.ShouldBindBodyWithJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Login Successful"})
}
