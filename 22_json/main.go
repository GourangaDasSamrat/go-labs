package main

import (
	"encoding/json"
	"fmt"
)

type player struct {
	Name      string `json:"name"`       // Full Name of the Player
	Age       int    `json:"age"`        // Age of the Player
	Country   string `json:"country"`    // Country of Origin
	Rank      int    `json:"rank"`       // WTA Ranking
	Hand      string `json:"hand"`       // Playing Hand (e.g., Left or Right)
	TitlesWon int    `json:"titles_won"` // Total Titles won in singles
}

func main()  {
    fmt.Println("Welcome to json in golang")
    EncodedJson()
}

func EncodedJson() {
	players := []player{
		{
			Name:      "Serena Williams",
			Age:       40,
			Country:   "USA",
			Rank:      15,
			Hand:      "Right",
			TitlesWon: 23,
		},
		{
			Name:      "Venus Williams",
			Age:       43,
			Country:   "USA",
			Rank:      52,
			Hand:      "Right",
			TitlesWon: 7,
		},
		{
			Name:      "Naomi Osaka",
			Age:       25,
			Country:   "Japan",
			Rank:      45,
			Hand:      "Right",
			TitlesWon: 4,
		},
		{
			Name:      "Simona Halep",
			Age:       31,
			Country:   "Romania",
			Rank:      10,
			Hand:      "Right",
			TitlesWon: 22,
		},
		{
			Name:      "Iga Świątek",
			Age:       21,
			Country:   "Poland",
			Rank:      1,
			Hand:      "Right",
			TitlesWon: 12,
		},
		{
			Name:      "Ashleigh Barty",
			Age:       26,
			Country:   "Australia",
			Rank:      1,
			Hand:      "Right",
			TitlesWon: 15,
		},
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(players, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	fmt.Println(string(jsonData))
}
