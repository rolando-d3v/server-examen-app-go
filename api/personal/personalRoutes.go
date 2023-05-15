package personal

import (
	"github.com/gofiber/fiber/v2"
)

func PersonalRoutes(app *fiber.App) {

	api := app.Group("/personal")
	api.Post("/", CreatePersonal)
	api.Get("/", GetAllPersonal)
	api.Get("/:id", GetIdPersonal)
	api.Put("/:id", UpdatePersonal)
	// api.Delete("/:id", DeleteIdPersonal)

}
