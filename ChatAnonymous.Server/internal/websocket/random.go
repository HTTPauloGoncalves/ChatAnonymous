package websocket

import (
	"net/http"

	"github.com/HTTPauloGoncalves/ChatAnonymous/ChatAnonymous.Server/internal/hub"
)

func RandomWebsocketHandler(h *hub.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		client := &hub.Client{
			Conn: conn,
			Send: make(chan []byte, 16),
			Hub:  h,
			Room: nil,
		}

		h.JoinRandom(client)

		go client.ReadPump()
		go client.WritePump()
	}
}
