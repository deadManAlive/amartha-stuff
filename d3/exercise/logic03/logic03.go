package logic03

import utils "github.com/deadManAlive/amartha-utils"

var Functions = [...]func(int) [][]int{
	// 1
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		s := 1

		for i := range n {
			for j := range n {
				if i%2 == 0 && i >= j {
					mat[i][j] = s
					s += 2
				} else if i%2 == 1 && n-i-1 <= j {
					mat[i][n-j-1] = s
					s += 2
				}
			}
		}

		return mat
	},
	// 2
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		s := 1

		for i := range n {
			for j := range n {
				if i%2 == 0 && i <= j {
					mat[i][j] = s
					s += 2
				} else if i%2 == 1 && n-i-1 >= j {
					mat[i][n-j-1] = s
					s += 2
				}
			}
		}

		return mat
	},
	// 3
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		s := 2

		for i := range n {
			for j := range n {
				if i%2 == 0 && i+j <= n-1 {
					mat[i][j] = s
					if i == 0 && j > 0 || i+j == n-1 {
						s += 3
					} else {
						s += 2
					}
				} else if i%2 == 1 && i <= j {
					mat[i][n-j-1] = s
					if j == n-1 {
						s += 2
					} else {
						s += 3
					}
				}
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
				if i%2 == 1 && i >= j {
					mat[i][n-1-i+j] = s
					s += 2
				} else if i%2 == 0 && n-i-1 <= j {
					mat[i][n-1-i+(n-j-1)] = s
					s += 2
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
				if i <= (n-1)/2 {
					if i%2 == 0 && i >= j && i+j <= n-1 {
						mat[i][j] = s
						mat[n-1-i][j] = s
						mat[i][n-1-j] = s
						mat[n-1-i][n-1-j] = s
						s += 2
					} else if i%2 == 1 && n-i-1 <= j && i+n-j-1 <= n-1 {
						mat[i][n-j-1] = s
						mat[n-1-i][n-j-1] = s
						mat[i][n-1-(n-1-j)] = s
						mat[n-1-i][n-1-(n-1-j)] = s
						s += 2
					}
				}
			}
		}

		return mat
	},
	// 6
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		s := 1

		for i := range n {
			for j := range n {
				if i <= (n-1)/2 {
					if i%2 == 0 && i <= j && i+j <= n-1 {
						mat[i][j] = s
						mat[n-1-i][j] = s
						s += 2
					} else if i%2 == 1 && n-i-1 >= j && i+n-j-1 <= n-1 {
						mat[i][n-j-1] = s
						mat[n-1-i][n-j-1] = s
						s += 2
					}
				}
			}
		}

		return mat
	},
	// 7
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		s := 1

		for i := range n {
			for j := range n {
				if i <= (n-1)/2 {
					if i%2 == 0 && i+j >= (n-1)/2 && i >= j-n/2 {
						mat[i][n-j-1] = s
						mat[n-1-i][n-j-1] = s
						s += 2
					} else if i%2 == 1 && i+j >= (n-1)/2 && i >= j-n/2 {
						mat[i][j] = s
						mat[n-1-i][j] = s
						s += 2
					}
				}
			}
		}

		return mat
	},
	// 8
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		s := 1

		for i := range n {
			for j := range n {
				if i <= (n-1)/2 {
					if i%2 == 0 && i+j >= (n-1)/2 && i >= j-n/2 {
						mat[n-j-1][i] = s
						mat[n-j-1][n-1-i] = s
						s += 2
					} else if i%2 == 1 && i+j >= (n-1)/2 && i >= j-n/2 {
						mat[j][i] = s
						mat[j][n-1-i] = s
						s += 2
					}
				}
			}
		}

		return mat
	},
	// 9
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
	// 10
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		for i := range n {
			s := 9
			for j := range n {
				if i <= (n-1)/2 {
					if i+j >= (n-1)/2 && i >= j-n/2 {
						mat[i][j] = s
						mat[n-1-i][j] = s
						if j >= n/2 {
							s += 2
						} else {
							s -= 2
						}
					}
				}
			}
		}

		return mat
	},
	// 11
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		s := 1

		for i := range n {
			for j := range n {
				if i <= (n-1)/2 {
					if i%2 == 0 && i >= j && i+j <= n-1 {
						mat[i][n-1-j] = s
						mat[n-1-i][n-1-j] = s
						s += 2
					} else if i%2 == 1 && n-i-1 <= j && i+n-j-1 <= n-1 {
						mat[i][n-1-(n-1-j)] = s
						mat[n-1-i][n-1-(n-1-j)] = s
						s += 2
					}
				}
				if (i == j) && j < n/2 {
					mat[i][j] = 2*(i+1) - 1
					mat[n-1-i][j] = 2*(i+1) - 1
				}
			}
		}

		return mat
	},
	// 12
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		s := 1

		for i := range n {
			for j := range n {
				if i <= (n-1)/2 {
					if i%2 == 0 && i >= j && i+j <= n-1 {
						mat[i][j] = s
						mat[n-1-i][j] = s
						s += 2
					} else if i%2 == 1 && n-i-1 <= j && i+n-j-1 <= n-1 {
						mat[i][n-j-1] = s
						mat[n-1-i][n-j-1] = s
						s += 2
					}
				}
				if (i == j) && j >= n/2 {
					mat[i][j] = 2*((n-1-i)+1) - 1
					mat[n-1-i][j] = 2*((n-1-i)+1) - 1
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
				if i <= (n-1)/2 {
					if i+j >= (n-1)/2 && i >= j-n/2 {
						if (i+j)%2 == 0 {
							mat[j][i] = 2*i + 1
							mat[j][n-1-i] = 2*i + 1
						}
					}
				}
			}
		}

		return mat
	},
	// 14
	func(n int) [][]int {
		mat := utils.EmptySquareMatrixGenerator(n)

		return mat
	},
}
