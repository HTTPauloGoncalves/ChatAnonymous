package utils

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func NewMessage(u, m string) *Message {
	return &Message{
		Username: u,
		Message:  m,
	}
}
