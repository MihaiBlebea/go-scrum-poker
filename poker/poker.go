package poker

import (
	"errors"

	"github.com/MihaiBlebea/go-scrum-poker/poker/room"
	"github.com/MihaiBlebea/go-scrum-poker/poker/user"
	"github.com/MihaiBlebea/go-scrum-poker/poker/vote"
	"gorm.io/gorm"
)

type Poker interface {
	CreateRoom() (string, error)
	AddUser(roomID, username string) (string, error)
	Vote(roomID, userID string, points uint) error
	NextTurn(roomID string) (uint, error)
	GetVoteOptions() []uint
	TurnVotes(roomID string) ([]UserVote, error)
}

type poker struct {
	voteOptions []uint
	conn        *gorm.DB
}

func New(db *gorm.DB) Poker {
	return &poker{fibonacci(6), db}
}

func (p *poker) CreateRoom() (string, error) {
	rm, err := room.New()
	if err != nil {
		return "", err
	}

	repo := room.NewRepo(p.conn)
	repo.Save(rm)

	return rm.ID, nil
}

func (p *poker) AddUser(roomID, username string) (string, error) {
	usr, err := user.New(roomID, username)
	if err != nil {
		return "", err
	}

	repo := user.NewRepo(p.conn)
	repo.Save(usr)

	return usr.ID, nil
}

func (p *poker) Vote(roomID, userID string, points uint) error {
	repo := user.NewRepo(p.conn)
	user, err := repo.GetByID(userID)
	if err != nil {
		return err
	}

	if user.RoomID != roomID {
		return errors.New("RoomID does not match the user")
	}

	var found bool
	for _, pts := range p.voteOptions {
		if pts == points {
			found = true
		}
	}
	if found == false {
		return errors.New("Vote is not a valid fibonacci number")
	}

	roomRepo := room.NewRepo(p.conn)
	rm, err := roomRepo.GetByID(roomID)
	if err != nil {
		return err
	}

	vt, err := vote.New(userID, roomID, rm.Turn, points)
	if err != nil {
		return err
	}

	voteRepo := vote.NewRepo(p.conn)
	voteRepo.Save(vt)

	return nil
}

func (p *poker) NextTurn(roomID string) (uint, error) {
	repo := room.NewRepo(p.conn)
	rm, err := repo.GetByID(roomID)
	if err != nil {
		return 0, err
	}

	rm.IncrementTurn()
	repo.UpdateTurn(rm)

	return rm.Turn, nil
}

func (p *poker) GetVoteOptions() []uint {
	return p.voteOptions
}

type UserVote struct {
	User user.User `json:"user"`
	Vote vote.Vote `json:"vote"`
}

func (p *poker) TurnVotes(roomID string) ([]UserVote, error) {
	userVotes := make([]UserVote, 0)

	repo := user.NewRepo(p.conn)
	voteRepo := vote.NewRepo(p.conn)
	roomRepo := room.NewRepo(p.conn)

	users, err := repo.GetByRoomID(roomID)
	if err != nil {
		return []UserVote{}, err
	}

	rm, err := roomRepo.GetByID(roomID)
	if err != nil {
		return []UserVote{}, err
	}

	for _, usr := range users {
		vt, err := voteRepo.GetByUserIDAndTurn(usr.ID, rm.Turn)
		if err != nil {
			return []UserVote{}, err
		}

		userVotes = append(userVotes, UserVote{usr, *vt})
	}

	return userVotes, nil
}
