package game_screen

import (
	"math"
	"math/rand"
)

type Dice struct {
}

func (dice *Dice) Roll() int {
	return 1 + rand.Intn(6)
}

func GetDices(numOfPlayers int) Dices {
	numOfDices := int(math.Floor(math.Sqrt(float64(numOfPlayers))))
	dices := make([]*Dice, numOfDices)
	for index := 0; index < numOfDices; index++ {
		dices[index] = &Dice{}
	}
	return dices
}

type Dices []*Dice

func (dices Dices) Roll() int {
	diceRoll := 0
	for _, dice := range dices {
		diceRoll += dice.Roll()
	}
	return diceRoll
}
