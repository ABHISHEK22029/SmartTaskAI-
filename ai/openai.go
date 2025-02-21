package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type OpenAIRequest struct {
	Prompt string `json:"prompt"`
}

type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func GetTaskBreakdown(task string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	url := "https://api.openai.com/v1/completions"

	requestBody, _ := json.Marshal(OpenAIRequest{Prompt: "Break down this task: " + task})
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var openAIResponse OpenAIResponse
	json.NewDecoder(resp.Body).Decode(&openAIResponse)

	if len(openAIResponse.Choices) > 0 {
		return openAIResponse.Choices[0].Text, nil
	}

	return "", fmt.Errorf("No response from OpenAI")
}
