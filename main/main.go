package main

type Cell bool

type Board struct {
	leBoard map[int]map[int]Cell
}

func MakeBoard(seed [][]int) Board {
	result := Board{
		leBoard: make(map[int]map[int]Cell),
	}

	for i := 0; i < len(seed); i++ {
		result.leBoard[i] = make(map[int]Cell)
		for j := 0; j < len(seed[i]); j++ {
			result.leBoard[i][j] = Cell(seed[i][j] == 1)
		}
	}
	return result
}

func (self Board) MakeAlive(x, y int) {
	if self.leBoard[x] == nil {
		self.leBoard[x] = make(map[int]Cell)
		self.leBoard[x][y] = Cell(true)
	}
}

func Next(input Board) (result Board) {
	result = Board{
		leBoard: make(map[int]map[int]Cell),
	}

	for i := 0; i < len(input.leBoard); i++ {
		result.leBoard[i] = input.leBoard[i]
	}

	return input
}

func (self Board) countAliveNeighbors(x, y int) int {
	alive := 0
	for xi := -1; xi <= 1; xi++ {
		theX := x + xi
		if theX < 0 || theX >= len(self.leBoard[theX]) {
			continue
		}

		for yi := -1; yi <= 1; yi++ {
			theY := y + yi
			if theY < 0 || theY >= len(self.leBoard[theX]) {
				continue
			}

			if yi == 0 && xi == 0 {
				continue
			}

			if self.leBoard[theX][theY] {
				alive++
			}
		}
	}

	return alive
}

func (self Board) ShouldDie(x, y int) bool {
	aliveNeighbors := self.countAliveNeighbors(x, y)
	if self.leBoard[x][y] { // self is alive
		return aliveNeighbors < 2 || aliveNeighbors > 3
	} else {
		return aliveNeighbors != 3
	}
}

func (self Board) Get(x, y int) Cell {
	return self.leBoard[x][y]
}
