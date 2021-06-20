package main

import "testing"

func TestTwoSumHashTableOnceLinear(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		answer := twoSumHashTableOnceLinear(nums, 9)
		if answer[0] != 0 {
			t.Error("answer first value must be 0")
		}
		if answer[1] != 1 {
			t.Error("answer second value must be 1")
		}
	})
}
