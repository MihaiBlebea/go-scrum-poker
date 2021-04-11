package server

import (
	"fmt"
	"log"

	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const prefix = "/api/v1/"

func server(handler Handler, logger Logger) {
	r := mux.NewRouter()

	api := r.PathPrefix(prefix).Subrouter()

	// Handle api calls
	api.Handle("/health-check", handler.HealthEndpoint()).
		Methods(http.MethodGet)

	api.Handle("/room", handler.CreateRoomEndpoint()).
		Methods(http.MethodPost)

	api.Handle("/room/user", handler.JoinRoomEndpoint()).
		Methods(http.MethodPost)

	api.Handle("/room/users", handler.RoomUsersEndpoint()).
		Methods(http.MethodGet)

	api.Handle("/room/next", handler.NextTurnEndpoint()).
		Methods(http.MethodPost)

	api.Handle("/vote", handler.VoteEndpoint()).
		Methods(http.MethodPost)

	api.Handle("/votes", handler.VoteOptionsEndpoint()).
		Methods(http.MethodGet)

	// Handle static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./webapp/dist/static"))))

	r.Handle("/", singlePageApp()).Methods(http.MethodGet)

	// Handle websocket
	r.HandleFunc("/ws/{room_id}", wsHandler)
	go echo()

	r.Use(loggerMiddleware(logger))

	srv := &http.Server{
		Handler:      cors.Default().Handler(r),
		Addr:         fmt.Sprintf("0.0.0.0:%s", os.Getenv("HTTP_PORT")),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info(fmt.Sprintf("Started server on port %s", os.Getenv("HTTP_PORT")))

	log.Fatal(srv.ListenAndServe())
}

func singlePageApp() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(200)

		http.ServeFile(w, r, "./webapp/dist/index.html")
	})
}
