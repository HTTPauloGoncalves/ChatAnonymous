package hub

import "github.com/gorilla/websocket"

type Room struct {
	Id        string
	Password  string
	Clients   map[*websocket.Conn]bool
	Join      chan *websocket.Conn
	Leave     chan *websocket.Conn
	Broadcast chan []byte
}

func NewRoom(id string, password string) *Room {
	return &Room{
		Id:        id,
		Password:  password,
		Clients:   make(map[*websocket.Conn]bool),
		Join:      make(chan *websocket.Conn),
		Leave:     make(chan *websocket.Conn),
		Broadcast: make(chan []byte),
	}
}

func (r *Room) Run() {
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
		}
	}
}
