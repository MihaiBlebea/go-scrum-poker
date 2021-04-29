package poker

import (
	"errors"

	"github.com/MihaiBlebea/go-scrum-poker/jwt"
	"github.com/MihaiBlebea/go-scrum-poker/poker/room"
	"github.com/MihaiBlebea/go-scrum-poker/poker/user"
	"github.com/MihaiBlebea/go-scrum-poker/poker/vote"
	"gorm.io/gorm"
)

type Poker interface {
	CreateRoom(name string) (string, error)
	CreateUser(username, email, token string) (string, error)
	AddUser(roomID, username, token string) (string, error)
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

func (p *poker) CreateRoom(name string) (string, error) {
	room, err := room.New(name)
	if err != nil {
		return "", err
	}
	p.roomRepo.Save(room)

	return room.JoinURL, nil
}

func (p *poker) CreateUser(username, email, token string) (string, error) {
	if _, ok := jwt.VerifyJWT(token, "scrumpoker-auth"); ok == false {
		return "", errors.New("Invalid auth token")
	}

	user, err := user.New(username, email, token)
	if err != nil {
		return "", err
	}
	p.userRepo.Save(user)

	return user.ID, nil
}

func (p *poker) AddUser(roomID, username, token string) (string, error) {
	usr, err := user.New(username, "", token)
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

	room, err := p.roomRepo.GetByID(roomID)
	if err != nil {
		return err
	}

	v, err := p.voteRepo.GetLatestVoteForUserInRoom(user.ID, room.ID, room.Turn)
	if err != nil && err.Error() != "record not found" {
		return err
	}

	if v.Vote == points {
		return errors.New("Cannot point twice with the same number of points")
	}

	vote, err := vote.New(userID, roomID, room.Turn, points)
	if err != nil {
		return err
	}
	p.voteRepo.Save(vote)

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
	if roomID == "" {
		return &State{}, errors.New("RoomID is an empty string")
	}

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

	for _, user := range users {
		vote, found := matchUserVote(votes, &user)
		if found == false {
			state.addUserState(user.Username, 0)
			continue
		}
		state.addUserState(user.Username, vote.Vote)
	}

	return state, nil
}
