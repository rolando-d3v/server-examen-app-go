package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/goccy/go-json"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rolando-d3v/server-examen-app-go/db"
	
	"github.com/swaggo/fiber-swagger"
	_ "github.com/swaggo/fiber-swagger/example/docs" 

	// ROUTES
	"github.com/rolando-d3v/server-examen-app-go/api/documento"
	"github.com/rolando-d3v/server-examen-app-go/api/personal"
	"github.com/rolando-d3v/server-examen-app-go/api/role"
	"github.com/rolando-d3v/server-examen-app-go/api/user"
)

func main() {
	// app := fiber.New()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
        JSONDecoder: json.Unmarshal,
		ServerHeader: "Fiber",
		AppName:      "Server Examen App Ver 2.0",
	})

	godotenv.Load(".env")
	// godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		port = ":4040"
	} else {
		port = ":" + port
	}

	//middleware
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	app.Use(cors.New())
	app.Use(logger.New()) //tipo morgan

	//statticos
	app.Static("/public", "./public")

	//routes
	documento.SetupRoutes(app)
	personal.PersonalRoutes(app)
	user.UserRoutes(app)
	role.RoleRoutes(app)

	//base de datos
	db.ConexionBD()

	app.Listen(port)

}
