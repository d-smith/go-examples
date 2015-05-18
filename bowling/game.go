package main

type Game struct {
	currentRoll int
	rolls       [20]int
}

func NewGame() *Game {
	return new(Game)
}

func (g *Game) roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll++
}

func (g *Game) score() int {
	score := 0
	rollIdx := 0

	for frame := 0; frame < 10; frame++ {
		if g.isStrike(rollIdx) {
			score = score + 10 + g.rolls[rollIdx+1] + g.rolls[rollIdx+2]
			rollIdx++
		} else if g.isSpare(rollIdx) {
			score = score + 10 + g.rolls[rollIdx+2]
			rollIdx += 2
		} else {
			score = score + g.rolls[rollIdx] + g.rolls[rollIdx+1]
			rollIdx += 2
		}
	}
	return score
}

func (g *Game) isSpare(rollIdx int) bool {
	return g.rolls[rollIdx]+g.rolls[rollIdx+1] == 10
}

func (g *Game) isStrike(rollIdx int) bool {
	return g.rolls[rollIdx] == 10
}
