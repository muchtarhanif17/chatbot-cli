package ollama

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Request struct {
	Model   string                 `json:"model"`
	Prompt  string                 `json:"prompt"`
	Stream  bool                   `json:"stream"`
	Options map[string]interface{} `json:"options"`
}

type Response struct {
	Response string `json:"response"`
}

func Generate(prompt string) (string, error) {
	req := Request{
		Model:  "deepseek-r1",
		Prompt: prompt,
		Stream: false,
	}

	data, _ := json.Marshal(req)

	resp, err := http.Post(
		"http://localhost:11434/api/generate",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res Response
	json.NewDecoder(resp.Body).Decode(&res)

	return res.Response, nil
}
