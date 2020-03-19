package handlers

import (
	"fmt"
	"net/http"
)

// Index method
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hey")
}

// Ping method
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
