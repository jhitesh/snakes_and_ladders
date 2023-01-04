package main

import (
	"math"
	"math/rand"
	"snake_and_ladder/game_board"
	"snake_and_ladder/game_screen"
	"snake_and_ladder/players"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	gamePlayers := make([]*players.Player, 0)
	gamePlayers = append(gamePlayers, &players.Player{Name: "Hitesh Jain", ID: 1})
	gamePlayers = append(gamePlayers, &players.Player{Name: "Richa Jain", ID: 2})
	gamePlayers = append(gamePlayers, &players.Player{Name: "Jaishri Jain", ID: 3})
	gamePlayers = append(gamePlayers, &players.Player{Name: "Lokesh Jain", ID: 4})
	gameBoard := game_board.NewBoard(int(math.Floor(math.Sqrt(float64(len(gamePlayers))))) * 100)

	gameScreen := game_screen.SetupGameScreen(gameBoard, gamePlayers)
	gameScreen.PlayUntilWin()
}
