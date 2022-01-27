package handlers

import "github.com/gofiber/fiber/v2"

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}

func Response(c *fiber.Ctx, data interface{}, err error) error {
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"data":    nil,
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"error":   nil,
		"data":    &data,
	})
}
