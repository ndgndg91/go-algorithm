package rso1a

import "fmt"

//https://leetcode.com/problems/running-sum-of-1d-array

func TestCase1() {
	arr := []int{1, 2, 3, 4}
	sums := runningSum(arr)
	fmt.Println(sums)
}

func runningSum(nums []int) []int {
	sums := make([]int, 0)
	for index, _ := range nums {
		sum := 0
		for i := 0; i <= index; i++ {
			sum += nums[i]
		}
		sums = append(sums, sum)
	}
	return sums
}
