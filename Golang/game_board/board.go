package game_board

import "fmt"

type Board struct {
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
