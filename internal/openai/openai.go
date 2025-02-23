package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"svnfluence/internal/models"
)

const openAIEndpoint = "https://api.openai.com/v1/chat/completions"

type OpenAIClient struct {
	APIKey string
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{APIKey: apiKey}
}

func (c *OpenAIClient) GetCommand(query string) (*models.Command, error) {
	prompt := fmt.Sprintf(
		"You are an expert in Apache Subversion (SVN). Given the user query '%s', provide an SVN command that best matches it. Return your response as a JSON object with these exact fields: 'command' (the SVN command), 'description' (a short explanation), 'usage' (syntax), and 'examples' (a list of 1-2 example uses). If no command matches, return an empty JSON object {}. Do not include any text outside the JSON.",
		query,
	)

	requestBody, _ := json.Marshal(map[string]interface{}{
		"model": "gpt-4-turbo", // or "gpt-3.5-turbo"
		"messages": []map[string]string{
			{"role": "system", "content": "You are an SVN expert."},
			{"role": "user", "content": prompt},
		},
		"temperature": 0.5,
	})

	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	content := result["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	var command models.Command
	if err := json.Unmarshal([]byte(content), &command); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	if command.Command == "" {
		return nil, nil
	}
	return &command, nil
}
