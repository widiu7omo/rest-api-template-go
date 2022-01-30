package handlers

import (
	"boilerplate/models"
	repositories "boilerplate/repositories"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func JadwalList(c *fiber.Ctx) error {
	if c.Query("page") != "" {
		page := c.Query("page", "1")
		page_size := c.Query("page_size", "10")
		jadwals, err := repositories.JadwalGetPagination(page, page_size)
		return Response(c, jadwals, err)
	}
	jadwals, err := repositories.JadwalGet()
	if err != nil {
		return Response(c, nil, err)
	}
	return Response(c, jadwals, nil)
}
func JadwalGetById(c *fiber.Ctx) error {
	fmt.Println(c.Params("id"))
	jadwal, err := repositories.JadwalGetById(c.Params("id"))
	if err != nil {
		return Response(c, nil, err)
	}
	return Response(c, jadwal, nil)
}

func JadwalCreate(c *fiber.Ctx) error {
	if c.Is("json") {
		var jadwal models.Jadwal
		err := json.Unmarshal(c.Body(), &jadwal)
		if err != nil {
			return Response(c, nil, err)
		}
		jadwal, err = repositories.JadwalCreate(jadwal)
		if err != nil {
			return Response(c, nil, err)
		}
		return Response(c, jadwal, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid Content-Type"))
}
func JadwalUpdate(c *fiber.Ctx) error {
	if c.Is("json") && c.Params("id") != "" {
		jadwal, errGet := repositories.JadwalGetById(c.Params("id"))
		if errGet != nil {
			return Response(c, nil, errGet)
		}
		err := json.Unmarshal(c.Body(), &jadwal)
		if err != nil {
			return Response(c, nil, err)
		}
		jadwal, err = repositories.JadwalUpdate(jadwal)
		if err != nil {
			return Response(c, nil, err)
		}
		return Response(c, jadwal, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid Content-Type or ID Param not found"))
}
func JadwalDelete(c *fiber.Ctx) error {
	if c.Params("id") != "" {
		jadwal, errGet := repositories.JadwalGetById(c.Params("id"))
		if errGet != nil {
			return Response(c, nil, errGet)
		}
		deleteErr := repositories.JadwalDelete(jadwal)
		if deleteErr != nil {
			return Response(c, nil, deleteErr)
		}
		return Response(c, jadwal, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid ID Param not found"))

}
