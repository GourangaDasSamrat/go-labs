package main

import (
	"encoding/json"
	"fmt"
)

type player struct {
	Name          string `json:"Player name"`
	Age           int
	Country       string
	Rank          int
	Hand          string
	TitlesWon     int
	PlayingStatus string `json:"-"`
	// PlayingStatus string `json:"-"` // this value not show on return
}

func main() {
	fmt.Println("Welcome to JSON in Go")
	EncodeJson()
}

func EncodeJson() {
	players := []player{
		{
			Name:          "Serena Williams",
			Age:           40,
			Country:       "USA",
			Rank:          15,
			Hand:          "Right",
			TitlesWon:     23,
			PlayingStatus: "Active",
		},
		{
			Name:          "Venus Williams",
			Age:           43,
			Country:       "USA",
			Rank:          52,
			Hand:          "Right",
			TitlesWon:     7,
			PlayingStatus: "Retired",
		},
		{
			Name:          "Naomi Osaka",
			Age:           25,
			Country:       "Japan",
			Rank:          45,
			Hand:          "Right",
			TitlesWon:     4,
			PlayingStatus: "Active",
		},
		{
			Name:          "Simona Halep",
			Age:           31,
			Country:       "Romania",
			Rank:          10,
			Hand:          "Right",
			TitlesWon:     22,
			PlayingStatus: "Active",
		},
	}

	jsonData, err := json.MarshalIndent(players, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
