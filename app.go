package main

import (
	"boilerplate/database"
	"boilerplate/handlers"
	"boilerplate/repositories"
	"boilerplate/routers"
	"flag"
	"fmt"
	"log"
	"os"

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
	var env = os.Getenv("APP_ENV")
	fmt.Println(env)
	var dbUrl string
	if env == "production" {
		dbUrl = "u6945793_erte5:JbAADXqobwWQ@tcp(127.0.0.1:3306)/u6945793_erte5?charset=utf8mb4"
	} else {
		dbUrl = "root:@tcp(127.0.0.1:3306)/posrt5?charset=utf8mb4"
	}
	// Parse command-line flags
	flag.Parse()
	// Connected with database
	database.Connect(dbUrl)
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
