package game_board

import (
	"encoding/json"
	"fmt"
	"log"
	"snake_and_ladder/MySQL"
)

type Board struct {
	ID      int
	Size    int
	Jumpers map[int]Jumper
}

func (board *Board) Print() {
	fmt.Printf("Board size is %v\n", board.Size)
	for _, jumper := range board.Jumpers {
		jumper.Print()
	}
}

func (board *Board) Move(position, steps int) int {
	newPosition := position + steps
	if jumper, ok := board.Jumpers[newPosition]; ok {
		newPosition = jumper.Jump()
	}
	if newPosition > board.Size {
		return position
	}
	return newPosition
}

func (board *Board) Snakes() []*Snake {
	snakes := make([]*Snake, 0)
	index := 0
	for _, jumper := range board.Jumpers {
		if !jumper.JumpsUp() {
			snakes = append(snakes, &Snake{})
			snakes[index].Start, snakes[index].End = jumper.EndPoints()
			index++
		}
	}
	return snakes
}

func (board *Board) Ladders() []*Ladder {
	ladders := make([]*Ladder, 0)
	index := 0
	for _, jumper := range board.Jumpers {
		if jumper.JumpsUp() {
			ladders = append(ladders, &Ladder{})
			ladders[index].Start, ladders[index].End = jumper.EndPoints()
			index++
		}
	}
	return ladders
}

func (board *Board) SaveBoard() {
	snakes := board.Snakes()
	snakesJSON, err := json.Marshal(snakes)
	if err != nil {
		log.Fatal("ERROR in marshalling snakes: ", err.Error())
	}

	ladders := board.Ladders()
	laddersJSON, err := json.Marshal(ladders)
	if err != nil {
		log.Fatal("ERROR in marshalling ladders: ", err.Error())
	}
	board.ID = MySQL.DB.SaveBoard(board.Size, len(board.Jumpers), snakesJSON, laddersJSON)
}

func NewBoard(size int) *Board {
	if size < 5 {
		return nil
	}
	board := &Board{
		Size:    size,
		Jumpers: make(map[int]Jumper),
	}

	PutSnakes(board)
	PutLadders(board)
	return board
}
