package poker

type UserState struct {
	Username string `json:"username"`
	Vote     uint   `json:"vote"`
}

type State struct {
	UserState []UserState `json:"user_state"`
	Turn      uint        `json:"current_turn"`
	RoomID    string      `json:"room_id"`
}

func newState(roomID string, turn uint) *State {
	return &State{RoomID: roomID, Turn: turn}
}

func (s *State) addUserState(username string, vote uint) {
	s.UserState = append(s.UserState, UserState{username, vote})
}
