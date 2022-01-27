package handlers

import (
	"boilerplate/models"
	repositores "boilerplate/repositories"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func UserList(c *fiber.Ctx) error {
	users, err := repositores.UserGet()
	return Response(c, users, err)
}
func UserGetById(c *fiber.Ctx) error {
	fmt.Println(c.Params("id"))
	user, err := repositores.UserGetById(c.Params("id"))
	return Response(c, user, err)
}

func UserCreate(c *fiber.Ctx) error {
	if c.Is("json") {
		var user models.User
		err := json.Unmarshal(c.Body(), &user)
		if err != nil {
			return Response(c, nil, err)
		}
		repositores.UserCreate(user)
		return Response(c, user, nil)
	}
	return Response(c, nil, fmt.Errorf("invalid Content-Type"))
}
