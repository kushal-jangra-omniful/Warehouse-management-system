package controllers

import (
	"net/http"
	"WMS/models"
	"WMS/config"
	"github.com/gin-gonic/gin"
)

// Create or Update Inventory
func UpsertInventory(c *gin.Context) {
	var inventory models.Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingInventory models.Inventory
	err := config.DB.Where("hub_id = ? AND tenant_id = ? AND seller_id = ? AND sku_id = ?", inventory.HubID, inventory.TenantID, inventory.SellerID, inventory.SKU_ID).First(&existingInventory).Error

	if err != nil {
		config.DB.Create(&inventory)
		c.JSON(http.StatusCreated, inventory)
	} else {
		existingInventory.AvailableQuantity = inventory.AvailableQuantity
		config.DB.Save(&existingInventory)
		c.JSON(http.StatusOK, existingInventory)
	}
}

// Get Inventory by Seller ID and Hub ID
func GetInventoryBySellerAndHub(c *gin.Context) {
	sellerID := c.Query("sellerid")
	hubID := c.Query("hubid")
    
	var inventories []models.Inventory
	if err := config.DB.Where("sellerid = ? AND hubid = ?", sellerID, hubID).Find(&inventories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, inventories)
}


