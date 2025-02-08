package controllers

import (
	"WMS/config"
	"WMS/models"

	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create SKU
func CreateSKU(c *gin.Context) {
	var sku models.SKU
	if err := c.ShouldBindJSON(&sku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&sku)
	c.JSON(http.StatusCreated, sku)
}

// Get SKU by ID
func GetSKUByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"is_success": false,
			"status_code": http.StatusBadRequest,
			"data": gin.H{
				"message": "ID parameter is required",
			},
			"meta": gin.H{},
		})
		return
	}

	var sku models.SKU

	if err := config.DB.First(&sku, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"is_success": false,
			"status_code": http.StatusNotFound,
			"data": gin.H{
				"message": err.Error(),
			},
			"meta": gin.H{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"is_success": true,
		"status_code": http.StatusOK,
		"data": sku,
		"meta": gin.H{},
	})
}


func GetAllSKUs(c *gin.Context) {
	var skus []models.SKU
	if err := config.DB.Find(&skus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, skus)
}

// Get SKUs by Tenant ID, Seller ID, and SKU Codes
func GetSKUsByFilters(c *gin.Context) {
	tenantID := c.Query("tenant_id")
	sellerID := c.Query("seller_id")
	skuCodes := c.QueryArray("sku_codes")

	var skus []models.SKU
	query := config.DB.Where("tenant_id = ? AND seller_id = ?", tenantID, sellerID)
	if len(skuCodes) > 0 {
		query = query.Where("id IN ?", skuCodes)
	}
	if err := query.Find(&skus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, skus)
}
