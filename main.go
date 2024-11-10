package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	reset   = "\033[0m"  // Reset color
	whiteBg = "\033[47m" // White background
	blackFg = "\033[30m" // Black foreground (text)
	redFg   = "\033[31m" // Red foreground (text)
)

// Available suits and ranks
var (
	suits = []string{"♠", "♥", "♦", "♣"}
	ranks = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
)

// drawCardRow creates a row for each card line in a horizontal format
func drawCardRow(cards [][]string, row int) string {
	rowString := ""
	for _, card := range cards {
		rowString += whiteBg + card[row] + reset + "  " // Add space between cards
	}
	return rowString
}

// printHand displays a hand of cards horizontally
func printHand(cards [][]string) {
	for row := 0; row < len(cards[0]); row++ { // Draw each row across all cards
		fmt.Println(drawCardRow(cards, row))
	}
}

// generateRandomCard returns a randomly generated card with rank and suit
func generateRandomCard() []string {
	rand.Seed(time.Now().UnixNano())
	rank := ranks[rand.Intn(len(ranks))]
	suit := suits[rand.Intn(len(suits))]

	color := blackFg
	if suit == "♥" || suit == "♦" {
		color = redFg
	}

	// Card representation as a slice of strings
	return []string{
		fmt.Sprintf("%s%s%-2s     ", whiteBg, color, rank), // Rank in top-left corner
		"       ",
		fmt.Sprintf("   %s%s   ", color, suit),
		"       ",
		fmt.Sprintf("     %s%s%2s", whiteBg, color, rank),
	}
}

func main() {
	// Generate a random hand of 6 cards
	var hand [][]string
	for i := 0; i < 13; i++ {
		hand = append(hand, generateRandomCard())
	}

	// Print the hand horizontally
	printHand(hand)
}
