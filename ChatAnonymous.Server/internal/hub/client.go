package hub

import (
	"github.com/HTTPauloGoncalves/ChatAnonymous/ChatAnonymous.Server/utils"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
	Room *Room
	Hub  *Hub
}

func (c *Client) ReadPump() {
	defer func() {
		if c.Room != nil {
			c.Room.Leave <- c
		}
		c.Conn.Close()
	}()

	for {
		var incoming utils.Message

		err := c.Conn.ReadJSON(&incoming)
		if err != nil {
			return
		}

		if incoming.Type == "join_random" {
			c.Hub.JoinRandom(c)
			continue
		}

		if c.Room == nil {
			continue
		}

		encoded, _ := utils.EncodeMessage(&incoming)
		c.Room.Broadcast <- BroadcastMessage{
			Sender: c,
			Data:   encoded,
		}
	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()

	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
