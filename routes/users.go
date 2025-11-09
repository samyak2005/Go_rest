package routes

import (
	"example/rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not create user",
			"error": err.Error()})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not save user",
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
		"user": user,
	})
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the input user data",
			"error": err.Error()})
		return
	}
	isValid := user.ValidateCredentials()
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": user,
	})
}