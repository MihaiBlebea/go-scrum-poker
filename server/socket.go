package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type ConnRequest struct {
	RoomID string `json:"room_id"`
	UserID string `json:"user_id"`
}

type Message struct {
	RoomID  string
	UserID  string
	Message string
}

var clients = make(map[string]map[*websocket.Conn]bool)

var Broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID := vars["room_id"]
	if roomID == "" {
		log.Fatal(errors.New("Room id is missing"))
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := clients[roomID]; ok == false {
		fmt.Println("Set initial connection")
		clients[roomID] = make(map[*websocket.Conn]bool)
	}

	clients[roomID][ws] = true
}

func writer(payload Message) {
	Broadcast <- payload
}

func echo() {
	for {
		message := <-Broadcast

		if _, ok := clients[message.RoomID]; ok == false {
			log.Fatal(errors.New("Room id not found in clients"))
		}

		for client := range clients[message.RoomID] {
			err := client.WriteMessage(websocket.TextMessage, []byte("Heeeey"))
			if err != nil {
				log.Printf("Websocket error: %s", err)
				client.Close()
				delete(clients[message.RoomID], client)
			}
		}
	}
}
