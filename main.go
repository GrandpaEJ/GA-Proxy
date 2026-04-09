package main

import (
	"fmt"
	"log"
	"os"

	"proxy-server/src"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// Version of the GA-Proxy
const Version = "1.0.1"

// Obfuscate is implemented in obfs_amd64.s
// It performs a hardware-level XOR between data and key using Plan 9 Assembly
func Obfuscate(data uint64, key uint64) uint64

func main() {
	_ = godotenv.Load()

	app := fiber.New(fiber.Config{
		AppName: "GA-Proxy-v" + Version,
	})

	app.Use(logger.New())
	app.Use(cors.New())

	// Security Middleware: Verify X-GA-Secret if set
	app.Use(func(c *fiber.Ctx) error {
		secret := os.Getenv("PROXY_SECRET")
		if secret == "" {
			return c.Next()
		}

		// Skip secret check for handshake and health endpoints
		if c.Path() == "/" || c.Path() == "/handshake" || c.Path() == "/version" {
			return c.Next()
		}

		clientSecret := c.Get("X-GA-Secret")
		if clientSecret != secret {
			return c.Status(403).JSON(fiber.Map{"error": "Unauthorized: Invalid GA Secret"})
		}

		return c.Next()
	})

	// --- Handshake & Versioning ---

	app.Get("/version", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"version": Version})
	})

	// Handshake: Panel sends a challenge, we return the assembly-obfuscated result
	app.Get("/handshake", func(c *fiber.Ctx) error {
		challenge := c.QueryInt("challenge", 0)
		secret := os.Getenv("PROXY_SECRET")
		if secret == "" {
			secret = "GA-DEFAULT-SECRET"
		}

		var key uint64
		for _, char := range secret {
			key += uint64(char)
		}

		result := Obfuscate(uint64(challenge), key)

		return c.JSON(fiber.Map{
			"status": "ready",
			"version": Version,
			"signature": fmt.Sprintf("%x", result),
		})
	})

	// --- Proxy Endpoints ---

	// Handlers are moved to src/ package for scalability
	app.All("/groq/*", src.HandleGroq)
	app.All("/openrouter/*", src.HandleOpenRouter)
	app.All("/cerebras/*", src.HandleCerebras)

	// Default Health Check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "online",
			"name": "GA-Proxy",
			"version": Version,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("GA-Proxy %s starting on port %s", Version, port)
	log.Fatal(app.Listen(":" + port))
}
