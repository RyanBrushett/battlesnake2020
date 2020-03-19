package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RyanBrushett/battlesnake2020/pkg/moves"
	"github.com/RyanBrushett/battlesnake2020/pkg/structs"
)

// Index method
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hey")
}

// Ping method
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

// Start method
func Start(w http.ResponseWriter, r *http.Request) {
	// What do I do with this.
	request := structs.ServerRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	startResponse := structs.StartResponse{
		Color:    structs.Color,
		HeadType: structs.Head,
		TailType: structs.Bum,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(startResponse)
}

// Move method
func Move(w http.ResponseWriter, r *http.Request) {
	request := structs.ServerRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	move := moves.RandomMove()

	moveResponse := structs.MoveResponse{Move: move}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(moveResponse)
}

// End method
func End(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "end")
}
