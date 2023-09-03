package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		e, err := entity.NewUserFactory("testeteste.com", "123456", "Gui", "Laranjeira")
		if err != nil {
			return c.JSON(err)
		}
		return c.JSON(e)
	})

	app.Listen(":3000")
}
