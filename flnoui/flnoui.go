package flnoui

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/

func TestCase1() {
	arr := []int{5, 5, 4}
	count := findLeastNumOfUniqueIntegers(arr, 1)
	fmt.Println(count)
	fmt.Println(count == 1)
}

func TestCase2() {
	arr := []int{4, 3, 1, 1, 3, 3, 2}
	count := findLeastNumOfUniqueIntegers(arr, 3)
	fmt.Println(count)
	fmt.Println(count == 2)
}

type kv struct {
	Key   int
	Value int
}

func findLeastNumOfUniqueIntegers(arr []int, k int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v] = m[v] + 1
	}

	if k == 0 {
		return len(m)
	}

	sortByValue := make([]kv, 0)

	for k, v := range m {
		sortByValue = append(sortByValue, kv{k, v})
	}

	sort.Slice(sortByValue, func(i, j int) bool {
		return sortByValue[i].Value < sortByValue[j].Value
	})

out:
	for _, v := range sortByValue {
		fmt.Printf("%d, %d\n", v.Key, v.Value)
		for v.Value > 0 {
			v.Value = v.Value - 1
			k--
			if v.Value == 0 {
				delete(m, v.Key)
			}
			if k == 0 {
				break out
			}
		}
	}

	fmt.Println(m)
	return len(m)
}
