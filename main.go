package main

import (
	"fmt"
	"math/rand"
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

// countCardsOfRank counts how many cards of a given rank a player has
func countCardsOfRank(cards [][]string, rank string) int {
	count := 0
	for _, card := range cards {
		fmt.Println("card: ", card[0])
		if strings.Contains(card[0], rank) {
			// if card[0][4:6] == rank { // Check the rank (first part of the card's rank in string)
			count++
		}
	}
	return count
}

// moveCards moves cards from one player's hand to another (here, from another player to player 1)
func moveCards(from *[][]string, to *[][]string, rank string) {
	// Find and remove the cards of the given rank from the 'from' player
	var cardsToMove [][]string
	for i, card := range *from {
		if card[0][4:6] == rank {
			cardsToMove = append(cardsToMove, card)
			*from = append((*from)[:i], (*from)[i+1:]...) // Remove card
			break
		}
	}
	// Add the found cards to the 'to' player's hand
	*to = append(*to, cardsToMove...)
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

	// Game loop for player1 to interact with other players
	for {
		// Ask player1 which AI player they want to talk to
		var playerChoice string
		fmt.Print("Which player do you want to talk to? (player2, player3, player4): ")
		fmt.Scanln(&playerChoice)
		playerChoice = strings.ToLower(playerChoice)

		// Choose the correct player hand based on playerChoice
		var selectedPlayer [][]string
		switch playerChoice {
		case "player2":
			selectedPlayer = player2
		case "player3":
			selectedPlayer = player3
		case "player4":
			selectedPlayer = player4
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
		actualCount := countCardsOfRank(selectedPlayer, rankChoice)

		// Check if the guess was correct
		if guess == actualCount {
			fmt.Printf("Correct! Player 1 takes all %d cards of rank %s from %s.\n", actualCount, rankChoice, playerChoice)
			// Move the cards to player 1's hand
			moveCards(&selectedPlayer, &player1, rankChoice)
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

		// Print updated hands
		printHand("Player 1", player1)
		printHand("Player 2", player2)
		printHand("Player 3", player3)
		printHand("Player 4", player4)
	}
}
