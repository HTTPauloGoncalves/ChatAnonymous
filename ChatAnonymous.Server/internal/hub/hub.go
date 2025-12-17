package hub

import "sync"

type Hub struct {
	mu          sync.RWMutex
	Rooms       map[string]*Room
	WaitingRoom []*Client
}

func NewHub() *Hub {
	return &Hub{
		Rooms:       make(map[string]*Room),
		WaitingRoom: make([]*Client, 0),
	}
}

func (h *Hub) AddNewRoom(uid string, room *Room) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.Rooms[uid]; exists {
		return false
	}

	h.Rooms[uid] = room
	return true
}

func (h *Hub) GetRoom(uid string) *Room {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.Rooms[uid]
}

func (h *Hub) RemoveRoom(uid string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.Rooms, uid)
}
