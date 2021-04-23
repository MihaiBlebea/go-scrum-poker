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
	query := `
		SELECT id, a.user_id, room_id, vote, turn, created_at, updated_at FROM (
			SELECT user_id, MAX(created_at) as max_created
			FROM votes
			WHERE room_id = ? AND turn = ?
			GROUP BY user_id
		) AS a
		INNER JOIN votes as b ON b.user_id = a.user_id AND b.created_at = a.max_created`

	err := r.db.Raw(query, roomID, turn).
		Scan(&votes).
		Error

	return votes, err
}

func (r *Repo) GetLatestVoteForUserInRoom(userID, roomID string) (*Vote, error) {
	vote := Vote{}
	// err := r.db.Where("user_id = ? AND room_id = ?", userID, roomID).Last(&vote).Error

	query := `
		SELECT * FROM votes
			WHERE user_id = ? AND room_id = ?
			ORDER BY created_at DESC
			LIMIT 1`

	err := r.db.Raw(query, userID, roomID).
		Scan(&vote).
		Error

	return &vote, err
}
