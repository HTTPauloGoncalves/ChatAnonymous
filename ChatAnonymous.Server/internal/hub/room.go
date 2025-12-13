package hub

import (
	"time"

	"github.com/gorilla/websocket"
)

type Room struct {
	Id        string
	Password  string
	Clients   map[*Client]bool
	Join      chan *Client
	Leave     chan *Client
	Broadcast chan BroadcastMessage
	Close     chan bool
}

type BroadcastMessage struct {
	Sender *Client
	Data   []byte
}

func NewRoom(id, password string) *Room {
	return &Room{
		Id:        id,
		Password:  password,
		Clients:   make(map[*Client]bool),
		Join:      make(chan *Client),
		Leave:     make(chan *Client),
		Broadcast: make(chan BroadcastMessage),
		Close:     make(chan bool),
	}
}

func (r *Room) Run(h *Hub) {
	timer := time.After(1 * time.Hour)

	for {
		select {
		case client := <-r.Join:
			r.Clients[client] = true

			systemMsg := []byte(`{"username":"System","message":"Um usuário entrou."}`)
			r.broadcastSystem(systemMsg)

		case client := <-r.Leave:
			if _, ok := r.Clients[client]; ok {
				delete(r.Clients, client)
				close(client.Send)

				systemMsg := []byte(`{"username":"System","message":"Um usuário saiu."}`)
				r.broadcastSystem(systemMsg)
			}

		case msg := <-r.Broadcast:
			for client := range r.Clients {
				if client == msg.Sender {
					continue
				}
				select {
				case client.Send <- msg.Data:
				default:
					close(client.Send)
					delete(r.Clients, client)
				}
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

func (r *Room) broadcastSystem(msg []byte) {
	for client := range r.Clients {
		client.Send <- msg
	}
}

func (r *Room) CloseRoom(h *Hub) {
	for client := range r.Clients {
		client.Conn.WriteMessage(websocket.TextMessage, []byte("Sala encerrada."))
		client.Conn.Close()
	}

	h.RemoveRoom(r.Id)
}
