package websocket

import (
	"fmt"
	"net/http"

	"github.com/HTTPauloGoncalves/ChatAnonymous/ChatAnonymous.Server/internal/hub"
	"github.com/HTTPauloGoncalves/ChatAnonymous/ChatAnonymous.Server/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebsocketHandler(h *hub.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.EnableCORS(w, r)

		if r.Method == http.MethodOptions {
			return
		}

		uidroom := r.URL.Query().Get("room")
		password := r.URL.Query().Get("password")

		if uidroom == "" {
			http.Error(w, "'room' parameter is mandatory", http.StatusBadRequest)
			return
		}

		if password == "" {
			http.Error(w, "'password' parameter is mandatory", http.StatusForbidden)
			return
		}

		room := h.GetRoom(uidroom)
		if room == nil {
			http.Error(w, "room not found", http.StatusNotFound)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Erro no upgrade:", err)
			return
		}

		room.Join <- conn

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				room.Leave <- conn
				conn.Close()
				return
			}

			room.Broadcast <- hub.Message{
				Conn: conn,
				Data: msg,
			}
		}
	}
}
