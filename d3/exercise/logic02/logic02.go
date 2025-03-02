package logic02

import utils "github.com/deadManAlive/amartha-utils"

var Functions = [...]func(int) [][]int{
	// 1
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)
		for i := range n {
			for j := range n {
				mat[i][j] = 2*j + 1
			}
		}
		return mat
	},
	// 2
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)
		for i := range n {
			for j := range n {
				mat[i][j] = 2 * (j + 1)
			}
		}
		return mat
	},
	// 3
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		s := 1

		for i := range n {
			for j := range n {
				mat[i][j] = s
				s += 2
			}
		}
		return mat
	},
	// 4
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		s := 1

		for i := range n {
			for j := range n {
				mat[i][j] = s

				if i == 0 {
					s += 3
				} else {
					if j == 1 || j == n-1 {
						s += 3
					} else {
						s += 2
					}
				}
			}
		}
		return mat
	},
	// 5
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		s := 1

		for i := range n {
			for j := range n {
				if i%2 == 0 {
					mat[i][j] = s
				} else {
					mat[i][n-1-j] = s
				}
				s += 2
			}
		}
		return mat
	},
	// 6
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
	// 7
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
	// 8
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
	// 9
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
	// 10
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
	// 11
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
	// 12
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
	// 13
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
