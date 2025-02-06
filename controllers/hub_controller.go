package controllers

import (
	"WMS/config"
	"WMS/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Hub
func CreateHub(c *gin.Context) {
	var hub models.Hub
	if err := c.ShouldBindJSON(&hub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&hub)
	c.JSON(http.StatusCreated, hub)
}

// Get Hub by ID
func GetHubByID(c *gin.Context) {
	id := c.Param("id")
	var hub models.Hub

	if err := config.DB.First(&hub, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hub)
}
