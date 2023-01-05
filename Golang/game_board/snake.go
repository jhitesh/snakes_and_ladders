package game_board

import (
	"fmt"
	"math/rand"
)

type Snake struct {
	Start int
	End   int
}

var _ Jumper = (*Snake)(nil)

func (snake *Snake) Print() {
	fmt.Printf("Snake: From %v to %v\n", snake.Start, snake.End)
}

func (snake *Snake) Jump() int {
	return snake.End
}

func NewSnake(start, end int) *Snake {
	if start <= end {
		return nil
	}
	return &Snake{start, end}
}

type SnakeUtils struct {
}

var snakeRandInt JumperUtils = &SnakeUtils{}

func (snakeRandInt *SnakeUtils) RandIntStart(board *Board) int {
	return 3 + rand.Intn(board.Size-3)
}

func (snakeRandInt *SnakeUtils) RandIntEnd(_ *Board, start int) int {
	return 2 + rand.Intn(start-2)
}

func (snakeRandInt *SnakeUtils) WithinRange(start, end int, board *Board) bool {
	return start > 2 && start < board.Size && end > 1 && end < start
}

func (snakeRandInt *SnakeUtils) NewJumper(start, end int) Jumper {
	return Jumper(NewSnake(start, end))
}

func PutSnakes(board *Board) {
	PutJumpers(board, snakeRandInt)
}
