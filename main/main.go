package main

type Cell struct {
	X, Y int
}

type Board [][]*Cell

func MakeBoard(seed [][]int) *Board {
	result := make(Board, len(seed))
	for i := 0; i < len(seed); i++ {
		result[i] = make([]*Cell, len(seed[i]))
		for j := 0; j < len(seed[i]); j++ {
			if seed[i][j] == 1 {
				result[i][j] = &Cell{
					X: i,
					Y: j,
				}
			}

		}
	}
	return &result
}
