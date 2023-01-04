package players

import "fmt"

type Player struct {
	ID   int
	Name string
}

func (player *Player) Print() {
	fmt.Printf("Player: %v, ID:%v\n", player.Name, player.ID)
}

func (player *Player) PrintMove(start, end, diceRoll int, jumped bool) {
	if jumped {
		if start < end {
			fmt.Printf("%v Got ladder at %v and moved from %v to %v\n", player.Name, start+diceRoll, start, end)
		} else {
			fmt.Printf("%v Was snake bitten at %v and moved from %v to %v\n", player.Name, start+diceRoll, start, end)
		}
	} else {
		if start != end {
			fmt.Printf("%v moved from %v to %v\n", player.Name, start, end)
		} else {
			fmt.Printf("%v Got %v on Dice Roll. But can't jump from %v to %v\n", player.Name, diceRoll, start, start+diceRoll)
		}
	}
}

func (player *Player) DeclareWin() {
	fmt.Printf("%v Won the Game. Congratulation!\n", player.Name)
}

func SetPlayersOnStart(players []*Player) map[int]int {
	playerPositions := make(map[int]int)
	for _, player := range players {
		playerPositions[player.ID] = 1
	}
	return playerPositions
}
