package structs

// Color is the colour of the snake
const Color string = "#00fbff"

// Head describes the shape of the Snake's head
const Head string = "tongue"

// Bum describes the shape of the Snake's rear
const Bum string = "curled"

// Game has an ID so you can tell which game is which
type Game struct {
	ID string `json:"id"`
}

// Coordinate refers to a square on a board
type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Board does stuff
type Board struct {
	Height int          `json:"height"`
	Width  int          `json:"width"`
	Snakes []Snake      `json:"snakes"`
	Food   []Coordinate `json:"food"`
}

// Snake is a snake oh shit
type Snake struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Health int          `json:"health"`
	Body   []Coordinate `json:"body"`
	Shout  string       `json:"shout"`
}

// ServerRequest has a Game, a Turn, a Board, and a Snake who is 'you'
type ServerRequest struct {
	Game  Game  `json:"game"`
	Turn  int   `json:"turn"`
	Board Board `json:"board"`
	You   Snake `json:"you"`
}

// StartResponse is what we send back once we get a request to start
type StartResponse struct {
	Color    string `json:"color"`
	HeadType string `json:"headType"`
	TailType string `json:"tailType"`
}

// MoveResponse is what we send back once we get a request to move
type MoveResponse struct {
	Move  string `json:"move"`
	Shout string `json:"shout"`
}
