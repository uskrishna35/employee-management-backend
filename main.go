// @title Fiber App API
// @version 1.0
// @description This is a sample server using Fiber and Swagger.
// @host localhost:3000
// @BasePath /
// @schemes http

package main

import (
	"fmt"
	"log"

	"fiber-app/database"
	"fiber-app/handlers"
	"fiber-app/routes"

	_ "fiber-app/docs"

	"github.com/gofiber/swagger"      
	 

	"github.com/gofiber/fiber/v2"
	//  "github.com/swaggo/files"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	if err := database.ConnectDB(); err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}


	handlers.InitCollections()


	defer database.DisconnectDB()

	// Initialize Fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:4200", // Change this if your Angular app is hosted elsewhere
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, orgid",
		AllowCredentials: true,
	}))


	routes.SetupRoutes(app)

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "http://localhost:3000/swagger/doc.json", // Change if needed
		DeepLinking: false,
	}))
	


	// Start the server
	port := ":3000"
	fmt.Println("Server is running on port", port)
	log.Fatal(app.Listen(port))
}
