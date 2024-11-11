package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	logMessageLimit = 30
)

// clearScreen clears the terminal screen.
func clearScreen() {
	fmt.Print("\033[H\033[2J") // ANSI escape code to clear screen and move cursor to top
}

func appendMessageLog(log *[]string, message string) []string {
	if len(*log) >= logMessageLimit {
		*log = (*log)[1:]
	}

	*log = append(*log, message)

	return *log
}

// printLayout prints the updated game layout after each turn.
func printLayout(p *players, messageLog []string) {
	clearScreen()

	// Print score board (top)
	fmt.Println("Score board")
	fmt.Printf("Player 1: %d\n", p.player[0].score)
	fmt.Printf("Player 2: %d\n", p.player[1].score)
	fmt.Printf("Player 3: %d\n", p.player[2].score)
	fmt.Printf("Player 4: %d\n", p.player[3].score)
	fmt.Println(strings.Repeat("=", 115))

	// Print chat messages (middle)
	for _, message := range messageLog {
		fmt.Println(message)
	}

	fmt.Println(strings.Repeat("=", 115))
}

// isGameOver checks if the sum of all players' scores equals 13, indicating the game is over.
func isGameOver(player1Score, player2Score, player3Score, player4Score int) bool {
	totalScore := player1Score + player2Score + player3Score + player4Score
	return totalScore == 13
}

// main game loop
func main() {
	// Generate and shuffle the deck
	deck := generateDeck()

	// var messageLog []string
	messageLog := make([]string, logMessageLimit)

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
			appendMessageLog(&messageLog, fmt.Sprintf("Game Master: Player %d's turn", currentPlayer.id))

			printLayout(&p, messageLog)

			var selectedPlayer *player
			var rankGuess string
			var countGuess int

			// Handle turn logic based on player
			switch currentPlayer.id {
			case 1:
				// Player 1's turn - human player

				// Print each player's hand
				printHand("Your hand", p.player[0].hand) // Human player

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
				fmt.Scanln(&rankGuess)
				rankGuess = strings.ToUpper(rankGuess)

				fmt.Printf("How many %s cards does Player %d have? ", rankGuess, selectedPlayer.id)
				fmt.Scanln(&countGuess)
			default:
				// Bot randomly chooses which player to talk to
				selectedPlayerIndex := selectAnotherRandomPlayer(currentPlayer.id)
				selectedPlayer = &p.player[selectedPlayerIndex]

				// Bot randomly chooses a rank to ask about
				rankGuess = ranks[rand.Intn(len(ranks))]

				// Bot randomly guesses the count
				countGuess = rand.Intn(3) + 1 // Guess between 1 and 3
			}

			actualCount := countCardsOfRank(selectedPlayer.hand, rankGuess)
			appendMessageLog(&messageLog, fmt.Sprintf("   Player %d: Player %d, do you have any cards of rank %s?", currentPlayer.id, selectedPlayer.id, rankGuess))

			if actualCount == 0 {
				appendMessageLog(&messageLog, fmt.Sprintf("   Player %d: No", selectedPlayer.id))
				continue
			}
			appendMessageLog(&messageLog, fmt.Sprintf("   Player %d: Yes", selectedPlayer.id))

			appendMessageLog(&messageLog, fmt.Sprintf("   Player %d: Player %d, do you have %d of these cards?", currentPlayer.id, selectedPlayer.id, countGuess))
			if countGuess != actualCount {
				appendMessageLog(&messageLog, fmt.Sprintf("   Player %d: No", selectedPlayer.id))
				continue
			}

			appendMessageLog(&messageLog, fmt.Sprintf("   Player %d: Yes", selectedPlayer.id))

			appendMessageLog(&messageLog, fmt.Sprintf("Game Master: Player %d takes %d cards of rank %s from Player %d", currentPlayer.id, actualCount, rankGuess, selectedPlayer.id))
			moveCards(&selectedPlayer.hand, &currentPlayer.hand, rankGuess)

			if checkForCompleteSet(&currentPlayer.hand, rankGuess) {
				currentPlayer.score++
				appendMessageLog(&messageLog, fmt.Sprintf("Game Master: Player %d has collected all 4 cards of rank %s and scored 1 point!", currentPlayer.id, rankGuess))
			}
		}

		// Check if the game is over
		if isGameOver(p.player[0].score, p.player[1].score, p.player[2].score, p.player[3].score) {
			break
		}
	}

	// Display final score
	fmt.Println("Game over!\n\nFinal Scores:")
	fmt.Printf("Player 1: %d\n", p.player[0].score)
	fmt.Printf("Player 2: %d\n", p.player[1].score)
	fmt.Printf("Player 3: %d\n", p.player[2].score)
	fmt.Printf("Player 4: %d\n", p.player[3].score)
}
