package role

import "github.com/gofiber/fiber/v2"

func RoleRoutes(app *fiber.App) {

	api := app.Group("/role")

	api.Get("/", GetAllRole)

}
