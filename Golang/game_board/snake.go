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

func NewSnakeJumper(start, end int) Jumper {
	return Jumper(NewSnake(start, end))
}

type SnakeRandInt struct {
}

var snakeRandInt JumperRandInt = &SnakeRandInt{}

func (snakeRandInt *SnakeRandInt) RandIntStart(board *Board) int {
	return 3 + rand.Intn(board.Size-3)
}

func (snakeRandInt *SnakeRandInt) RandIntEnd(_ *Board, start int) int {
	return 2 + rand.Intn(start-2)
}

func PutSnakes(board *Board) {
	PutJumpers(board, snakeRandInt, NewSnakeJumper)
}
