package game_screen

import (
	"fmt"
	"snake_and_ladder/MySQL"
	"snake_and_ladder/game_board"
	"snake_and_ladder/players"
)

type GameScreen struct {
	ID              int
	Board           *game_board.Board
	Dices           Dices
	Players         []*players.Player
	TurnNo          int
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
		TurnNo:          0,
		PlayerPositions: players.SetPlayersOnStart(players_),
		GameDeclared:    false,
	}
	gameScreen.Print()
	gameScreen.ID = MySQL.DB.SaveGameScreen(len(gameScreen.Dices), gameScreen.Board.ID, gameScreen.PlayerIDs())
	return gameScreen
}

func (gameScreen *GameScreen) PlayerIDs() []int {
	playerIDs := make([]int, len(gameScreen.Players))
	for index, player := range gameScreen.Players {
		playerIDs[index] = player.ID
	}
	return playerIDs
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
	currentPlayer := gameScreen.CurrentPlayer()

	diceRoll := gameScreen.Dices.Roll()
	currPosition := gameScreen.PlayerPositions[currentPlayer.ID]
	newPosition := gameScreen.Board.Move(currPosition, diceRoll)

	_, jumped := gameScreen.Board.Jumpers[currPosition+diceRoll]
	gameScreen.SaveMove(currentPlayer, diceRoll, currPosition, newPosition, jumped)
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
	gameScreen.TurnNo = gameScreen.TurnNo + 1
}

func (gameScreen *GameScreen) CurrentPlayer() *players.Player {
	return gameScreen.Players[gameScreen.TurnNo%len(gameScreen.Players)]
}

func (gameScreen *GameScreen) DeclareWin(winner *players.Player) {
	winner.DeclareWin()
	MySQL.DB.MarkGameWinner(gameScreen.ID, winner.ID)
	gameScreen.GameDeclared = true
}

func (gameScreen *GameScreen) SaveMove(player *players.Player, diceRoll, start, end int, jumped bool) {
	newMove := &Move{
		GameID:   gameScreen.ID,
		TurnNo:   gameScreen.TurnNo + 1,
		PlayerID: player.ID,
		DiceRoll: diceRoll,
		Start:    start,
		End:      end,
		Jumped:   jumped,
	}
	newMove.SaveMove()
}
