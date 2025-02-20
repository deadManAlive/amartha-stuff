package logic01

import (
	"fmt"
	"math"
)

var Functions = [...]func(int) []int{
	func(n int) []int {
		arr := make([]int, 0, n)

		for i := range n {
			arr = append(arr, 2*i+1)
		}

		return arr
	},
	func(n int) []int {
		arr := make([]int, 0, n)

		for i := range n {
			arr = append(arr, 2*(i+1))
		}

		return arr
	},
	func(n int) []int {
		arr := make([]int, 0, n)

		for i := range n {
			arr = append(arr, 3*(i+1))
		}

		return arr
	},
	func(n int) []int {
		arr := make([]int, 0, n)

		for i := range n {
			arr = append(arr, 19-(2*i))
		}

		return arr
	},
	func(n int) []int {
		arr := make([]int, 0, n)

		for i := range n {
			arr = append(arr, 20-(2*i))
		}

		return arr
	},
	func(n int) []int {
		arr := make([]int, 0, n)

		for i := range n {
			arr = append(arr, 30-(3*i))
		}

		return arr
	},
	func(n int) []int {
		arr := make([]int, 0, n)
		j := 1

		for i := range n {
			arr = append(arr, j)
			if float32(i+1) < float32(n)/2 {
				j += 2
			} else if float32(i+1) > float32(n)/2 {
				j -= 2
			}
		}

		return arr
	},
	func(n int) []int {
		arr := make([]int, 0, n)
		j := 2

		for i := range n {
			arr = append(arr, j)
			if float32(i+1) < float32(n)/2 {
				j += 2
			} else if float32(i+1) > float32(n)/2 {
				j -= 2
			}
		}

		return arr
	},
	func(n int) []int {
		arr := make([]int, 0, n)
		j := 3

		for i := range n {
			arr = append(arr, j)
			if float32(i+1) < float32(n)/2 {
				j += 3
			} else if float32(i+1) > float32(n)/2 {
				j -= 3
			}
		}

		return arr
	},
	func(int) []int {
		return []int{}
	},
	func(int) []int {
		return []int{}
	},
	func(n int) []int {
		arr := make([]int, 0, n)

		for i := range n {
			arr = append(arr, 1+2*(i%4))
		}

		return arr
	},
}

var FizBuzz = [2]func(int) []string{
	func(n int) []string {
		arr := make([]string, 0, n)
		j := 2

		for i := range n {
			if i%2 == 1 {
				arr = append(arr, "fizz")
			} else {
				arr = append(arr, fmt.Sprintf("%d", j))
				j *= 2
			}
		}

		return arr
	},
	func(n int) []string {
		arr := make([]string, 0, n)
		j := -1

		for i := range n {
			if i%2 == 1 {
				arr = append(arr, fmt.Sprintf("%d", int(3*math.Pow(2, float64(j)))))
				j += 1

			} else {
				arr = append(arr, "buzz")
			}
		}

		return arr
	},
}
