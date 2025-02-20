package main

import (
	"fmt"
)

func cellDrawer(cellStatus bool) string {
	if cellStatus {
		return "x"
	} else {
		return "."
	}
}

func main() {
	col := 3
	row := 3
	iter := 5

	cell := make([][]bool, col)
	for i := range col {
		cell[i] = make([]bool, row)
	}

	ncell := make([][]int, col)
	for i := range col {
		ncell[i] = make([]int, row)
	}

	dir := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	prem := [][]int{{0, 1}, {1, 0}, {2, 1}}

	for i := range prem {
		px := prem[i][0]
		py := prem[i][1]

		cell[px][py] = true
	}

	for n := range iter {
		fmt.Printf("Iteration: %d\n", n)
		for i := range col {
			for j := range row {
				fmt.Printf("%s ", cellDrawer(cell[i][j]))
			}
			fmt.Println()
		}
		for i := range col {
			for j := range row {
				neighbors := 0

				for _, n := range dir {
					ni, nj := i+n[0], j+n[1]
					if ni >= 0 && nj >= 0 && ni < row && nj < col && cell[ni][nj] {
						neighbors++
					}
				}

				ncell[i][j] = neighbors

			}
		}

		for i := range col {
			for j := range row {
				if cell[i][j] {
					if ncell[i][j] < 2 {
						cell[i][j] = false
					} else if ncell[i][j] > 3 {
						cell[i][j] = false
					}
				} else {
					if ncell[i][j] == 3 {
						cell[i][j] = true
					}
				}
			}
		}
	}
}
