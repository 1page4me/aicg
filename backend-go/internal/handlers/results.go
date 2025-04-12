package handlers

import (
	"net/http"

	"aicg/internal/database"
	"aicg/internal/models"

	"github.com/gin-gonic/gin"
)

func GetResults(c *gin.Context) {
	var results []models.Result
	if err := database.DB.Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch results"})
		return
	}
	c.JSON(http.StatusOK, results)
}

func GetResult(c *gin.Context) {
	id := c.Param("id")
	var result models.Result
	if err := database.DB.First(&result, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Result not found"})
		return
	}
	c.JSON(http.StatusOK, result)
}
