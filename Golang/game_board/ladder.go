package game_board

import (
	"fmt"
	"math/rand"
)

type Ladder struct {
	Start int
	End   int
}

var _ Jumper = (*Ladder)(nil)

func (ladder *Ladder) Print() {
	fmt.Printf("Ladder: From %v to %v\n", ladder.Start, ladder.End)
}

func (ladder *Ladder) Jump() int {
	return ladder.End
}

func NewLadder(start, end int) *Ladder {
	if end <= start {
		return nil
	}
	return &Ladder{start, end}
}

func NewLadderJumper(start, end int) Jumper {
	return Jumper(NewLadder(start, end))
}

type LadderRandInt struct {
}

var ladderRandInt JumperRandInt = &LadderRandInt{}

func (ladderRandInt *LadderRandInt) RandIntStart(board *Board) int {
	return 2 + rand.Intn(board.Size-2)
}

func (ladderRandInt *LadderRandInt) RandIntEnd(board *Board, start int) int {
	return start + 1 + rand.Intn(board.Size-start-1)
}

func PutLadders(board *Board) {
	PutJumpers(board, ladderRandInt, NewLadderJumper)
}
