package twosum

import "fmt"

//https://leetcode.com/problems/two-sum

func TestCase1() {
	nums := []int{2, 7, 11, 15}
	answer1 := twoSumBruteForce(nums, 9)
	fmt.Println(answer1)
	answer2 := twoSumHashTable(nums, 9)
	fmt.Println()
	fmt.Println(answer2)
	answer3 := twoSumHashTableOnceLinear(nums, 9)
	fmt.Println()
	fmt.Println(answer3)

}

func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				answer := make([]int, 2)
				answer[0] = i
				answer[1] = j
				return answer
			}
		}
	}

	return make([]int, 2)
}

func twoSumHashTable(nums []int, target int) []int {
	numbers := make(map[int]int)
	for i, v := range nums {
		numbers[v] = i
	}

	for i, v := range nums {
		complement := target - v
		value, exists := numbers[complement]
		if exists && value != i {
			answer := make([]int, 2)
			answer[0] = i
			answer[1] = value
			return answer
		}
	}

	return make([]int, 2)
}

func twoSumHashTableOnceLinear(nums []int, target int) []int {
	numbers := make(map[int]int)
	for i, v := range nums {
		complement := target - v
		value, exists := numbers[complement]
		if exists {
			answer := make([]int, 2)
			answer[0] = value
			answer[1] = i
			return answer
		}

		numbers[v] = i
	}

	return make([]int, 2)
}
