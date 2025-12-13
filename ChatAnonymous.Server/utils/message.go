package utils

import "encoding/json"

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func EncodeMessage(msg *Message) ([]byte, error) {
	return json.Marshal(msg)
}
