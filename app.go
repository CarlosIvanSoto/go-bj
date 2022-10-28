package main

import (
	"encoding/json"
	"log"

	"github.com/CarlosIvanSoto/go-bj/database"
	"github.com/CarlosIvanSoto/go-bj/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// func tmpHandler(c *fiber.Ctx) error {
// 	return c.SendString("Hello, world!")
// }

func setRoutesPlayer(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/players", routes.AllPlayers)
	p := api.Group("/player")
	p.Get("/", routes.SearchPlayer)
	p.Post("/", routes.AddPlayer)
	p.Put("/", routes.UpdatePlayer)
	p.Delete("/", routes.DeletePlayer)
}

func main() {
	database.ConnectDB()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	setRoutesPlayer(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":5000"))
}
