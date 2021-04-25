package poker

import (
	"github.com/MihaiBlebea/go-scrum-poker/poker/user"
	"github.com/MihaiBlebea/go-scrum-poker/poker/vote"
)

func fibonacci(n int) []uint {
	a := 1
	b := 1
	list := make([]uint, 0)

	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
		list = append(list, uint(a))
	}

	return list
}

func matchUserVote(votes []vote.Vote, u *user.User) (_ *vote.Vote, found bool) {
	for _, v := range votes {
		if v.UserID == u.ID {
			return &v, true
		}
	}

	return nil, false
}
