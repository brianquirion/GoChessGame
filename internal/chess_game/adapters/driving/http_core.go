package driving

import "github.com/gofiber/fiber/v2"

func Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./assets/404.html")
}
