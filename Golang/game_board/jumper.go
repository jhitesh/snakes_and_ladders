package game_board

type Jumper interface {
	Jump() int
	Print()
	JumpsUp() bool
	EndPoints() (int, int)
}

type JumperUtils interface {
	RandIntStart(*Board) int
	RandIntEnd(*Board, int) int
	WithinRange(int, int, *Board) bool
	NewJumper(int, int) Jumper
}

func PutJumpers(board *Board, jumperUtils JumperUtils) {
	numOfJumpers := board.Size / 12

	for count := 0; count < numOfJumpers; count++ {
		jumperStart := jumperUtils.RandIntStart(board)
		if _, ok := board.Jumpers[jumperStart]; !ok {
			jumperEnd := jumperUtils.RandIntEnd(board, jumperStart)
			if !jumperUtils.WithinRange(jumperStart, jumperEnd, board) {
				continue
			}
			jumper := jumperUtils.NewJumper(jumperStart, jumperEnd)
			board.Jumpers[jumperStart] = jumper
		}
	}
}
