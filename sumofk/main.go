package main

import (
	"container/list"
)

func main() {
	var ts = []int{50, 55, 56, 57, 58}
	ts = []int{91, 74, 73, 85, 73, 81, 87}
	//ts = []int{50}
	bestSum := ChooseBestSum(331, 5, ts)
	println(bestSum)
}

func ChooseBestSum(t, k int, ls []int) int {
	if len(ls) < k {
		return -1
	}

	sums := sumsOfCombinations(ls, k)
	filteredSums := filterUnderValue(sums, t)
	return bestSum(filteredSums)
}

func bestSum(filteredSums *list.List) int {
	max := 0
	for e := filteredSums.Front(); e != nil; e = e.Next() {
		if e.Value.(int) >= max {
			max = e.Value.(int)
		}
	}

	if max == 0 {
		return -1
	} else {
		return max
	}
}

func filterUnderValue(sums *list.List, t int) *list.List {
	underTValue := list.New()
	for e := sums.Front(); e != nil; e = e.Next() {
		if e.Value.(int) <= t {
			underTValue.PushBack(e.Value.(int))
		}

	}

	return underTValue
}

func sumsOfCombinations(target []int, depth int) *list.List {
	return doRecursive(0, make([]int, depth), 0, target, depth, list.New())
}

func doRecursive(index int, arr []int, start int, target []int, depth int, sums *list.List) *list.List {
	if index == depth {
		sum := 0
		for _, v := range arr {
			print(v)
			print(" ")
			sum += v
		}
		println()

		sums.PushBack(sum)
		return sums
	}

	for i := start; i < len(target); i++ {
		arr[index] = target[i]
		doRecursive(index+1, arr, i+1, target, depth, sums)
	}
	return sums
}
