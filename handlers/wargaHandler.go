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
	if err != nil {
		return Response(c, nil, err)
	}
	return Response(c, wargas, nil)
}
func WargaListWithIuranWargas(c *fiber.Ctx) error {
	if c.Query("page") != "" {
		page := c.Query("page", "1")
		page_size := c.Query("page_size", "10")
		iuranId := c.Query("iuran_id", "")
		wargas, err := repositories.WargaWithIuranWargaGetPagination(page, page_size, iuranId)
		return Response(c, wargas, err)
	}
	wargas, err := repositories.WargaGet()
	if err != nil {
		return Response(c, nil, err)
	}
	return Response(c, wargas, nil)
}
func WargaGetById(c *fiber.Ctx) error {
	fmt.Println(c.Params("id"))
	warga, err := repositories.WargaGetById(c.Params("id"))
	if err != nil {
		return Response(c, nil, err)
	}
	return Response(c, warga, nil)
}

func WargaCreate(c *fiber.Ctx) error {
	if c.Is("json") {
		var warga models.Warga
		err := json.Unmarshal(c.Body(), &warga)
		if err != nil {
			return Response(c, nil, err)
		}
		warga, err = repositories.WargaCreate(warga)
		if err != nil {
			return Response(c, nil, err)
		}
		return Response(c, warga, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid Content-Type"))
}
func WargaUpdate(c *fiber.Ctx) error {
	if c.Is("json") && c.Params("id") != "" {
		warga, errGet := repositories.WargaGetById(c.Params("id"))
		if errGet != nil {
			return Response(c, nil, errGet)
		}
		err := json.Unmarshal(c.Body(), &warga)
		if err != nil {
			return Response(c, nil, err)
		}
		warga, err = repositories.WargaUpdate(warga)
		if err != nil {
			return Response(c, nil, err)
		}
		return Response(c, warga, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid Content-Type or ID Param not found"))
}
func WargaDelete(c *fiber.Ctx) error {
	if c.Params("id") != "" {
		warga, errGet := repositories.WargaGetById(c.Params("id"))
		if errGet != nil {
			return Response(c, nil, errGet)
		}
		deleteErr := repositories.WargaDelete(warga)
		if deleteErr != nil {
			return Response(c, nil, deleteErr)
		}
		return Response(c, warga, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid ID Param not found"))

}
