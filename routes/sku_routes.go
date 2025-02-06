package routes

import (
	"WMS/controllers"
    "github.com/omniful/go_commons/http"
	
)

func SKURoutes(r *http.Server) {
	skus := r.Group("/skus")
	{
		skus.POST("/", controllers.CreateSKU)
		skus.GET("/:id", controllers.GetSKUByID)
		skus.GET("/", controllers.GetSKUsByFilters)
		skus.GET("/all", controllers.GetAllSKUs)
	}
}