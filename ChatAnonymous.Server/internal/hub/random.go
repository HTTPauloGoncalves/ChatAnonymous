package hub

import "github.com/HTTPauloGoncalves/ChatAnonymous/ChatAnonymous.Server/utils"

func (h *Hub) JoinRandom(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if client.Room != nil {
		return
	}

	for _, c := range h.WaitingRoom {
		if c == client {
			return
		}
	}

	if len(h.WaitingRoom) > 0 {
		other := h.WaitingRoom[0]
		h.WaitingRoom = h.WaitingRoom[1:]

		roomID, _ := utils.NewUUID()
		room := NewRoom(roomID, "")
		h.Rooms[roomID] = room

		client.Room = room
		other.Room = room

		go room.Run(h)

		room.Join <- other
		room.Join <- client

		return
	}

	h.WaitingRoom = append(h.WaitingRoom, client)
}
