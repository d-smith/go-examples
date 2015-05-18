package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func rollMany(rolls int, pins int, game *Game) {
	for i := 0; i < rolls; i++ {
		game.roll(pins)
	}
}

//1st Test - roll all gutter balls
func TestGutterGame(t *testing.T) {
	game := NewGame()
	t.Log("All gutter balls scores 0")
	rollMany(20, 0, game)
	assert.Equal(t, 0, game.score())
}

//2nd Test - roll all ones
func TestAllOnes(t *testing.T) {
	game := NewGame()
	t.Log("1 pins per roll scores 20")
	rollMany(20, 1, game)
	assert.Equal(t, 20, game.score())
}

//3rd Test - roll one spare
func TestRollOneSpace(t *testing.T) {
	game := NewGame()
	t.Log("One spare, a three, then gutters scores 16")
	game.roll(5)
	game.roll(5)
	game.roll(3)
	rollMany(17, 0, game)
	assert.Equal(t, 16, game.score())
}

//3rd Test - roll one strike
func TestRollOneStrike(t *testing.T) {
	game := NewGame()
	t.Log("One strike, a 3, a 4, then zeros scores 24")
	game.roll(10)
	game.roll(3)
	game.roll(4)
	rollMany(16, 0, game)
	assert.Equal(t, 24, game.score())
}

//4th Test - roll all strikes
func TestPerfectGame(t *testing.T) {
	game := NewGame()
	t.Log("Perfect game scores 300")
	rollMany(12, 10, game)
	assert.Equal(t, 300, game.score())
}
