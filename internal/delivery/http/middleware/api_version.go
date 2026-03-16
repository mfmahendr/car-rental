package middleware

import "github.com/gofiber/fiber/v2"

func APIVersion(version string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("X-API-Version", version)
		return c.Next()
	}
}