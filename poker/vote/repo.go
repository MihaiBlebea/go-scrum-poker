package vote

import "gorm.io/gorm"

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) Save(vote *Vote) {
	r.db.Create(vote)
}

func (r *Repo) GetByID(ID string) (*Vote, error) {
	vote := Vote{}
	err := r.db.Where("id = ?", ID).Find(&vote).Error

	return &vote, err
}

func (r *Repo) GetLatestByTurnAndRoomID(turn uint, roomID string) ([]Vote, error) {
	votes := make([]Vote, 0)
	err := r.db.Where("room_id = ? AND turn = ?", roomID, turn).Group("user_id").Find(&votes).Error

	return votes, err
}
