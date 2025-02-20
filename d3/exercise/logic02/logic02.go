package logic02

import utils "github.com/deadManAlive/amartha-utils"

var Functions = [...]func(int) [][]int{
	func(n int) [][]int {
		mat := utils.SqEmpMatGen(n)
		for i := range n {
			for j := range n {
				mat[i][j] = 2*j + 1
			}
		}
		return mat
	},
	func(n int) [][]int {
		mat := utils.SqEmpMatGen(n)
		for i := range n {
			for j := range n {
				mat[i][j] = 2 * (j + 1)
			}
		}
		return mat
	},
	func(n int) [][]int {
		mat := utils.SqEmpMatGen(n)
		for i := range n {
			for j := range n {
				mat[i][j] = 1 + (18 * i) + (2 * j)
			}
		}
		return mat
	},
	func(n int) [][]int {
		mat := utils.SqEmpMatGen(n)

		for i := range n {
			for j := range n {
				if i == 0 {
					mat[i][j] = 1 + (3 * j)
				} else {
					if j == 1 {
						mat[i][j] = 8 + (20 * i) + (3 * j)
					} else {
						mat[i][j] = 9 + (20 * i) + (2 * j)
					}
				}
			}
		}
		return mat
	},
	func(n int) [][]int {
		mat := utils.SqEmpMatGen(n)
		for i := range n {
			for j := range n {
				if i%2 == 0 {
					mat[i][j] = 1 + (18 * i) + (2 * j)
				} else {
					mat[i][j] = (18 * (i + 1)) - 1 - (2 * j)
				}
			}
		}
		return mat
	},
	func(n int) [][]int {
		mat := utils.SqEmpMatGen(n)
		s := 1
		for i := range n {
			for j := range n {
				if i%2 == 0 {
					mat[i][j] = s
					s += 3
				} else {
					mat[i][n-1-j] = s
					s += 2
				}
			}
		}
		return mat
	},
}
