package main

import "fmt"

func main() {
	fmt.Printf("%v\n", Pyramid(0))
	fmt.Printf("%v\n", Pyramid(1))
	fmt.Printf("%v\n", Pyramid(2))
}

func Pyramid(n int) [][]int {
	pyramid := make([][]int, n)
	//fmt.Printf("%v\n", pyramid)
	for i := 0; i < n; i++ {
		floor := make([]int, i+1)
		for j := 0; j < len(floor); j++ {
			floor[j] = 1
		}
		pyramid[i] = floor
	}
	return pyramid
}
