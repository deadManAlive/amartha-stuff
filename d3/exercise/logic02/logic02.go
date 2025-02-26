package logic02

import utils "github.com/deadManAlive/amartha-utils"

var Functions = [...]func(int) [][]int{
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)
		for i := range n {
			for j := range n {
				mat[i][j] = 2*j + 1
			}
		}
		return mat
	},
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)
		for i := range n {
			for j := range n {
				mat[i][j] = 2 * (j + 1)
			}
		}
		return mat
	},
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)
		for i := range n {
			for j := range n {
				mat[i][j] = 1 + (18 * i) + (2 * j)
			}
		}
		return mat
	},
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

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
		mat := utils.EmptySquareMatrixGenerator(n)
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
		mat := utils.EmptySquareMatrixGenerator(n)

		e := 1

		for i := range n {
			for j := range n {
				if i%2 == 0 {
					mat[i][j] = e
				} else {
					mat[i][n-1-j] = e
				}

				if (i == 0) || (i == 1 && j == n-1) || (i == 2 && j == 6) || (i == 3 && j == 0) || (i == 4 && j == 0) || (i == 6 && j == 7 || j == 8) {
					e += 3
				} else {
					e += 2
				}
				if e == 107 || e == 125 {
					e -= 1
				}
			}
		}

		return mat
	},
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		for i := range n {
			for j := range n {
				if i == j {
					mat[i][j] = 2*i + 1
				}
			}
		}

		return mat
	},
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		for i := range n {
			for j := range n {
				if i+j == n-1 {
					mat[i][j] = (2*n - 1) - 2*i
				}
			}
		}

		return mat
	},
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		for i := range n {
			for j := range n {
				if i+j == n-1 {
					mat[i][j] = (2*n - 1) - 2*i
				} else if i == j {
					mat[i][j] = 2*i + 1
				}
			}
		}

		return mat
	},
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		for i := range n {
			for j := range n {
				if i >= j {
					mat[i][j] = 2*i + 1
				}
			}
		}

		return mat
	},
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		for i := range n {
			for j := range n {
				if i <= j {
					mat[i][j] = 2*i + 1
				}
			}
		}

		return mat
	},
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		for i := range n {
			for j := range n {
				if i >= j && i+j <= n-1 || i <= j && i+j >= n-1 {
					mat[i][j] = 2*j + 1
				}
			}
		}

		return mat
	},
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		for i := range n {
			for j := range n {
				if i <= j && i+j <= n-1 || i >= j && i+j >= n-1 {
					mat[i][j] = 2*j + 1
				}
			}
		}

		return mat
	},
}
