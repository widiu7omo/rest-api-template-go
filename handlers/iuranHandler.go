package handlers

import (
	"boilerplate/models"
	repositories "boilerplate/repositories"
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
	if err != nil {
		return Response(c, nil, err)
	}
	return Response(c, iurans, nil)
}
func IuranGetById(c *fiber.Ctx) error {
	fmt.Println(c.Params("id"))
	iuran, err := repositories.IuranGetById(c.Params("id"))
	if err != nil {
		return Response(c, nil, err)
	}
	return Response(c, iuran, nil)
}

func IuranCreate(c *fiber.Ctx) error {
	if c.Is("json") {
		var iuran models.Iuran
		err := json.Unmarshal(c.Body(), &iuran)
		if err != nil {
			return Response(c, nil, err)
		}
		iuran, err = repositories.IuranCreate(iuran)
		if err != nil {
			return Response(c, nil, err)
		}
		return Response(c, iuran, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid Content-Type"))
}
func IuranUpdate(c *fiber.Ctx) error {
	if c.Is("json") && c.Params("id") != "" {
		iuran, errGet := repositories.IuranGetById(c.Params("id"))
		if errGet != nil {
			return Response(c, nil, errGet)
		}
		err := json.Unmarshal(c.Body(), &iuran)
		if err != nil {
			return Response(c, nil, err)
		}
		iuran, err = repositories.IuranUpdate(iuran)
		if err != nil {
			return Response(c, nil, err)
		}
		return Response(c, iuran, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid Content-Type or ID Param not found"))
}
func IuranDelete(c *fiber.Ctx) error {
	if c.Params("id") != "" {
		iuran, errGet := repositories.IuranGetById(c.Params("id"))
		if errGet != nil {
			return Response(c, nil, errGet)
		}
		deleteErr := repositories.IuranDelete(iuran)
		if deleteErr != nil {
			return Response(c, nil, deleteErr)
		}
		return Response(c, iuran, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid ID Param not found"))

}
