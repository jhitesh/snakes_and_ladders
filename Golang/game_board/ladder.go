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

type LadderUtils struct {
}

var ladderRandInt JumperUtils = &LadderUtils{}

func (ladderRandInt *LadderUtils) RandIntStart(board *Board) int {
	return 2 + rand.Intn(board.Size-2)
}

func (ladderRandInt *LadderUtils) RandIntEnd(board *Board, start int) int {
	return start + 1 + rand.Intn(board.Size-start-1)
}

func (ladderRandInt *LadderUtils) WithinRange(start, end int, board *Board) bool {
	return start > 1 && start < board.Size-1 && end > start && end < board.Size
}

func (ladderRandInt *LadderUtils) NewJumper(start, end int) Jumper {
	return Jumper(NewLadder(start, end))
}

func PutLadders(board *Board) {
	PutJumpers(board, ladderRandInt)
}
