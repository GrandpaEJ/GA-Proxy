package src

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

// HandleGroq proxies requests to Groq API
func HandleGroq(c *fiber.Ctx) error {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		return c.Status(500).JSON(fiber.Map{"error": "GROQ_API_KEY not set"})
	}

	path := strings.TrimPrefix(c.Path(), "/groq")
	url := "https://api.groq.com/openai" + path
	c.Request().Header.Set("Authorization", "Bearer "+apiKey)
	
	return proxy.Do(c, url)
}
