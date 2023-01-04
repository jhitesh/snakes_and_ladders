package game_screen

import (
	"fmt"
	"snake_and_ladder/Golang/game_board"
	"snake_and_ladder/Golang/players"
)

type GameScreen struct {
	Board           *game_board.Board
	Dices           Dices
	Players         []*players.Player
	PlayerTurn      int
	PlayerPositions map[int]int
	GameDeclared    bool
}

func SetupGameScreen(board *game_board.Board, players_ []*players.Player) *GameScreen {
	if len(players_) < 2 {
		return nil
	}
	gameScreen := &GameScreen{
		Board:           board,
		Dices:           GetDices(len(players_)),
		Players:         players_,
		PlayerTurn:      0,
		PlayerPositions: players.SetPlayersOnStart(players_),
		GameDeclared:    false,
	}
	gameScreen.Print()
	return gameScreen
}

func (gameScreen *GameScreen) Print() {
	gameScreen.Board.Print()
	fmt.Printf("Number of dices assigned to the game: %v\n", len(gameScreen.Dices))
	fmt.Printf("Players playing in the game:\n")
	for _, player := range gameScreen.Players {
		player.Print()
	}
}

func (gameScreen *GameScreen) PlayUntilWin() {
	for !gameScreen.GameDeclared {
		gameScreen.PlayCurrentPlayer()
	}
}

func (gameScreen *GameScreen) PlayCurrentPlayer() {
	currentPlayer := gameScreen.Players[gameScreen.PlayerTurn]

	diceRoll := gameScreen.Dices.Roll()
	currPosition := gameScreen.PlayerPositions[currentPlayer.ID]
	newPosition := gameScreen.Board.Move(currPosition, diceRoll)

	_, jumped := gameScreen.Board.Jumpers[currPosition+diceRoll]
	currentPlayer.PrintMove(currPosition, newPosition, diceRoll, jumped)

	if newPosition == gameScreen.Board.Size {
		gameScreen.DeclareWin(currentPlayer)
	}

	gameScreen.updatePlayerPosition(currentPlayer, newPosition)

	gameScreen.updatePlayerTurn()
}

func (gameScreen *GameScreen) updatePlayerPosition(player *players.Player, position int) {
	gameScreen.PlayerPositions[player.ID] = position
}

func (gameScreen *GameScreen) updatePlayerTurn() {
	playerTurn := gameScreen.PlayerTurn + 1
	if playerTurn >= len(gameScreen.Players) {
		playerTurn = 0
	}
	gameScreen.PlayerTurn = playerTurn
}

func (gameScreen *GameScreen) DeclareWin(winner *players.Player) {
	winner.DeclareWin()
	gameScreen.GameDeclared = true
}
