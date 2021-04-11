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

func (r *Repo) GetByUserIDAndTurn(ID string, turn uint) (*Vote, error) {
	vote := Vote{}
	err := r.db.Where("user_id = ? AND turn = ?", ID, turn).Order("created_at desc").Last(&vote).Error

	return &vote, err
}
