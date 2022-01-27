package routers

import "github.com/gofiber/fiber/v2"

func WebRouter(app *fiber.App) {
	// Setup static files
	app.Static("/", "./static/public")
}
