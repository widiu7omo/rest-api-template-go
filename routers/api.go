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
	v1.Put("/users/:id", handlers.UserUpdate)
	v1.Get("/users/:id", handlers.UserGetById)
	v1.Post("/users", handlers.UserCreate)
	v1.Delete("/users/:id", handlers.UserDelete)
	// Iurans handlers
	v1.Get("/iurans", handlers.IuranList)
	v1.Get("/iurans/:id", handlers.IuranGetById)
	v1.Post("/iurans", handlers.IuranCreate)
	v1.Put("/iurans/:id", handlers.IuranUpdate)
	v1.Delete("/iurans/:id", handlers.IuranDelete)
	// Warga handlers
	v1.Get("/wargas", handlers.WargaList)
	v1.Get("/wargasWithIuran", handlers.WargaListWithIuranWargas)
	v1.Get("/wargasWithIuran/:iuranId/:wargaId", handlers.WargaWithIuranWargaGetById)
	v1.Get("/wargas/:id", handlers.WargaGetById)
	v1.Post("/wargas", handlers.WargaCreate)
	v1.Put("/wargas/:id", handlers.WargaUpdate)
	v1.Delete("/wargas/:id", handlers.WargaDelete)
	// IuranWargas handlers
	v1.Get("/iuranWargas", handlers.IuranWargaList)
	v1.Get("/iuranWargas/:id", handlers.IuranWargaGetById)
	v1.Put("/iuranWargas/:id", handlers.IuranWargaUpdate)
	v1.Delete("/iuranWargas/:id", handlers.IuranWargaDelete)
	v1.Post("/iuranWargas", handlers.IuranWargaCreate)
}
