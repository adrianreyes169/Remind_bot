package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func sendMessage(message string, config Config) error {

	var w bytes.Buffer

	err := json.NewEncoder(&w).Encode(map[string]any{"text": message, "chat_id": config.Chat_ID})
	if err != nil {
		return err
	}

	_, err = http.Post("https://api.telegram.org/bot"+config.Token+"/sendMessage", "application/json", &w)
	if err != nil {
		return err
	}

	return nil
}
