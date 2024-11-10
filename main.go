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

// Card represents a playing card with a rank, suit, and color
type Card struct {
	rank  string
	suit  string
	color string
}

// drawCardRow creates a row for each card line in a horizontal format
func drawCardRow(cards [][]string, row int) string {
	rowString := ""
	for _, card := range cards {
		rowString += whiteBg + card[row] + reset + "  " // Add space between cards
	}
	return rowString
}

// printHand displays a hand of cards horizontally
func printHand(player string, cards [][]string) {
	fmt.Println(player + "'s hand:")
	for row := 0; row < len(cards[0]); row++ { // Draw each row across all cards
		fmt.Println(drawCardRow(cards, row))
	}
	fmt.Println()
}

// generateCard creates a string representation of a card
func generateCard(card Card) []string {
	color := blackFg
	if card.suit == "♥" || card.suit == "♦" {
		color = redFg
	}

	// Card representation as a slice of strings
	return []string{
		fmt.Sprintf("%s%s%-2s     ", whiteBg, color, card.rank), // Rank in top-left corner
		"       ",
		fmt.Sprintf("   %s%s   ", color, card.suit),
		"       ",
		fmt.Sprintf("     %s%s%2s", whiteBg, color, card.rank),
	}
}

// generateDeck creates a shuffled deck of 52 cards
func generateDeck() []Card {
	var deck []Card
	for _, suit := range suits {
		for _, rank := range ranks {
			color := blackFg
			if suit == "♥" || suit == "♦" {
				color = redFg
			}
			deck = append(deck, Card{rank: rank, suit: suit, color: color})
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

// distributeCards distributes 13 cards to each player from the shuffled deck
func distributeCards(deck []Card) ([][]string, [][]string, [][]string, [][]string) {
	var player1, player2, player3, player4 [][]string
	for i := 0; i < 13; i++ {
		player1 = append(player1, generateCard(deck[i]))
		player2 = append(player2, generateCard(deck[i+13]))
		player3 = append(player3, generateCard(deck[i+26]))
		player4 = append(player4, generateCard(deck[i+39]))
	}
	return player1, player2, player3, player4
}

func main() {
	// Generate and shuffle the deck
	deck := generateDeck()

	// Distribute cards to players
	player1, player2, player3, player4 := distributeCards(deck)

	// Print each player's hand
	printHand("Player 1", player1) // Human player
	printHand("Player 2", player2) // AI bot 1
	printHand("Player 3", player3) // AI bot 2
	printHand("Player 4", player4) // AI bot 3
}
