package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

type CreateRoomRequest struct {
	Name string `json:"name"`
}

type CreateRoomResponse struct {
	JoinURL string `json:"join_url,omitempty"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func createRoomEndpoint(poker Poker, logger Logger) http.Handler {
	validate := func(r *http.Request) (*CreateRoomRequest, error) {
		request := CreateRoomRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return &request, err
		}

		if request.Name == "" {
			return &request, errors.New("Invalid request param name")
		}

		return &request, nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := CreateRoomResponse{}

		request, err := validate(r)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest, logger)
			return
		}

		joinUrl, err := poker.CreateRoom(request.Name)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest, logger)
			return
		}

		response.JoinURL = joinUrl
		response.Success = true

		sendResponse(w, response, 200, logger)
	})
}
