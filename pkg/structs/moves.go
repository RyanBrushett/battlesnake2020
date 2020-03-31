package structs

import (
	"math/rand"
)

// Directions are the valid directions a snake can move
var Directions = [...]string{"up", "down", "left", "right"}

// Moves is Directions mapped to coordinates
var Moves = map[string]Coordinate{
	"up": Coordinate{
		X: 0, Y: -1,
	},
	"down": Coordinate{
		X: 0, Y: 1,
	},
	"left": Coordinate{
		X: -1, Y: 0,
	},
	"right": Coordinate{
		X: 1, Y: 0,
	},
}

// randomMove makes the snake move erratically. This kills the snake.
func (state State) randomMove() string {
	randIndex := rand.Intn(len(Directions))
	move := Directions[randIndex]
	return move
}

// sortOfSafeRandom will avoid immediate death but this still kills the
func (state State) sortOfSafeRandom() string {
	for {
		randIndex := rand.Intn(len(Directions))
		move := Directions[randIndex]
		if state.isSafeMove(move) {
			return move
		}
	}
}
