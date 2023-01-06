package main

import (
	"math"
	"math/rand"
	"snake_and_ladder/JSON"
	"snake_and_ladder/MySQL"
	"snake_and_ladder/game_board"
	"snake_and_ladder/game_screen"
	"snake_and_ladder/players"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	players_ := JSON.GetPlayers()
	JSON.FillDatabaseConf()
	MySQL.CreateConnection()
	gamePlayers := OrganizePlayers(players_)
	gameBoard := game_board.NewBoard(int(math.Floor(math.Sqrt(float64(len(gamePlayers))))) * 100)

	gameScreen := game_screen.SetupGameScreen(gameBoard, gamePlayers)
	gameScreen.PlayUntilWin()
}

func OrganizePlayers(players_ []string) []*players.Player {
	gamePlayers := make([]*players.Player, len(players_))

	for index, playerName := range players_ {
		player := &players.Player{
			Name: playerName,
		}

		playerID, ok := MySQL.DB.CheckPlayer(player.Name)
		if !ok {
			playerID = MySQL.DB.SavePlayer(player.Name)
		}
		player.ID = playerID
		gamePlayers[index] = player
	}
	return gamePlayers
}
