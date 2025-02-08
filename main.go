package main

import (
	"WMS/config"
	"WMS/routes"
	"fmt"
	"time"
    // "oms/kafkaconsumer"
	"github.com/omniful/go_commons/http"
	// "github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World!")

	server := http.InitializeServer(
		":8082",        // listen address
		10*time.Second, // read timeout
		10*time.Second, // write timeout
		70*time.Second, // idle timeout
	)
	config.InitializeDatabase()

	
    
	routes.HubRoutes(server)
	routes.SKURoutes(server)
	routes.InventoryRoutes(server)

	if err := server.StartServer("MyService"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}

	fmt.Println("Server started successfully")
	
}

