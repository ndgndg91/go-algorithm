package lswrc

import (
	"fmt"
)

func TestCase1()  {
	length := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(length == 3)
}

func TestCase2()  {
	length := lengthOfLongestSubstring("pwwkew")
	fmt.Println(length == 3)
}

func TestCase3()  {
	length := lengthOfLongestSubstring("dvdf")
	fmt.Println(length == 3)
}

func TestCase4()  {
	length := lengthOfLongestSubstring("abba")
	fmt.Println(length == 2)
}

func lengthOfLongestSubstring(s string) int  {
	m := make(map[int32]int)

	answer := 0
	start := 0
	for i, v := range s{
		index := m[v]
		if index != 0 {
			start = max(index, start)
		}

		fmt.Println(answer, i + 1, start)
		answer = max(answer, i + 1 - start)
		m[v] = i + 1
	}

	return answer
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}