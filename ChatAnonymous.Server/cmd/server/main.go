package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HTTPauloGoncalves/ChatAnonymous/ChatAnonymous.Server/internal/hub"
	"github.com/HTTPauloGoncalves/ChatAnonymous/ChatAnonymous.Server/internal/middleware"
	"github.com/HTTPauloGoncalves/ChatAnonymous/ChatAnonymous.Server/internal/websocket"
	"github.com/HTTPauloGoncalves/ChatAnonymous/ChatAnonymous.Server/utils"
)

var h = hub.NewHub()

func main() {
	serverRun()
}

func serverRun() {

	http.Handle(
		"/",
		middleware.Chain(
			http.HandlerFunc(home),
			middleware.EnableCORS(),
			middleware.RateLimit(),
		),
	)

	http.Handle(
		"/newroom",
		middleware.Chain(
			http.HandlerFunc(newRoom),
			middleware.EnableCORS(),
			middleware.RateLimit(),
		),
	)

	http.Handle(
		"/closeroom",
		middleware.Chain(
			http.HandlerFunc(closeRoom),
			middleware.EnableCORS(),
			middleware.RateLimit(),
		),
	)

	http.Handle(
		"/ws",
		middleware.Chain(
			websocket.WebsocketHandler(h),
			middleware.EnableCORS(),
		),
	)

	fmt.Println("Servidor rodando em http://localhost:8080 ...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"API is running"}`))
}

func newRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
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

	go room.Run(h)

	response := map[string]string{
		"roomUUID": uidroom,
		"password": uidpass,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func closeRoom(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	roomUUID := r.URL.Query().Get("room")
	password := r.URL.Query().Get("password")

	if roomUUID == "" {
		http.Error(w, "'room' parameter is mandatory", http.StatusBadRequest)
		return
	}

	if password == "" {
		http.Error(w, "'password' parameter is mandatory", http.StatusBadRequest)
		return
	}

	room := h.GetRoom(roomUUID)
	if room == nil {
		http.Error(w, "room not found", http.StatusNotFound)
		return
	}

	if room.Password != password {
		http.Error(w, "invalid password", http.StatusForbidden)
		return
	}

	room.Close <- true

	json.NewEncoder(w).Encode(map[string]string{
		"message": "room closed successfully",
	})
}
