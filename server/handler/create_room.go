package handler

import (
	"net/http"
)

type CreateRoomRequest struct {
	Name string `json:"room_name"`
}

type CreateRoomResponse struct {
	ID      string `json:"id,omitempty"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func createRoomEndpoint(poker Poker, logger Logger) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := CreateRoomResponse{}

		id, err := poker.CreateRoom()
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest, logger)
			return
		}

		response.ID = id
		response.Success = true

		sendResponse(w, response, 200, logger)
	})
}
