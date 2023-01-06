package MySQL

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"snake_and_ladder/JSON"
)

type DatabaseConn struct {
	Conn *sql.DB
}

var _ Database = (*DatabaseConn)(nil)

var DB = &DatabaseConn{}

func CreateConnection() {
	var err error
	dbConf := JSON.DBConf
	dataSourceName := fmt.Sprintf("%v:%v@%v/%v", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Database)
	dbConn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("ERROR in making connection to database", err.Error())
	}
	fmt.Println("Successfully connected to MySQL Database")
	DB.Conn = dbConn
}

func (db *DatabaseConn) SaveBoard(size, numOfJumpers int, snakes, ladders []byte) int {
	res, err := db.Conn.Exec("INSERT INTO boards (size, num_of_jumpers, snakes, ladders) VALUES (?, ?, ?, ?)",
		size, numOfJumpers, snakes, ladders)
	if err != nil {
		log.Fatal("ERROR in inserting into boards table: ", err.Error())
	}
	boardID, err := res.LastInsertId()
	if err != nil {
		log.Fatal("ERROR in getting inserted board id: ", err.Error())
	}
	return int(boardID)
}

func (db *DatabaseConn) CheckPlayer(playerName string) (int, bool) {
	var playerID int
	err := db.Conn.QueryRow("SELECT ID FROM players WHERE name=?", playerName).Scan(&playerID)
	if err != nil {
		fmt.Println("ERROR in scanning data into id: ", err.Error())
		return playerID, false
	}
	return playerID, true
}

func (db *DatabaseConn) SavePlayer(playerName string) int {
	res, err := db.Conn.Exec("INSERT INTO players (name) VALUES (?)", playerName)
	if err != nil {
		log.Fatal("ERROR in saving player: ", err.Error())
	}
	playerID, err := res.LastInsertId()
	if err != nil {
		log.Fatal("ERROR in getting inserted player id: ", err.Error())
	}
	return int(playerID)
}

func (db *DatabaseConn) SaveGameScreen(numOfDices, boardID int, playerIDs []int) int {
	playerIDsJSON, err := json.Marshal(playerIDs)
	if err != nil {
		log.Fatal("ERROR in marshalling: ", playerIDs, err.Error())
	}
	res, err := db.Conn.Exec("INSERT INTO game_screen (num_of_dices, players_in_game, board_id) VALUES (?, ?, ?)",
		numOfDices, playerIDsJSON, boardID)
	if err != nil {
		log.Fatal("ERROR in saving game screen: ", err.Error())
	}
	screenID, err := res.LastInsertId()
	if err != nil {
		log.Fatal("ERROR in getting inserted screen id: ", err.Error())
	}
	return int(screenID)
}

func (db *DatabaseConn) MarkGameWinner(screenID, winnerID int) {
	_, err := db.Conn.Exec("UPDATE game_screen SET winner_id=? WHERE ID=?", winnerID, screenID)
	if err != nil {
		log.Fatal("ERROR in updating winner of game: ", err.Error())
	}
}

func (db *DatabaseConn) SaveGameMove(gameId, turnNo, playerId, diceRoll, start, end int, jumped bool) {
	_, err := db.Conn.Exec("INSERT INTO game_history (game_id, turn_no, player_id, dice_roll, start, end, jumped) VALUES (?, ?, ?, ?, ?, ?, ?)",
		gameId, turnNo, playerId, diceRoll, start, end, jumped)
	if err != nil {
		log.Fatal("ERROR in saving game move: ", err.Error())
	}
}
