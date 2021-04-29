package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type CreateUserResponse struct {
	ID      string `json:"id,omitempty"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func createUserEndpoint(poker Poker, logger Logger) http.Handler {
	validate := func(r *http.Request) (*CreateUserRequest, error) {
		request := CreateUserRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return &request, err
		}

		if request.Username == "" || len(request.Username) < 3 {
			return &request, errors.New("Invalid request param username")
		}

		if request.Email == "" || len(request.Email) < 3 {
			return &request, errors.New("Invalid request param email")
		}

		if request.Token == "" || len(request.Token) < 3 {
			return &request, errors.New("Invalid request param token")
		}

		return &request, nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := CreateUserResponse{}

		request, err := validate(r)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest, logger)
			return
		}

		id, err := poker.CreateUser(
			request.Username,
			request.Email,
			request.Token,
		)
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
