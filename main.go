package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	whiteBg = "\033[47m" // White background
	blackFg = "\033[30m" // Black foreground (text)
	redFg   = "\033[31m" // Red foreground (text)
)

// isGameOver checks if the sum of all players' scores equals 13, indicating the game is over.
func isGameOver(player1Score, player2Score, player3Score, player4Score int) bool {
	totalScore := player1Score + player2Score + player3Score + player4Score
	return totalScore == 13
}

// main game loop
func main() {
	// Generate and shuffle the deck
	deck := generateDeck()

	// Distribute cards to players
	hand1, hand2, hand3, hand4 := distributeCards(deck)

	// Create an instance of players with 4 players
	p := players{
		player: []player{
			{id: 1, hand: hand1, score: 0},
			{id: 2, hand: hand2, score: 0},
			{id: 3, hand: hand3, score: 0},
			{id: 4, hand: hand4, score: 0},
		},
	}

	// Game loop
	for {
		// Players taking turns
		for i := range p.player {
			currentPlayer := &p.player[i]
			fmt.Printf("Game Master: Player %d's turn\n", currentPlayer.id)

			var selectedPlayer *player
			var rankChoice string
			var guess int

			// Handle turn logic based on player
			switch currentPlayer.id {
			case 1:
				// Player 1's turn - human player

				// Print each player's hand after every update
				printHand("Your hand", p.player[0].hand)       // Human player
				printHand("Player 2's hand", p.player[1].hand) // AI bot 1
				printHand("Player 3's hand", p.player[2].hand) // AI bot 2
				printHand("Player 4's hand", p.player[3].hand) // AI bot 3

				var playerChoice int
				fmt.Print("Which player do you want to talk to? (2, 3 or 4): ")
				fmt.Scanln(&playerChoice)

				if playerChoice < 2 || playerChoice > 4 {
					fmt.Println("Invalid player choice. Please try again.")
					continue
				}

				// convert player id to index
				selectedPlayer = &p.player[playerChoice-1]

				fmt.Print("Enter the card rank you want to ask about (A, 2, 3, ..., K): ")
				fmt.Scanln(&rankChoice)
				rankChoice = strings.ToUpper(rankChoice)

				fmt.Printf("How many %s cards does Player %d have? ", rankChoice, selectedPlayer.id)
				fmt.Scanln(&guess)
			default:
				// Bot randomly chooses which player to talk to
				selectedPlayerIndex := selectAnotherRandomPlayer(currentPlayer.id)
				selectedPlayer = &p.player[selectedPlayerIndex]

				// Bot randomly chooses a rank to ask about
				rankChoice = ranks[rand.Intn(len(ranks))]

				// Bot randomly guesses the count
				guess = rand.Intn(4) + 1 // Guess between 1 and 4
			}

			actualCount := countCardsOfRank(selectedPlayer.hand, rankChoice)
			fmt.Printf("   Player %d: Player %d, do you have any cards of rank %s?\n", currentPlayer.id, selectedPlayer.id, rankChoice)

			if actualCount == 0 {
				fmt.Printf("   Player %d: No\n", selectedPlayer.id)
				continue
			}
			fmt.Printf("   Player %d: Yes\n", selectedPlayer.id)

			fmt.Printf("   Player %d: Player %d, do you have %d of these cards\n", currentPlayer.id, selectedPlayer.id, guess)
			if guess != actualCount {
				fmt.Printf("   Player %d: No\n", selectedPlayer.id)
				continue
			}

			fmt.Printf("   Player %d: Yes\n", selectedPlayer.id)

			fmt.Printf("Game Master: Player %d takes %d cards of rank %s from Player %d.\n", currentPlayer.id, actualCount, rankChoice, selectedPlayer.id)
			moveCards(&selectedPlayer.hand, &currentPlayer.hand, rankChoice)

			if checkForCompleteSet(&currentPlayer.hand, rankChoice) {
				currentPlayer.score++
				fmt.Printf("Game Master: Player %d has collected all 4 cards of rank %s and scored 1 point!\n", currentPlayer.id, rankChoice)
			}
		}

		// Check if the game is over
		if isGameOver(p.player[0].score, p.player[1].score, p.player[2].score, p.player[3].score) {
			break
		}
	}

	// Display final score
	fmt.Println("===================================")
	fmt.Println("Final Scores:")
	fmt.Printf("Player 1: %d\n", p.player[0].score)
	fmt.Printf("Player 2: %d\n", p.player[1].score)
	fmt.Printf("Player 3: %d\n", p.player[2].score)
	fmt.Printf("Player 4: %d\n", p.player[3].score)
}
