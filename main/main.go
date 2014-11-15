package main

type Cell bool

type Board [][]Cell

func MakeBoard(seed [][]int) Board {
	result := make(Board, len(seed))
	for i := 0; i < len(seed); i++ {
		result[i] = make([]Cell, len(seed[i]))
		for j := 0; j < len(seed[i]); j++ {
			if seed[i][j] == 1 {
				result[i][j] = Cell(true)
			}

		}
	}
	return result
}

func Next(input Board) (result Board) {
	result = make(Board, len(input))
	for i := 0; i < len(input); i++ {
		copy(result[i], input[i])
	}

	return input
}

func (self Board) countAliveNeighbors(x, y int) int {
	alive := 0
	for xi := -1; xi <= 1; xi++ {
		theX := x + xi
		if theX < 0 || theX >= len(self[theX]) {
			continue
		}

		for yi := -1; yi <= 1; yi++ {
			theY := y + yi
			if theY < 0 || theY >= len(self[theX]) {
				continue
			}

			if yi == 0 && xi == 0 {
				continue
			}

			if self[theX][theY] {
				alive++
			}
		}
	}

	return alive
}

func (self Board) ShouldDie(x, y int) bool {
	aliveNeighbors := self.countAliveNeighbors(x, y)
	if self[x][y] { // self is alive
		return aliveNeighbors < 2 || aliveNeighbors > 3
	} else {
		return false
	}

}
