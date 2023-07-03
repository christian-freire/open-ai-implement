package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string `json:"text"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}

func main() {
	apiKey := "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"

	input := "Qual a origem da guitarra Gibson Les Paul?"

	requestBody, err := json.Marshal(map[string]interface{}{
		"model":      "text-davinci-003",
		"prompt":     input,
		"max_tokens": 4000,
	})
	if err != nil {
		fmt.Println("Erro ao construir corpo da requisição: ", err.Error())
		return
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Erro ao criar requisição: ", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao executar requisição: ", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta: ", err.Error())
		return
	}

	var apiResponse APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Erro ao analisar resposta JSON: ", err.Error())
		return
	}

	if len(apiResponse.Choices) > 0 {
		fmt.Println(apiResponse.Choices[0].Text)
	} else {
		fmt.Println("Nenhuma resposta encontrada.")
	}
}
