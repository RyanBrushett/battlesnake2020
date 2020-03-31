package structs

// Coordinate refers to a square on a board
type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Add the values of Coordinate other to Coordinate c
func (c Coordinate) Add(other Coordinate) Coordinate {
	result := Coordinate{
		X: c.X + other.X,
		Y: c.Y + other.Y,
	}
	return result
}
