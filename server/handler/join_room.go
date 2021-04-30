package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type JoinRoomRequest struct {
	RoomCode string
	Token    string
}

type JoinRoomResponse struct {
	ID      string `json:"id,omitempty"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func joinRoomEndpoint(poker Poker, logger Logger) http.Handler {
	validate := func(r *http.Request) (*JoinRoomRequest, error) {
		request := JoinRoomRequest{}

		token, err := extractBearerToken(r)
		if err != nil {
			return &request, err
		}

		params := mux.Vars(r)
		roomCode, ok := params["room_code"]
		if ok == false {
			return &request, errors.New("Invalid request param room_code")
		}

		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return &request, err
		}

		if roomCode == "" || len(roomCode) < 6 {
			return &request, errors.New("Invalid request param room_code")
		}

		request.RoomCode = roomCode
		request.Token = token

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

		id, err := poker.JoinRoom(request.RoomCode, request.Token)
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
