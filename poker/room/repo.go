package room

import (
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) Save(room *Room) {
	r.db.Create(room)
}

func (r *Repo) UpdateTurn(room *Room) {
	r.db.Model(&Room{}).Where("id = ?", room.ID).Update("turn", room.Turn)
}

func (r *Repo) GetByID(ID string) (*Room, error) {
	room := Room{}
	err := r.db.Where("id = ?", ID).Find(&room).Error

	return &room, err
}

func (r *Repo) GetByJoinCode(code string) (*Room, error) {
	room := Room{}
	err := r.db.Where("join_code = ?", code).Find(&room).Error

	return &room, err
}
