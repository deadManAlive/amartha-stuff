package main

import (
	"fmt"

	"github.com/deadManAlive/amartha-stuff/d3/exercise/logic01"
	"github.com/deadManAlive/amartha-stuff/d3/exercise/logic02"
	utils "github.com/deadManAlive/amartha-utils"
)

func main() {
	fmt.Println("Logic-01")
	s := 10
	for i, f := range logic01.Functions {

		r := f(s)
		fmt.Printf("%2d: ", i+1)

		if i == 9 || i == 10 {
			r := logic01.FizBuzz[i-9](s)
			for _, v := range r {
				fmt.Printf("%4s ", v)
			}
			fmt.Println()
			continue
		}

		utils.PrintSlice(r)
	}

	fmt.Println("Logic-02")
	s = 9

	for i, f := range logic02.Functions {
		r := f(s)
		fmt.Printf("%2d:\n", i+1)
		utils.PrintSlice2(r)
	}
}
