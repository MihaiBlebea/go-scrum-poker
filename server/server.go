package server

import (
	"fmt"
	"io/fs"
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

	api.Handle("/room", handler.RoomStateEndpoint()).
		Methods(http.MethodGet)

	api.Handle("/room/next", handler.NextTurnEndpoint()).
		Methods(http.MethodPost)

	api.Handle("/vote", handler.VoteEndpoint()).
		Methods(http.MethodPost)

	api.Handle("/votes", handler.VoteOptionsEndpoint()).
		Methods(http.MethodGet)

	// Handle static files
	webapp, err := fs.Sub(static, "webapp")
	if err != nil {
		fmt.Println(err)
	}

	static, err := fs.Sub(webapp, "static")
	if err != nil {
		fmt.Println(err)
	}

	r.PathPrefix("/static/css/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(static))))
	r.PathPrefix("/static/js/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(static))))

	r.Handle("/", http.FileServer(http.FS(webapp)))

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
