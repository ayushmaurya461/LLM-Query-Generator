package handler

import (
	"fmt"

	ollama_client "github.com/ayushmaurya461/llm-query-generator.git/ollama-client"
	"github.com/gofiber/fiber/v2"
)

type QueryRequest struct {
	Schema string `json:"schema"`
	Prompt string `json:"prompt"`
	Model  string `json:"model"`
}

func QueryHandler(c *fiber.Ctx) error {
	fmt.Println("Got a request")
	var req QueryRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	fullPrompt := fmt.Sprintf(`Schema:
		%s
		Prompt:
		%s
		Model:
		%s

		-----
		Only return the query, No need to explain the query
		`, req.Schema, req.Prompt, req.Model)

	response, err := ollama_client.GenerateQuery(fullPrompt, req.Model)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"query": response})
}
