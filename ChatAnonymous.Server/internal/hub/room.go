package hub

import (
	"time"
)

type Room struct {
	Id        string
	Password  string
	Random    bool
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

func NewRoom(id, password string, random bool) *Room {
	return &Room{
		Id:        id,
		Password:  password,
		Random:    random,
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

			systemMsg := []byte(`{"type":"system","message":"Um usuário entrou."}`)
			r.broadcastSystem(systemMsg)

		case client := <-r.Leave:

			if _, ok := r.Clients[client]; !ok {
				continue
			}

			delete(r.Clients, client)
			close(client.Send)

			if r.Random {
				r.CloseRoom(h)
				return
			}

			r.broadcastSystem([]byte(`{"type":"system","message":"Um usuário saiu."}`))

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
		select {
		case client.Send <- msg:
		default:
		}
	}
}

func (r *Room) CloseRoom(h *Hub) {
	for client := range r.Clients {
		select {
		case client.Send <- []byte(`{"type":"system","message":"Sala encerrada."}`):
		default:
		}

		select {
		case <-client.Send:
		default:
			close(client.Send)
		}
	}

	h.RemoveRoom(r.Id)
}
