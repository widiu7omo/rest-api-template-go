package handlers

import (
	"boilerplate/models"
	repositores "boilerplate/repositories"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func WargaList(c *fiber.Ctx) error {
	wargas, err := repositores.WargaGet()
	return Response(c, wargas, err)
}
func WargaGetById(c *fiber.Ctx) error {
	fmt.Println(c.Params("id"))
	warga, err := repositores.WargaGetById(c.Params("id"))
	return Response(c, warga, err)
}

func WargaCreate(c *fiber.Ctx) error {
	if c.Is("json") {
		var warga models.Warga
		err := json.Unmarshal(c.Body(), &warga)
		if err != nil {
			return Response(c, nil, err)
		}
		repositores.WargaCreate(warga)
		return Response(c, warga, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid Content-Type"))
}
