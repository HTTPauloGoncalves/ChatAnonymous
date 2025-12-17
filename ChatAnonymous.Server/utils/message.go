package utils

import "encoding/json"

type Message struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func EncodeMessage(msg *Message) ([]byte, error) {
	return json.Marshal(msg)
}
