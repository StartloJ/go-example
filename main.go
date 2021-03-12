package main

import (
	"go-example/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	user.UserRoute(app)
	app.Listen(":3000")
}
