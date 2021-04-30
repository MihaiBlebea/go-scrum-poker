package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func sendResponse(w http.ResponseWriter, resp interface{}, code int, logger Logger) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	b, _ := json.Marshal(resp)

	log := logger.Info

	if code != 200 {
		log = logger.Error
	}

	log(string(b))
	w.Write(b)
}

func extractBearerToken(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return "", errors.New("Expected auth token is missing")
	}

	token := strings.Replace(auth, "Bearer ", "", 1)
	if token == "" {
		return "", errors.New("Expected auth token is missing")
	}

	return token, nil
}
