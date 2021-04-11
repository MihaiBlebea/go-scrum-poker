package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

type NextTurnRequest struct {
	RoomID string `json:"room_id"`
}

type NextTurnResponse struct {
	Turn    uint   `json:"turn"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func nextTurnEndpoint(poker Poker, logger Logger) http.Handler {
	validate := func(r *http.Request) (*NextTurnRequest, error) {
		request := NextTurnRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return &request, err
		}

		if request.RoomID == "" || len(request.RoomID) < 3 {
			return &request, errors.New("Invalid request param room_id")
		}

		return &request, nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := NextTurnResponse{}

		request, err := validate(r)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest, logger)
			return
		}

		turn, err := poker.NextTurn(request.RoomID)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest, logger)
			return
		}

		response.Turn = turn
		response.Success = true

		sendResponse(w, response, 200, logger)
	})
}
