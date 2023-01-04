package game_board

type Jumper interface {
	Jump() int
	Print()
}

type JumperRandInt interface {
	RandIntStart(*Board) int
	RandIntEnd(*Board, int) int
}

func PutJumpers(board *Board, jumperRandInt JumperRandInt, newJumper func(int, int) Jumper) {
	numOfJumpers := board.Size / 12

	count := 0
	for count < numOfJumpers {
		jumperPositionStart := jumperRandInt.RandIntStart(board)
		if _, ok := board.Jumpers[jumperPositionStart]; !ok {
			jumperPositionEnd := jumperRandInt.RandIntEnd(board, jumperPositionStart)
			jumper := newJumper(jumperPositionStart, jumperPositionEnd)
			board.Jumpers[jumperPositionStart] = jumper
			count++
		}
	}
}
