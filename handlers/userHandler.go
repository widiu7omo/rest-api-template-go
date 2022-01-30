package handlers

import (
	"boilerplate/models"
	repositories "boilerplate/repositories"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func UserList(c *fiber.Ctx) error {
	if c.Query("page") != "" {
		page := c.Query("page", "1")
		pageSize := c.Query("page_size", "10")
		users, err := repositories.UserGetPagination(page, pageSize)
		return Response(c, users, err)
	}
	users, err := repositories.UserGet()
	if err != nil {
		return Response(c, nil, err)
	}
	return Response(c, users, nil)
}
func UserGetById(c *fiber.Ctx) error {
	fmt.Println(c.Params("id"))
	user, err := repositories.UserGetById(c.Params("id"))
	if err != nil {
		return Response(c, nil, err)
	}
	return Response(c, user, nil)
}

func UserCreate(c *fiber.Ctx) error {
	if c.Is("json") {
		var user models.User
		err := json.Unmarshal(c.Body(), &user)
		if err != nil {
			return Response(c, nil, err)
		}
		user, err = repositories.UserCreate(user)
		if err != nil {
			return Response(c, nil, err)
		}
		return Response(c, user, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid Content-Type"))
}
func UserUpdate(c *fiber.Ctx) error {
	if c.Is("json") && c.Params("id") != "" {
		user, errGet := repositories.UserGetById(c.Params("id"))
		if errGet != nil {
			return Response(c, nil, errGet)
		}
		err := json.Unmarshal(c.Body(), &user)
		if err != nil {
			return Response(c, nil, err)
		}
		user, err = repositories.UserUpdate(user)
		if err != nil {
			return Response(c, nil, err)
		}
		return Response(c, user, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid Content-Type or ID Param not found"))
}
func UserDelete(c *fiber.Ctx) error {
	if c.Params("id") != "" {
		user, errGet := repositories.UserGetById(c.Params("id"))
		if errGet != nil {
			return Response(c, nil, errGet)
		}
		deleteErr := repositories.UserDelete(user)
		if deleteErr != nil {
			return Response(c, nil, deleteErr)
		}
		return Response(c, user, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid ID Param not found"))

}
