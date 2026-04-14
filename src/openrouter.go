package src

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

// HandleOpenRouter proxies requests to OpenRouter API
func HandleOpenRouter(c *fiber.Ctx) error {
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		return c.Status(500).JSON(fiber.Map{"error": "OPENROUTER_API_KEY not set"})
	}

	path := strings.TrimPrefix(c.Path(), "/openrouter")
	url := "https://openrouter.ai/api" + path
	c.Request().Header.Set("Authorization", "Bearer "+apiKey)
	c.Request().Header.Set("HTTP-Referer", "https://zero-bot.net")
	c.Request().Header.Set("X-Title", "GA-Proxy")

	return proxy.Do(c, url)
}
