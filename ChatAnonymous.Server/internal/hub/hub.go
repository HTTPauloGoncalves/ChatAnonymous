package hub

type Hub struct {
	Rooms map[string]*Room
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]*Room),
	}
}

func (h *Hub) AddNewRoom(uid string, room *Room) {
	h.Rooms[uid] = room
}

func (h *Hub) GetRoom(uid string) *Room {
	room, exist := h.Rooms[uid]
	if !exist {
		return nil
	}
	return room
}
