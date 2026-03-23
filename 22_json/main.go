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

func main() {
	fmt.Println("Welcome to json in golang")
	EncodedJson()
	DecodeJson()
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
	checkNilErr(err)

	fmt.Println(string(jsonData))
}

func DecodeJson() {
	jsonData := []byte(`
	[
	  {
		"name": "Serena Williams",
		"age": 40,
		"country": "USA",
		"rank": 15,
		"hand": "Right",
		"titles_won": 23
	  },
	  {
		"name": "Ashleigh Barty",
		"age": 26,
		"country": "Australia",
		"rank": 1,
		"hand": "Right",
		"titles_won": 15
	  }
	]
	`)

	var players []player

	isValid := json.Valid(jsonData)

	if isValid {
		err := json.Unmarshal(jsonData, &players)
		checkNilErr(err)
		fmt.Println(players)
	} else {
		fmt.Println("JSON is invalid")
	}

	var playerData []map[string]any

	err := json.Unmarshal(jsonData, &playerData)
	checkNilErr(err)
	fmt.Println(playerData)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
