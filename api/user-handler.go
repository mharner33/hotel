package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mharner33/hotel/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Mike",
		LastName:  "Harner",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "user api endpoint"})
}
