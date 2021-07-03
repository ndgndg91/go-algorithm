package main

import "fmt"

func main() {
	firstCase := []string{"aabb", "abcd", "bbaa", "dada"}
	secondCase := []string{"crazer", "carer", "racar", "caers", "racer"}
	thirdCase := []string{"lazing", "lazy", "lacer"}
	fmt.Printf("%v\n", anagrams("abba", firstCase))
	fmt.Printf("%v\n", anagrams("racer", secondCase))
	fmt.Printf("%v\n", anagrams("laser", thirdCase))
}

func anagrams(word string, words []string) []string {
	var sum int32 = 0
	for _, c := range word {
		sum = sum + c
	}

	var result []string
	for _, w := range words {
		var temp int32 = 0
		for _, c := range w {
			temp = temp + c
		}

		if sum == temp {
			result = append(result, w)
		}
	}

	return result
}
