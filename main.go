package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Player struct to represent a player
type Player struct {
	ID     int
	Dices  []int
	Points int
}

// RollDices rolls dices for each player
func RollDices(players []*Player) {
	for _, player := range players {
		for i := range player.Dices {
			player.Dices[i] = rand.Intn(6) + 1
		}
	}
}

// EvaluateDices evaluates the dices rolled by each player
func EvaluateDices(players []*Player) {
	for _, player := range players {
		for i := 0; i < len(player.Dices); i++ {
			switch player.Dices[i] {
			case 6:
				player.Points++
				player.Dices = append(player.Dices[:i], player.Dices[i+1:]...)
				i--
			case 1:
				nextPlayerID := (player.ID + 1) % len(players)
				players[nextPlayerID].Dices = append(players[nextPlayerID].Dices, 1)
				player.Dices = append(player.Dices[:i], player.Dices[i+1:]...)
				i--
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var numPlayers, numDices int
	fmt.Print("Pemain = ")
	fmt.Scanln(&numPlayers)
	fmt.Print("Dadu = ")
	fmt.Scanln(&numDices)

	// Initialize players
	players := make([]*Player, numPlayers)
	for i := range players {
		players[i] = &Player{
			ID:    i,
			Dices: make([]int, numDices),
		}
	}

	// Game loop
	round := 1
	for {
		fmt.Printf("==================\nGiliran %d lempar dadu:\n", round)

		RollDices(players)
		for _, player := range players {
			dicesStr := ""
			if len(player.Dices) == 0 {
				dicesStr = "_ (Berhenti bermain karena tidak memiliki dadu)"
			} else {
				for i, dice := range player.Dices {
					if i > 0 {
						dicesStr += ", "
					}
					dicesStr += fmt.Sprintf("%d", dice)
				}
			}
			fmt.Printf("Pemain #%d (%d): %v\n", player.ID+1, player.Points, dicesStr)
		}

		EvaluateDices(players)
		fmt.Println("Setelah evaluasi:")
		for _, player := range players {
			dicesStr := ""
			if len(player.Dices) == 0 {
				dicesStr = "_ (Berhenti bermain karena tidak memiliki dadu)"
			} else {
				for i, dice := range player.Dices {
					if i > 0 {
						dicesStr += ", "
					}
					dicesStr += fmt.Sprintf("%d", dice)
				}
			}
			fmt.Printf("Pemain #%d (%d): %v\n", player.ID+1, player.Points, dicesStr)
		}

		// Check if game ends
		activePlayers := 0
		var winner *Player
		for _, player := range players {
			if len(player.Dices) > 0 {
				activePlayers++
				winner = player
			}
		}
		if activePlayers == 1 {
			fmt.Printf("==================\nGame berakhir karena hanya pemain #%d yang memiliki dadu.\n", winner.ID+1)
			break
		}

		round++
	}

	// Determine winner
	maxPoints := 0
	var gameWinner *Player
	for _, player := range players {
		if player.Points > maxPoints {
			maxPoints = player.Points
			gameWinner = player
		}
	}
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n", gameWinner.ID+1)
}
