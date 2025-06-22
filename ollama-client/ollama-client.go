package ollama_client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type OllamaRequest struct {
	Prompt string `json:"prompt"`
	Model  string `json:"model"`
}

type OllamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func GenerateQuery(prompt string, model string) (string, error) {
	payload := OllamaRequest{
		Model:  model,
		Prompt: prompt,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "http://localhost:11434/api/generate", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Timeout: 60 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Stream the response
	var result string
	decoder := json.NewDecoder(res.Body)

	for {
		var chunk OllamaResponse
		if err := decoder.Decode(&chunk); err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
		result += chunk.Response
		if chunk.Done {
			break
		}
	}

	return result, nil
}
