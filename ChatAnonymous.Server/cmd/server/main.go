package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HTTPauloGoncalves/ChatAnonymous/ChatAnonymous.Server/internal/hub"
	"github.com/HTTPauloGoncalves/ChatAnonymous/ChatAnonymous.Server/internal/websocket"
	"github.com/HTTPauloGoncalves/ChatAnonymous/ChatAnonymous.Server/utils"
)

var h = hub.NewHub()

func main() {
	serverRun()
}

func serverRun() {
	http.HandleFunc("/", home)
	http.HandleFunc("/newroom", newRoom)
	http.HandleFunc("/ws", websocket.WebsocketHandler(h))

	fmt.Println("Servidor rodando em http://localhost:8080 ...")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg := []byte("Welcome to Chat Anonymous")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write(msg)

	if err != nil {
		fmt.Println("erro ao escrever resposta:", err)
	}
}

func newRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uidroom, err := utils.NewUUID()

	if err != nil {
		panic("error creating room id")
	}

	uidpass, err := utils.NewUUID()
	if err != nil {
		panic("error creating password")
	}

	room := hub.NewRoom(uidroom, uidpass)

	h.AddNewRoom(uidroom, room)

	go room.Run()

	response := map[string]string{
		"roomUUID": uidroom,
		"password": uidpass,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
