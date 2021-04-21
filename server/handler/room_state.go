package handler

import (
	"errors"
	"net/http"

	"github.com/MihaiBlebea/go-scrum-poker/poker"
)

type RoomUsersRequest struct {
	RoomID string `json:"room_id"`
}

type RoomUsersResponse struct {
	State   poker.State `json:"state"`
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
}

func roomStateEndpoint(poker Poker, logger Logger) http.Handler {
	validate := func(r *http.Request) (*RoomUsersRequest, error) {
		request := RoomUsersRequest{}

		roomID := r.URL.Query().Get("room_id")

		if roomID == "" || len(roomID) < 3 {
			return &request, errors.New("Invalid request param room_id")
		}

		request.RoomID = roomID

		return &request, nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := RoomUsersResponse{}

		request, err := validate(r)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest, logger)
			return
		}

		state, err := poker.GetState(request.RoomID)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest, logger)
			return
		}

		response.State = *state
		response.Success = true

		sendResponse(w, response, 200, logger)
	})
}
