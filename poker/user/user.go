package user

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	RoomID    string    `json:"room_id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(roomId, username string) (*User, error) {
	user := User{}

	id, err := uuid.NewV4()
	if err != nil {
		return &user, err
	}
	user.ID = id.String()
	user.Username = username
	user.RoomID = roomId

	return &user, nil
}
