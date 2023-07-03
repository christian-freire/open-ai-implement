package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func main() {

	viper.AutomaticEnv()

	viper.SetConfigFile("configs/env.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Erro ao ler arquivo de configuração:", err.Error())
		return
	}

	apiKey := viper.GetString("apiKey")

	input := "Vale a pena um notebook mac para programar?"

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

	fmt.Println(string(body))
}
