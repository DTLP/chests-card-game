package main

import (
	"math/rand"
)

type player struct {
	id    int
	hand  [][]string
	score int
}

type players struct {
	player []player
}

// getRandomPlayer returns a random player ID between 1 and 4
func getRandomPlayer() int {
	return rand.Intn(4)
}

// selectAnotherRandomPlayer returns a random player ID that is not equal to
// current player's ID
func selectAnotherRandomPlayer(currentPlayerID int) int {
	selectedPlayerIndex := getRandomPlayer()
	// make sure the randomly selected player is not the current player
	// current player's ID needs to be converted to index for comparison
	if selectedPlayerIndex == currentPlayerID-1 {
		return selectAnotherRandomPlayer(currentPlayerID)
	}
	return selectedPlayerIndex
}
