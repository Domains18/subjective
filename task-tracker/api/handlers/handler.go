package handlers

import (
	"net/http"

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

	response := repository.CreateAccount(user)

	if response.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": response.Message})
		return
	}

	if response.User == nil {
		c.JSON(http.StatusConflict, gin.H{"error": response.Message})
	}

	c.JSON(http.StatusCreated, gin.H{"message": response.Message, "user": response.User})
}
