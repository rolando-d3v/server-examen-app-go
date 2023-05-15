package role

import "github.com/gofiber/fiber/v2"

func GetAllRole(c *fiber.Ctx) error {

	return c.Status(400).JSON(fiber.Map{
		"msj": "successfully",
	})

}
