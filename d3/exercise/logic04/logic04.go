package logic04

import (
	utils "github.com/deadManAlive/amartha-utils"
)

// array of function to generate pattern for the main matrices
var elementaryPatternGenerator = [...]func(n int) [][]int{
	// 0
	// += 2, alternating, from left-to-right, from logic2-5
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
	// 1
	// from logic3-7
	func(n int) [][]int {
		mat := utils.EmptyMatrixGenerator(n, 2*n-1)

		s := 1
		w := 2*n - 1

		for i := range n {
			for j := range w {
				if i%2 == 0 && i+j >= (w-1)/2 && i >= j-w/2 {
					mat[i][w-j-1] = s
					s += 2
				} else if i%2 == 1 && i+j >= (w-1)/2 && i >= j-w/2 {
					mat[i][j] = s
					s += 2
				}
			}
		}

		return mat
	},
	// 2
	// also from logic3-7
	func(n int) [][]int {
		mat := utils.EmptyMatrixGenerator(n, 2*n-1)

		s := 1
		w := 2*n - 1

		for i := range n {
			for j := range w {
				if i%2 == 0 && i+j >= (w-1)/2 && i >= j-w/2 {
					mat[n-1-i][w-j-1] = s
					s += 2
				} else if i%2 == 1 && i+j >= (w-1)/2 && i >= j-w/2 {
					mat[n-1-i][j] = s
					s += 2
				}
			}
		}

		return mat
	},
	// 3
	// another one from logic3-7
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		s := 1

		for i := range n {
			for j := range n {
				if i <= (n-1)/2 {
					if i%2 == 0 && i+j >= (n-1)/2 && i >= j-n/2 {
						mat[i][n-j-1] = s
						mat[n-1-i][n-j-1] = n*n + 1 - s
						s += 2
					} else if i%2 == 1 && i+j >= (n-1)/2 && i >= j-n/2 {
						mat[i][j] = s
						mat[n-1-i][j] = n*n + 1 - s
						s += 2
					}
				}
			}
		}

		return mat
	},
	// 4
	// from logic3-9
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		for i := range n {
			s := 1
			for j := range n {
				if i <= (n-1)/2 {
					if i+j >= (n-1)/2 && i >= j-n/2 {
						mat[i][j] = s
						mat[n-1-i][j] = s
						if j >= n/2 {
							s -= 2
						} else {
							s += 2
						}
					}
				}
			}
		}

		return mat
	},
}

var Functions = [...]func(n int) [][]int{
	// 1
	func(n int) [][]int {
		pos := func(n int) int { return (2 + (n - 1)) * n / 2 }

		mat := utils.EmptySquareMatrixGenerator(pos(n))

		for i := range n {
			emb := elementaryPatternGenerator[0](i + 1)
			utils.EmbedMatrix(emb, mat, pos(i), pos(i))
		}

		return mat
	},
	// 2
	func(n int) [][]int {
		pos := func(n int) int { return (2 + (n - 1)) * n / 2 }

		mat := utils.EmptySquareMatrixGenerator(pos(n))

		for i := range n {
			emb := elementaryPatternGenerator[0](i + 1)
			utils.EmbedMatrix(emb, mat, pos(n)-pos(i+1), pos(i))
		}

		return mat
	},
	// 3
	func(n int) [][]int {
		size := n * n

		mat := utils.EmptySquareMatrixGenerator(size)

		pos := func(n int) int { return (2 + (n - 1)) * n / 2 }

		for i := range n {
			emb := elementaryPatternGenerator[0](i + 1)
			utils.EmbedMatrix(emb, mat, pos(i), pos(i))
			// prevent redoing the middle pattern
			if i < n-1 {
				utils.EmbedMatrix(emb, mat, size-pos(i+1), pos(i))
				utils.EmbedMatrix(emb, mat, pos(i), size-pos(i+1))
				utils.EmbedMatrix(emb, mat, size-pos(i+1), size-pos(i+1))
			}
		}

		return mat
	},
	// 4?
	// 5
	func(n int) [][]int {
		pos := func(n int) int { return (2 + (n - 1)) * n / 2 }

		mat := utils.EmptyMatrixGenerator(pos(n), n*n)

		for i := range n {
			emb := elementaryPatternGenerator[1](i + 1)
			utils.EmbedMatrix(emb, mat, pos(i), i*i)
		}

		return mat
	},
	// 6
	func(n int) [][]int {
		pos := func(n int) int { return (2 + (n - 1)) * n / 2 }

		mat := utils.EmptyMatrixGenerator(pos(n), n*n)

		for i := range n {
			emb := elementaryPatternGenerator[1](i + 1)
			utils.EmbedMatrix(emb, mat, pos(n)-pos(i+1), i*i)
		}

		return mat
	},
	// 7
	func(n int) [][]int {
		pos := func(n int) int { return (2 + (n - 1)) * n / 2 }

		mat := utils.EmptyMatrixGenerator(pos(n), n*n)

		for i := range n {
			emb := elementaryPatternGenerator[2](i + 1)
			utils.EmbedMatrix(emb, mat, pos(i), i*i)
		}

		return mat
	},
	// 8
	func(n int) [][]int {
		pos := func(n int) int { return (2 + (n - 1)) * n / 2 }

		mat := utils.EmptyMatrixGenerator(pos(n), n*n)

		for i := range n {
			emb := elementaryPatternGenerator[2](i + 1)
			utils.EmbedMatrix(emb, mat, pos(n)-pos(i+1), i*i)
		}

		return mat
	},
	// 8?
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n * n)

		for i := range n {
			emb := elementaryPatternGenerator[3](2*(i+1) - 1)
			utils.EmbedMatrix(emb, mat, i*i, i*i)
		}

		return mat
	},
	// 9
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n * n)

		for i := range n {
			emb := elementaryPatternGenerator[3](2*(i+1) - 1)
			utils.EmbedMatrix(emb, mat, n*n-(i+1)*(i+1), i*i)
		}

		return mat
	},
	// 10
	func(n int) [][]int {
		size := n*n + n - 1

		pos := func(n int) int { return n * (n + 1) / 2 }
		dual := func(n int) int { return (n*n + 3*n - 2) / 2 }

		mat := utils.EmptySquareMatrixGenerator(size)

		for i := range n {
			m := n - 1 - i
			emb := elementaryPatternGenerator[4](2*(m+1) - 1)
			utils.EmbedMatrix(emb, mat, pos(m), pos(m))
			// skip middle pattern redoing
			if m < n-1 {
				utils.EmbedMatrix(emb, mat, size-dual(m+1), pos(m))
				utils.EmbedMatrix(emb, mat, pos(m), size-dual(m+1))
				utils.EmbedMatrix(emb, mat, size-dual(m+1), size-dual(m+1))
			}
		}

		return mat
	},
}
