package user



import (
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {

	api := app.Group("/user")
	api.Post("/", CreateUser)
	api.Get("/", GetAllUser)
	// api.Get("/:id", GetIdPersonal)
	api.Put("/:id", UpdateUser)
	api.Delete("/:id", DeleteUser)

}
