package llm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Generate(prompt, apiKey string) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Can not load env")
	}

	body := map[string]interface{}{
		"model": os.Getenv("AI_MODEL"),
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	jsonData, _ := json.Marshal(body)

	req, err := http.NewRequest(
		"POST",
		os.Getenv("API_ENDPOINT"),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	// HANDLE API ERROR PROPERLY
	if errMsg, ok := result["error"]; ok {
		return "", fmt.Errorf("LLM error: %v", errMsg)
	}

	choices, ok := result["choices"]
	if !ok || choices == nil {
		return "", errors.New("no choices returned from LLM (check API key or model)")
	}

	choicesArr, ok := choices.([]interface{})
	if !ok || len(choicesArr) == 0 {
		return "", errors.New("invalid choices format from LLM")
	}

	first := choicesArr[0].(map[string]interface{})
	message := first["message"].(map[string]interface{})
	content := message["content"].(string)

	return content, nil
}