package moves

import "math/rand"

// Directions are the valid directions a snake can move
var Directions = [...]string{"up", "down", "left", "right"}

// RandomMove makes the snake move erratically. This kills the snake.
func RandomMove() string {
	randIndex := rand.Intn(len(Directions))
	move := Directions[randIndex]
	return move
}
