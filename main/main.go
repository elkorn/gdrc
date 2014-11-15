package main

type Board struct {
	leBoard map[int]map[int]bool
}

func MakeBoard(seed [][]int) Board {
	result := Board{
		leBoard: make(map[int]map[int]bool),
	}

	for i := 0; i < len(seed); i++ {
		result.leBoard[i] = make(map[int]bool)
		for j := 0; j < len(seed[i]); j++ {
			result.leBoard[i][j] = bool(seed[i][j] == 1)
		}
	}
	return result
}

func (self Board) Size() int {
	return len(self.leBoard)
}

func (self Board) RowSize(row int) int {
	return len(self.leBoard[row])
}

func (self Board) MakeAlive(x, y int) {
	if self.leBoard[x] == nil {
		self.leBoard[x] = make(map[int]bool)
		self.leBoard[x][y] = bool(true)
	}
}

func Next(input Board) (result Board) {
	result = Board{
		leBoard: make(map[int]map[int]bool),
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
		// if theX < 0 || theX >= len(self.leBoard[theX]) {
		// 	continue
		// }

		for yi := -1; yi <= 1; yi++ {
			theY := y + yi
			// if theY < 0 || theY >= len(self.leBoard[theX]) {
			// 	continue
			// }

			if yi == 0 && xi == 0 {
				continue
			}

			if self.Get(theX, theY) {
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

func (self Board) Get(x, y int) bool {
	return self.leBoard[x][y]
}
