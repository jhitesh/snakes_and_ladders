package game_screen

import "snake_and_ladder/MySQL"

type Move struct {
	GameID   int
	TurnNo   int
	PlayerID int
	DiceRoll int
	Start    int
	End      int
	Jumped   bool
}

func (move *Move) SaveMove() {
	MySQL.DB.SaveGameMove(move.GameID, move.TurnNo, move.PlayerID, move.DiceRoll, move.Start, move.End, move.Jumped)
}
