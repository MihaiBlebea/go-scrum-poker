package handler

import (
	"net/http"

	"github.com/MihaiBlebea/go-scrum-poker/poker"
)

type Logger interface {
	Info(args ...interface{})
	Trace(args ...interface{})
	Debug(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

type Poker interface {
	CreateRoom() (string, error)
	AddUser(roomID, username string) (string, error)
	GetVoteOptions() []uint
	TurnVotes(roomID string) ([]poker.UserVote, error)
	Vote(roomID, userID string, points uint) error
	NextTurn(roomID string) (uint, error)
}

type Service struct {
	poker  Poker
	logger Logger
}

func New(poker Poker, logger Logger) *Service {
	return &Service{poker, logger}
}

func (s *Service) JoinRoomEndpoint() http.Handler {
	return joinRoomEndpoint(s.poker, s.logger)
}

func (s *Service) CreateRoomEndpoint() http.Handler {
	return createRoomEndpoint(s.poker, s.logger)
}

func (s *Service) VoteOptionsEndpoint() http.Handler {
	return voteOptionsEndpoint(s.poker, s.logger)
}

func (s *Service) RoomUsersEndpoint() http.Handler {
	return roomUsersEndpoint(s.poker, s.logger)
}

func (s *Service) VoteEndpoint() http.Handler {
	return voteEndpoint(s.poker, s.logger)
}

func (s *Service) NextTurnEndpoint() http.Handler {
	return nextTurnEndpoint(s.poker, s.logger)
}

func (s *Service) HealthEndpoint() http.Handler {
	return healthEndpoint(s.logger)
}
