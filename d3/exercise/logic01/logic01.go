package logic01

import "fmt"

func Hello() {
	fmt.Println("Hello!")
}

func Problem01(n int) []int {
	arr := make([]int, n)

	for i := range n {
		arr = append(arr, 2*(i+1)+1)
	}

	return arr
}
