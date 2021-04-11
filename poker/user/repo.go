package user

import "gorm.io/gorm"

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) Save(user *User) {
	r.db.Create(user)
}

func (r *Repo) GetByID(ID string) (*User, error) {
	user := User{}
	err := r.db.Where("id = ?", ID).Find(&user).Error

	return &user, err
}

func (r *Repo) GetByRoomID(ID string) ([]User, error) {
	users := make([]User, 0)
	err := r.db.Where("room_id = ?", ID).Find(&users).Error

	return users, err
}
