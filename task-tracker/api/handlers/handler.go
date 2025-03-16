package handlers

import (
	"net/http"

	"github.com/Domains18/subjective.git/api/auth"
	"github.com/Domains18/subjective.git/internal/model"
	"github.com/Domains18/subjective.git/pkg/repository"
	"github.com/gin-gonic/gin"
)

func CreateAccountHandler(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	response, err := repository.CreateAccount(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": response.Message})
		return
	}

	if response.User == nil {
		c.JSON(http.StatusConflict, gin.H{"error": response.Message})
	}

	c.JSON(http.StatusCreated, gin.H{"message": response.Message, "user": response.User})
}



func LoginHandler(c *gin.Context){
	var dto model.Login_types

	if err := c.ShouldBindBodyWithJSON(&dto);  err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	response, err := auth.Login(dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": response.Message})
		return
	}

	if response.User == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": response.Message})
	}

	c.JSON(http.StatusOK, gin.H{"token": response.Token, "user": response.User})
}


func GetUsers(c *gin.Context){
	response, err := repository.GetUsers()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": response})
}