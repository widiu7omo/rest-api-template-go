package main

import (
	"boilerplate/database"
	"boilerplate/handlers"
	"boilerplate/repositories"
	"boilerplate/routers"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func InitializeRepositories() {
	repositories.Init()
	//add more here
}
func main() {
	// Parse command-line flags
	flag.Parse()
	// Connected with database
	database.Connect()
	InitializeRepositories()
	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	routers.ApiRouter(app)
	routers.WebRouter(app)
	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
