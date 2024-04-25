package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

// House represents a Hogwarts house
type House string

const (
	Gryffindor House = "Gryffindor"
	Hufflepuff House = "Hufflepuff"
	Ravenclaw  House = "Ravenclaw"
	Slytherin  House = "Slytherin"
)

// Question represents a personality question
type Question struct {
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}

// HouseResult represents the result of sorting into a house
type HouseResult struct {
	House House `json:"house"`
}

// Questions to ask the user
var questions = []Question{
	{
		Question: "Which trait do you value the most?",
	},
	{
		Question: "What animal would you bring to Hogwarts as a pet?",
	},
}

func main() {
	http.HandleFunc("/sortinghat", sortingHatHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sortingHatHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Randomly select a house based on answers
	house := selectHouse(r.URL.Query().Get("answers"))

	// Send the result back as JSON
	result := HouseResult{House: house}
	json.NewEncoder(w).Encode(result)
}

// Select a Hogwarts house randomly based on answers
func selectHouse(answers string) House {
	// Hash the string using SHA-256
	hash := sha256.Sum256([]byte(answers))

	// Convert the first 8 bytes of the hash into an int64
	seed := binary.BigEndian.Uint64(hash[:8])
  log.Printf("Seed: %d %s", seed, hash)

  r := rand.New(rand.NewSource(int64(seed)))

	// Randomly select a house
	switch r.Intn(4) {
	case 0:
		return Gryffindor
	case 1:
		return Hufflepuff
	case 2:
		return Ravenclaw
	default:
		return Slytherin
	}
}
