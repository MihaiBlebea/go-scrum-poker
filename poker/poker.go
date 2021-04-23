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
	GetState(roomID string) (*State, error)
}

type poker struct {
	voteOptions []uint
	conn        *gorm.DB
	voteRepo    *vote.Repo
	roomRepo    *room.Repo
	userRepo    *user.Repo
}

func New(db *gorm.DB) Poker {
	return &poker{
		voteOptions: fibonacci(6),
		conn:        db,
		voteRepo:    vote.NewRepo(db),
		roomRepo:    room.NewRepo(db),
		userRepo:    user.NewRepo(db),
	}
}

func (p *poker) CreateRoom() (string, error) {
	rm, err := room.New()
	if err != nil {
		return "", err
	}
	p.roomRepo.Save(rm)

	return rm.ID, nil
}

func (p *poker) AddUser(roomID, username string) (string, error) {
	usr, err := user.New(roomID, username)
	if err != nil {
		return "", err
	}
	p.userRepo.Save(usr)

	return usr.ID, nil
}

func (p *poker) Vote(roomID, userID string, points uint) error {
	user, err := p.userRepo.GetByID(userID)
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

	rm, err := p.roomRepo.GetByID(roomID)
	if err != nil {
		return err
	}

	vt, err := vote.New(userID, roomID, rm.Turn, points)
	if err != nil {
		return err
	}
	p.voteRepo.Save(vt)

	return nil
}

func (p *poker) NextTurn(roomID string) (uint, error) {
	room, err := p.roomRepo.GetByID(roomID)
	if err != nil {
		return 0, err
	}

	room.IncrementTurn()
	p.roomRepo.UpdateTurn(room)

	return room.Turn, nil
}

func (p *poker) GetVoteOptions() []uint {
	return p.voteOptions
}

func (p *poker) GetState(roomID string) (*State, error) {
	room, err := p.roomRepo.GetByID(roomID)
	if err != nil {
		return &State{}, err
	}

	users, err := p.userRepo.GetByRoomID(roomID)
	if err != nil {
		return &State{}, err
	}

	votes, err := p.voteRepo.GetLatestByTurnAndRoomID(room.Turn, room.ID)
	if err != nil {
		return &State{}, err
	}
	state := newState(room.ID, room.Turn)

	for _, vote := range votes {
		for _, user := range users {
			if user.ID == vote.UserID {
				state.addUserState(user.Username, vote.Vote)
			}
		}
	}

	return state, nil
}
