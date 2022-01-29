package handlers

import (
	"boilerplate/models"
	"boilerplate/repositories"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func WargaList(c *fiber.Ctx) error {
	if c.Query("page") != "" {
		page := c.Query("page", "1")
		page_size := c.Query("page_size", "10")
		wargas, err := repositories.WargaGetPagination(page, page_size)
		return Response(c, wargas, err)
	}
	wargas, err := repositories.WargaGet()
	return Response(c, wargas, err)
}
func WargaGetById(c *fiber.Ctx) error {
	fmt.Println(c.Params("id"))
	warga, err := repositories.WargaGetById(c.Params("id"))
	return Response(c, warga, err)
}

func WargaCreate(c *fiber.Ctx) error {
	if c.Is("json") {
		var warga models.Warga
		err := json.Unmarshal(c.Body(), &warga)
		if err != nil {
			return Response(c, nil, err)
		}
		repositories.WargaCreate(warga)
		return Response(c, warga, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid Content-Type"))
}
