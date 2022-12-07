package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world (ID: %d)\n", id)
	})

	fmt.Printf("Serving on http://localhost:7070 (ID: %d)\n", id)
	log.Fatal(http.ListenAndServe(":7070", nil))
}
