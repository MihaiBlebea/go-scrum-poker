package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

type JoinRoomRequest struct {
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}

type JoinRoomResponse struct {
	ID      string `json:"id,omitempty"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func joinRoomEndpoint(poker Poker, logger Logger) http.Handler {
	validate := func(r *http.Request) (*JoinRoomRequest, error) {
		request := JoinRoomRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return &request, err
		}

		if request.RoomID == "" || len(request.RoomID) < 3 {
			return &request, errors.New("Invalid request param room_id")
		}

		if request.Username == "" || len(request.Username) < 3 {
			return &request, errors.New("Invalid request param username")
		}

		return &request, nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := JoinRoomResponse{}

		request, err := validate(r)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest, logger)
			return
		}

		id, err := poker.AddUser(request.RoomID, request.Username)
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
