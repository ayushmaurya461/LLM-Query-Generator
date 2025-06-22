package main

import (
	"github.com/ayushmaurya461/llm-query-generator.git/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/generate-query", handler.QueryHandler)

	app.Listen(":3000")
}
