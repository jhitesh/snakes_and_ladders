package MySQL

type Database interface {
	SaveBoard(size, numOfJumpers int, snakes, ladders []byte) (boardId int)

	CheckPlayer(playerName string) (playersId int, isPresent bool)

	SavePlayer(playerName string) (playerId int)

	SaveGameScreen(numOfDices, boardID int, playerIDs []int) (screenId int)

	MarkGameWinner(screenID, winnerID int)

	SaveGameMove(gameId, turnNo, playerId, diceRoll, start, end int, jumped bool)
}
