package handlers

import (
	"boilerplate/models"
	"boilerplate/repositories"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func IuranList(c *fiber.Ctx) error {
	fmt.Println(c.Query("with"))
	if c.Query("with") == "total" {
		iurans, err := repositories.IuranWithTotal()
		return Response(c, iurans, err)
	}
	iurans, err := repositories.IuranGet()
	return Response(c, iurans, err)
}
func IuranGetById(c *fiber.Ctx) error {
	fmt.Println(c.Params("id"))
	iuran, err := repositories.IuranGetById(c.Params("id"))
	return Response(c, iuran, err)
}

func IuranCreate(c *fiber.Ctx) error {
	if c.Is("json") {
		var iuran models.Iuran
		err := json.Unmarshal(c.Body(), &iuran)
		if err != nil {
			return Response(c, nil, err)
		}
		repositories.IuranCreate(iuran)
		return Response(c, iuran, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid Content-Type"))
}
