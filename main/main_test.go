package main_test

import (
	"testing"

	"github.com/elkorn/code_retreat-game_of_life/main"
	"github.com/stretchr/testify/assert"
)

var seed [][]int = [][]int{
	[]int{0, 1, 0},
	[]int{0, 1, 1},
	[]int{1, 1, 0},
}

func TestMakebool(t *testing.T) {
	c := true

	assert.True(t, c)
}

func TestHaveBoard(t *testing.T) {
	board := main.MakeBoard(seed)

	assert.NotNil(t, board)
	assert.Equal(t, len(seed), board.Size())
	for i := 0; i < board.Size(); i++ {
		assert.Equal(t, len(seed[i]), board.RowSize(i))
	}
}

func TestBoardSeed(t *testing.T) {
	board := main.MakeBoard(seed)

	assert.True(t, bool(board.Get(0, 1)))
	assert.True(t, bool(board.Get(1, 2)))
	assert.True(t, bool(board.Get(2, 0)))
}

func TestNext(t *testing.T) {
	board1 := main.MakeBoard(seed)
	next := board1.Next()
	assert.NotNil(t, next)
	assert.Equal(t, board1.Size(), next.Size())
	for i := 0; i < next.Size(); i++ {
		assert.Equal(t, board1.RowSize(i), next.RowSize(i))
	}
}

func defaultBoard() (board main.Board) {
	board = main.MakeBoard(seed)
	return
}

func TestShouldDieWhenLessThan2(t *testing.T) {
	seed := [][]int{
		[]int{1, 0, 0},
		[]int{0, 1, 1},
		[]int{1, 1, 0},
	}

	board := main.MakeBoard(seed)
	assert.True(t, board.ShouldDie(0, 0))
}

func TestShouldNotDieWhenAliveAnd2LiveNeighbors(t *testing.T) {
	seed := [][]int{
		[]int{1, 1, 0},
		[]int{0, 1, 0},
		[]int{1, 1, 0},
	}

	board := main.MakeBoard(seed)
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

func TestShouldComeAliveWhenDeadWith3AliveNeighbors(t *testing.T) {
	seed := [][]int{
		[]int{1, 1, 0},
		[]int{0, 0, 1},
		[]int{0, 0, 0},
	}

	board := main.MakeBoard(seed)
	assert.False(t, board.ShouldDie(1, 1))
}

func TestSpaceIsExpandable(t *testing.T) {
	board := main.MakeBoard([][]int{
		[]int{
			1,
		},
	})

	assert.True(t, board.ShouldDie(0, 0))

	board.MakeAlive(100, 100)

	assert.True(t, board.ShouldDie(0, 0))
	assert.True(t, board.ShouldDie(100, 100))
}

func TestNeighborhoodsInExpandedSpaceAreCohesive(t *testing.T) {
	board := main.MakeBoard([][]int{
		[]int{
			1,
		},
	})

	assert.True(t, board.ShouldDie(0, 0))

	board.MakeAlive(100, 100)
	board.MakeAlive(99, 100)
	board.MakeAlive(101, 100)

	assert.False(t, board.ShouldDie(100, 100))
}

func TestNextIterationProgressesPopulation(t *testing.T) {
	board := defaultBoard()
	next := board.Next()

	assert.False(t, next.Get(1, 1))
}
