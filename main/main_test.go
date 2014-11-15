package main_test

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/elkorn/code_retreat-game_of_life/main"


func TestMakeCell(t *testing.T) {
    c := &main.Cell{
        X: 0,
        Y: 0,
    }

    assert.NotNil(t, c)
    assert.Equal(t, 0, c.X)
    assert.Equal(t, 0, c.Y)
}

func TestHaveBoard(t *testing.T) {
    seed := [][] int {
        []int{0, 1, 0},
        []int{0, 1, 1},
        []int{1, 1, 0},
    }

    board := main.MakeBoard(seed)
    assert.NotNil(t, board)
    assert.Equal(t, len(seed), len(*board))
    for i := 0; i < len(*board); i++ {
        assert.Equal(t, len(seed[i]), len((*board)[i]))
    }

}

func TestBoardSeed(t *testing.T) {
    seed := [][] int {
        []int{0, 1, 0},
        []int{0, 1, 1},
        []int{1, 1, 0},
    }

    board := main.MakeBoard(seed)

    assert.NotNil(t, (*board)[0][1])
    assert.NotNil(t, (*board)[1][2])
    assert.NotNil(t, (*board)[2][0])
}