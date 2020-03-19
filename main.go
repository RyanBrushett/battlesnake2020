package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RyanBrushett/battlesnake2020/pkg/handlers"
)

func main() {
	port := "8080"

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/ping", handlers.Ping)

	fmt.Printf("Starting at http://0.0.0.0:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
