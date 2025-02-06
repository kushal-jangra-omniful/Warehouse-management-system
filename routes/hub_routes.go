package routes

import (
	"WMS/controllers"
    "github.com/omniful/go_commons/http"
	
)

func HubRoutes(r *http.Server) {
	hubs := r.Group("/hubs")
	{
		hubs.POST("/", controllers.CreateHub)
		hubs.GET("/:id", controllers.GetHubByID)
	}
}
