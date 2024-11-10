package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
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
// If the hand contains more than 13 cards, it will be printed in multiple lines with a blank line between the two parts
func printHand(player string, cards [][]string) {
	fmt.Println(player + "'s hand:")
	numCards := len(cards)
	numLines := (numCards + 12) / 13 // Determine how many lines to print

	// Print first 13 cards
	for line := 0; line < numLines; line++ {
		start := line * 13
		end := (line + 1) * 13
		if end > numCards {
			end = numCards
		}
		printCardsInLine := cards[start:end]                  // Slice the hand for the current line
		for row := 0; row < len(printCardsInLine[0]); row++ { // Draw each row across the cards
			fmt.Println(drawCardRow(printCardsInLine, row))
		}
		if line == 0 && numCards > 13 {
			// Add a blank line after the first 13 cards
			fmt.Println()
		}
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

// ExtractRank extracts the rank of the card without color codes or extra spaces
func ExtractRank(card string) string {
	re := regexp.MustCompile(`\033\[[0-9;]*m`)                   // This regex matches any ANSI escape sequence
	cardRank := re.ReplaceAllString(strings.TrimSpace(card), "") // Remove color codes and spaces
	return cardRank
}

// countCardsOfRank counts how many cards of a given rank a player has
func countCardsOfRank(cards [][]string, rank string) int {
	count := 0
	for _, card := range cards {
		// Extract the rank from card[0] without affecting the card itself
		cardRank := ExtractRank(card[0])
		// Directly compare the extracted card rank to the requested rank
		if cardRank == rank {
			count++
		}
	}
	return count
}

// moveCards moves cards of a specific rank from one player's hand to another player's hand
func moveCards(from *[][]string, to *[][]string, rank string) {
	// Find and remove all cards of the given rank from the 'from' player
	var cardsToMove [][]string
	var remainingCards [][]string

	// Loop over the hand and separate the cards of the specific rank
	for _, card := range *from {
		// Extract the rank from card[0]
		cardRank := ExtractRank(card[0]) // Extract rank without formatting issues
		// If the card matches the desired rank, move it to the "cardsToMove" list
		if cardRank == rank {
			cardsToMove = append(cardsToMove, card)
		} else {
			remainingCards = append(remainingCards, card)
		}
	}
	// Update the 'from' player hand by keeping only the remaining cards
	*from = remainingCards
	// Add the found cards to the 'to' player's hand
	*to = append(*to, cardsToMove...)
}

// main game loop
func main() {
	// Generate and shuffle the deck
	deck := generateDeck()

	// Distribute cards to players
	player1, player2, player3, player4 := distributeCards(deck)

	// Game loop for player1 to interact with other players
	for {
		// Print each player's hand after every update
		printHand("Player 1", player1) // Human player
		printHand("Player 2", player2) // AI bot 1
		printHand("Player 3", player3) // AI bot 2
		printHand("Player 4", player4) // AI bot 3

		// Ask player1 which AI player they want to talk to
		var playerChoice string
		fmt.Print("Which player do you want to talk to? (player2, player3, player4): ")
		fmt.Scanln(&playerChoice)
		playerChoice = strings.ToLower(playerChoice)

		// Choose the correct player hand based on playerChoice
		var selectedPlayer *[][]string
		switch playerChoice {
		case "player2":
			selectedPlayer = &player2
		case "player3":
			selectedPlayer = &player3
		case "player4":
			selectedPlayer = &player4
		default:
			fmt.Println("Invalid player choice. Please try again.")
			continue
		}

		// Prompt player1 for a card rank
		var rankChoice string
		fmt.Print("Enter the card rank you want to ask about (A, 2, 3, ..., K): ")
		fmt.Scanln(&rankChoice)
		rankChoice = strings.ToUpper(rankChoice)

		// Prompt player1 to guess how many cards the selected player has of that rank
		var guess int
		fmt.Printf("How many %s cards does %s have? ", rankChoice, playerChoice)
		fmt.Scanln(&guess)

		// Count how many cards the selected player has of the chosen rank
		actualCount := countCardsOfRank(*selectedPlayer, rankChoice)

		// Check if the guess was correct
		if guess == actualCount {
			fmt.Printf("Correct! Player 1 takes all %d cards of rank %s from %s.\n", actualCount, rankChoice, playerChoice)
			// Move the cards to player 1's hand
			moveCards(selectedPlayer, &player1, rankChoice)
		} else {
			fmt.Printf("Wrong! %s has %d cards of rank %s, not %d.\n", playerChoice, actualCount, rankChoice, guess)
		}

		// Ask if player 1 wants to continue
		var continueGame string
		fmt.Print("Do you want to continue? (yes/no): ")
		fmt.Scanln(&continueGame)
		if strings.ToLower(continueGame) != "yes" {
			break
		}
	}
}
