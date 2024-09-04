package Handler

import (
	"ESayur/Services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NewUserHandler(app fiber.Router, userServices Services.UserServices) {
	app.Get("/", GetUsers(userServices))
}

func GetUsers(userServices Services.UserServices) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := userServices.GetUsers(c.Context())
		fmt.Println(c.Context())
		if err != nil {
			if err.Error() == "data kosong" {
				return c.SendStatus(404)
			}
			return c.SendStatus(500)
		}

		return c.Status(200).JSON(users)
	}
}

func GetUser(userServices Services.UserServices) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	}
}
