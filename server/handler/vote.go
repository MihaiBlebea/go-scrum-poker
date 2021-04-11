package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/MihaiBlebea/go-scrum-poker/server"
)

type VoteRequest struct {
	RoomID string `json:"room_id"`
	UserID string `json:"user_id"`
	Vote   uint   `json:"vote"`
}

type VoteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func voteEndpoint(poker Poker, logger Logger) http.Handler {
	validate := func(r *http.Request) (*VoteRequest, error) {
		request := VoteRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return &request, err
		}

		if request.RoomID == "" || len(request.RoomID) < 3 {
			return &request, errors.New("Invalid request param room_id")
		}

		if request.UserID == "" || len(request.UserID) < 3 {
			return &request, errors.New("Invalid request param user_id")
		}

		return &request, nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := VoteResponse{}

		request, err := validate(r)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest, logger)
			return
		}

		err = poker.Vote(request.RoomID, request.UserID, request.Vote)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest, logger)
			return
		}

		response.Success = true

		server.Broadcast <- server.Message{
			RoomID:  request.RoomID,
			UserID:  request.UserID,
			Message: "HEY",
		}

		sendResponse(w, response, 200, logger)
	})
}
