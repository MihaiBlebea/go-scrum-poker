package room

import (
	"time"

	"github.com/gofrs/uuid"
)

type Room struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Turn      uint      `json:"turn"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(name string) (*Room, error) {
	room := Room{Name: name}

	id, err := uuid.NewV4()
	if err != nil {
		return &room, err
	}
	room.ID = id.String()
	room.Turn = 1

	return &room, nil
}

func (r *Room) IncrementTurn() {
	r.Turn += 1
}
