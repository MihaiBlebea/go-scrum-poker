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

	fmt.Println("IS CONNECTING")

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// _, b, err := ws.ReadMessage()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// request := ConnRequest{}
	// err = json.Unmarshal(b, &request)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(request)

	// register client
	// clients[roomID] = append(clients[roomID], ws)
	if _, ok := clients[roomID]; ok == false {
		fmt.Println("Set initial connection")
		clients[roomID] = make(map[*websocket.Conn]bool)
	}

	clients[roomID][ws] = true
	fmt.Println("CLIENTS on CONN", clients)
	// for client := range clients {
	// 	err := client.WriteMessage(websocket.TextMessage, []byte("helloooo back"))
	// 	if err != nil {
	// 		log.Printf("Websocket error: %s", err)
	// 		client.Close()
	// 		delete(clients, client)
	// 	}
	// }
	// writer(Message{roomID: roomID, userID: "1234", message: "Hellooooo"})

	// for client := range clients[roomID] {
	// 	err := client.WriteMessage(websocket.TextMessage, []byte("Heeeey"))
	// 	if err != nil {
	// 		log.Printf("Websocket error: %s", err)
	// 		client.Close()
	// 		// clients[message.roomID] = remove(clients[message.roomID], index)
	// 		delete(clients[roomID], client)
	// 	}
	// }
}

func writer(payload Message) {
	fmt.Println("Send message")
	Broadcast <- payload
}

func echo() {
	for {
		message := <-Broadcast
		fmt.Println("Received message to send", message)

		fmt.Println("CLIENTS", clients)

		if _, ok := clients[message.RoomID]; ok == false {
			log.Fatal(errors.New("Room id not found in clients"))
		}

		// latlong := fmt.Sprintf("%f %f %s", val.Lat, val.Long)

		for client := range clients[message.RoomID] {
			err := client.WriteMessage(websocket.TextMessage, []byte("Heeeey"))
			if err != nil {
				log.Printf("Websocket error: %s", err)
				client.Close()
				// clients[message.roomID] = remove(clients[message.roomID], index)
				delete(clients[message.RoomID], client)
			}
		}
	}
}
