package repository

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"backend/ai"
)

func GetTaskSuggestions(c *gin.Context) {
	var input struct {
		Task string `json:"task"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	suggestions, err := ai.GetTaskBreakdown(input.Task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate suggestions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"suggestions": suggestions})
}
