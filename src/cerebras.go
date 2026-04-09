package src

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

// HandleCerebras proxies requests to Cerebras AI API
func HandleCerebras(c *fiber.Ctx) error {
	apiKey := os.Getenv("CEREBRAS_API_KEY")
	if apiKey == "" {
		return c.Status(500).JSON(fiber.Map{"error": "CEREBRAS_API_KEY not set"})
	}

	path := strings.TrimPrefix(c.Path(), "/cerebras")
	url := "https://api.cerebras.ai" + path
	c.Request().Header.Set("Authorization", "Bearer "+apiKey)
	
	return proxy.Do(c, url)
}
