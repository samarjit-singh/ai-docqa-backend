package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables: ", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" 
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Hello from Go Fiber!")
	})

	log.Println("🌍 Server is running on port", port)
	log.Fatal(app.Listen(":" + port))
}
