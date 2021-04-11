package handler

import (
	"net/http"
)

type VoteOptionsResponse struct {
	Votes   []uint `json:"votes"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func voteOptionsEndpoint(poker Poker, logger Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := VoteOptionsResponse{}
		response.Votes = poker.GetVoteOptions()
		response.Success = true

		sendResponse(w, response, 200, logger)
	})
}
