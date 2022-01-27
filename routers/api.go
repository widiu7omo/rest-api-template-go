package routers

import (
	"boilerplate/handlers"
	"github.com/gofiber/fiber/v2"
)

func ApiRouter(app *fiber.App) {
	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")
	//
	// Users handlers
	v1.Get("/users", handlers.UserList)
	v1.Get("/users/:id", handlers.UserGetById)
	v1.Post("/users", handlers.UserCreate)
	// Iurans handlers
	v1.Get("/iurans", handlers.IuranList)
	v1.Get("/iurans/:id", handlers.IuranGetById)
	v1.Post("/iurans", handlers.IuranCreate)
	// Warga handlers
	v1.Get("/wargas", handlers.WargaList)
	v1.Get("/wargas/:id", handlers.WargaGetById)
	v1.Post("/wargas", handlers.WargaCreate)
	// IuranWargas handlers
	v1.Get("/iuranWargas", handlers.IuranWargaList)
	v1.Get("/iuranWargas/:id", handlers.IuranWargaGetById)
	v1.Put("/iuranWargas/:id", handlers.IuranWargaUpdate)
	v1.Delete("/iuranWargas/:id", handlers.IuranWargaDelete)
	v1.Post("/iuranWargas", handlers.IuranWargaCreate)

}
