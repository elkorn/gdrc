package main_test

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/elkorn/code_retreat-game_of_life/main"

var seed [][]int = [][]int{
	[]int{0, 1, 0},
	[]int{0, 1, 1},
	[]int{1, 1, 0},
}

func TestMakeCell(t *testing.T) {
	c := main.Cell(true)

	assert.True(t, bool(c))
}

func TestHaveBoard(t *testing.T) {
	board := main.MakeBoard(seed)
	assert.NotNil(t, board)
	assert.Equal(t, len(seed), len(board))
	for i := 0; i < len(board); i++ {
		assert.Equal(t, len(seed[i]), len(board[i]))
	}
}

func TestBoardSeed(t *testing.T) {
	board := main.MakeBoard(seed)

	assert.True(t, bool(board[0][1]))
	assert.True(t, bool(board[1][2]))
	assert.True(t, bool(board[2][0]))
}

func TestNext(t *testing.T) {
	board1 := main.MakeBoard(seed)
	next := main.Next(board1)
	assert.NotNil(t, next)
	assert.Equal(t, len(board1), len(next))
	for i := 0; i < len(next); i++ {
		assert.Equal(t, len(board1[i]), len(next[i]))
	}
}

func defaultBoard() (board main.Board) {
	board = main.MakeBoard(seed)
	return
}

func TestShouldDieWhenLessThan2(t *testing.T) {
	seed := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 1},
		[]int{1, 1, 0},
	}

	board := main.MakeBoard(seed)
	assert.True(t, board.ShouldDie(0, 0))
	seed = [][]int{
		[]int{0, 1, 0},
		[]int{0, 1, 1},
		[]int{1, 1, 0},
	}

	board = main.MakeBoard(seed)
	assert.False(t, board.ShouldDie(0, 0))
}

func TestShouldDieWhenMoreThan3(t *testing.T) {
	seed := [][]int{
		[]int{1, 1, 1},
		[]int{1, 1, 1},
		[]int{1, 1, 1},
	}

	board := main.MakeBoard(seed)
	assert.True(t, board.ShouldDie(0, 1))
}

func TestShouldLiveWhenLiveAnd2Or3(t *testing.T) {
	seed := [][]int{
		[]int{1, 1, 0},
		[]int{0, 1, 1},
		[]int{0, 1, 0},
	}

	board := main.MakeBoard(seed)
	assert.False(t, board.ShouldDie(0, 0))
	assert.False(t, board.ShouldDie(1, 2))
}
