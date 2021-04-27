package user

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	RoomID    string    `json:"room_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(username, email, token string) (*User, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return &User{}, err
	}

	return &User{
		ID:       id.String(),
		Username: username,
		Email:    email,
		Token:    token,
	}, nil
}
