package websocket

import (
	"fmt"
	"net/http"

	"github.com/HTTPauloGoncalves/ChatAnonymous/ChatAnonymous.Server/internal/hub"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebsocketHandler(h *hub.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodOptions {
			return
		}

		uidroom := r.URL.Query().Get("room")
		password := r.URL.Query().Get("password")

		if uidroom == "" || password == "" {
			http.Error(w, "'room' and 'password' parameters are mandatory", http.StatusBadRequest)
			return
		}

		room := h.GetRoom(uidroom)
		if room == nil {
			http.Error(w, "room not found", http.StatusNotFound)
			return
		}

		if room.Password != password {
			http.Error(w, "invalid password", http.StatusForbidden)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Erro no upgrade:", err)
			return
		}

		client := &hub.Client{
			Conn: conn,
			Send: make(chan []byte, 16),
			Room: room,
		}

		room.Join <- client

		go client.ReadPump()
		go client.WritePump()
	}
}
