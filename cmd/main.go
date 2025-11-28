package cmd

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	r := fiber.New()
	r.Listen(":3000")
}
