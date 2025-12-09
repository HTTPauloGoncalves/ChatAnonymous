package hub

import (
	"time"

	"github.com/gorilla/websocket"
)

type Room struct {
	Id        string
	Password  string
	Clients   map[*websocket.Conn]bool
	Join      chan *websocket.Conn
	Leave     chan *websocket.Conn
	Broadcast chan []byte
	Close     chan bool
}

func (r *Room) Run(h *Hub) {
	timer := time.After(1 * time.Hour)
	for {
		select {
		case conn := <-r.Join:
			r.Clients[conn] = true

		case conn := <-r.Leave:
			delete(r.Clients, conn)

		case msg := <-r.Broadcast:
			for client := range r.Clients {
				client.WriteMessage(websocket.TextMessage, msg)
			}

		case <-timer:
			r.CloseRoom(h)
			return

		case <-r.Close:
			r.CloseRoom(h)
			return
		}
	}
}

func NewRoom(id string, password string) *Room {
	return &Room{
		Id:        id,
		Password:  password,
		Clients:   make(map[*websocket.Conn]bool),
		Join:      make(chan *websocket.Conn),
		Leave:     make(chan *websocket.Conn),
		Broadcast: make(chan []byte),
		Close:     make(chan bool),
	}
}

func (r *Room) CloseRoom(h *Hub) {
	for client := range r.Clients {
		client.WriteMessage(websocket.TextMessage, []byte("Sala encerrada."))
		client.Close()
	}

	h.RemoveRoom(r.Id)
}
