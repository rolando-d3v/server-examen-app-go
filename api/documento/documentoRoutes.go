package documento

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/personales", logger.New())
	api.Post("/", CreatePersonal)
	api.Get("/", GetAllPersonal)
	api.Get("/:id", GetIdPersonal)
	api.Delete("/:id", DeleteIdPersonal)

}
