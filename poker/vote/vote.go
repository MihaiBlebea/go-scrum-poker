package vote

import (
	"time"

	"github.com/gofrs/uuid"
)

type Vote struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"user_id"`
	RoomID    string    `json:"room_id"`
	Vote      uint      `json:"vote"`
	Turn      uint      `json:"turn"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(userID, roomID string, turn, vote uint) (*Vote, error) {
	vt := Vote{}

	id, err := uuid.NewV4()
	if err != nil {
		return &vt, err
	}
	vt.ID = id.String()
	vt.UserID = userID
	vt.RoomID = roomID
	vt.Turn = turn
	vt.Vote = vote

	return &vt, nil
}
