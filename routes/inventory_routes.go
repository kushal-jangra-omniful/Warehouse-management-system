package routes

import (
	"WMS/controllers"
    "github.com/omniful/go_commons/http"
	
)

func InventoryRoutes(r *http.Server) {
	inventories := r.Group("/inventories")
	{
		inventories.POST("/", controllers.UpsertInventory)
		inventories.GET("/", controllers.GetInventoryBySellerAndHub)
	}
}