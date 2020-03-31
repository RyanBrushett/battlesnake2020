package structs

// State has a Game, a Turn, a Board, and a Snake who is 'you'
type State struct {
	Game  Game  `json:"game"`
	Turn  int   `json:"turn"`
	Board Board `json:"board"`
	You   Snake `json:"you"`
}

// GetMove returns the next move for the snake
func (state State) GetMove() (move string) {
	return state.sortOfSafeRandom()
}

func (state State) isSafeMove(move string) bool {
	if state.isOutOfBounds(move) {
		return false
	}
	if state.isSnakeCollision(move) {
		return false
	}
	return true
}

func (state State) isOutOfBounds(move string) bool {
	appliedMove := state.applyMove(state.You, move)
	if appliedMove.X >= state.Board.Width || appliedMove.X < 0 {
		return true
	}
	if appliedMove.Y >= state.Board.Height || appliedMove.Y < 0 {
		return true
	}
	return false
}

func (state State) isSnakeCollision(move string) bool {
	snakes := state.Board.Snakes
	appliedMove := state.applyMove(state.You, move)
	for _, snake := range snakes {
		for _, chunk := range snake.Body[:len(snake.Body)-1] {
			if chunk.X == appliedMove.X && chunk.Y == appliedMove.Y {
				return true
			}
		}
	}
	return false
}

func (state State) applyMove(you Snake, move string) (result Coordinate) {
	head := you.Body[0]
	result = head.Add(Moves[move])
	return
}
