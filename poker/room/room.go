package room

import (
	"math/rand"
	"time"

	"github.com/gofrs/uuid"
)

type Room struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Turn      uint      `json:"turn"`
	JoinURL   string    `json:"join_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(name string) (*Room, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return &Room{}, err
	}

	return &Room{
		ID:      id.String(),
		Name:    name,
		Turn:    1,
		JoinURL: generateJoinURL(6),
	}, nil
}

func (r *Room) IncrementTurn() {
	r.Turn += 1
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func generateJoinURL(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
