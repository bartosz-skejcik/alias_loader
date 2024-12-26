package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Messages []Message `json:"messages"`
	Model    string    `json:"model"`
}

type Response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func QueryAI(systemPrompt, userPrompt string) (string, error) {
	reqBody := Request{
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
		Model: "llama3-8b-8192",
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("GROQ_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("error parsing response: %v", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response content received")
	}

	return response.Choices[0].Message.Content, nil
}

func GetTableTitle(category string) (string, error) {
	return QueryAI(
		"Generate a short, descriptive title for a category of shell aliases. Respond with just the title.",
		category,
	)
}

func GetAliasDescription(alias string) (string, error) {
	return QueryAI(
		"Your task is to respond with a short and brief description of a given alias. Dont use the word 'alias', only respond with what this command does",
		alias,
	)
}
