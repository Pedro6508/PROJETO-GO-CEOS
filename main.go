package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	t "tree-ceos/projeto"
)

func main() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	app.Static("/", "./static")

	app.Get("/submit", t.Fazer_arvore)

	app.Get("/", func(ctx *fiber.Ctx) error {
		ctx.Redirect("/formulario.html")
		return nil
	})

	app.Listen(":3300")
}
