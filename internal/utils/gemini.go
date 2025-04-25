package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func AskGemini(document, question string) (string, error) {
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=" + os.Getenv("GEMINI_API_KEY")

	reqBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{"parts": []map[string]string{
				{"text": fmt.Sprintf("Document:\n%s\n\nQuestion:\n%s", document, question)},
			}},
		},
	}

	body, _ := json.Marshal(reqBody)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBody, _ := ioutil.ReadAll(res.Body)

	var parsed map[string]interface{}
	json.Unmarshal(resBody, &parsed)

	text := parsed["candidates"].([]interface{})[0].(map[string]interface{})["content"].(map[string]interface{})["parts"].([]interface{})[0].(map[string]interface{})["text"].(string)

	return text, nil
}
