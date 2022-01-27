package handlers

import (
	"boilerplate/models"
	repositores "boilerplate/repositories"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func IuranWargaList(c *fiber.Ctx) error {
	iuranWargas, err := repositores.IuranWargaGet()
	if err != nil {
		return Response(c, nil, err)
	}
	return Response(c, iuranWargas, nil)
}
func IuranWargaGetById(c *fiber.Ctx) error {
	fmt.Println(c.Params("id"))
	iuranWarga, err := repositores.IuranWargaGetById(c.Params("id"))
	if err != nil {
		return Response(c, nil, err)
	}
	return Response(c, iuranWarga, nil)
}

func IuranWargaCreate(c *fiber.Ctx) error {
	if c.Is("json") {
		var iuranWarga models.IuranWarga
		err := json.Unmarshal(c.Body(), &iuranWarga)
		if err != nil {
			return Response(c, nil, err)
		}
		createErr := repositores.IuranWargaCreate(iuranWarga)
		if createErr != nil {
			return Response(c, nil, createErr)
		}
		return Response(c, iuranWarga, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid Content-Type"))
}
func IuranWargaUpdate(c *fiber.Ctx) error {
	if c.Is("json") && c.Params("id") != "" {
		iuranWarga, errGet := repositores.IuranWargaGetById(c.Params("id"))
		if errGet != nil {
			return Response(c, nil, errGet)
		}
		err := json.Unmarshal(c.Body(), &iuranWarga)
		if err != nil {
			return Response(c, nil, err)
		}
		updateErr := repositores.IuranWargaUpdate(iuranWarga)
		if updateErr != nil {
			return Response(c, nil, updateErr)
		}
		return Response(c, iuranWarga, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid Content-Type or ID Param not found"))
}
func IuranWargaDelete(c *fiber.Ctx) error {
	if c.Params("id") != "" {
		iuranWarga, errGet := repositores.IuranWargaGetById(c.Params("id"))
		if errGet != nil {
			return Response(c, nil, errGet)
		}
		deleteErr := repositores.IuranWargaDelete(iuranWarga)
		if deleteErr != nil {
			return Response(c, nil, deleteErr)
		}
		return Response(c, "Successfully Deleted", nil)
	}
	return Response(c, nil, fmt.Errorf("invalid ID Param not found"))

}
