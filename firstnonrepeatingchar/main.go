package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	nonRepeating := FirstNonRepeating("streSs")
	println(nonRepeating == "t")
}

type info struct {
	count      int
	firstIndex int
}

func FirstNonRepeating(str string) string {
	m := make(map[string]*info)
	for i, c := range str {
		s := string(c)
		s = strings.ToLower(s)
		if in := m[s]; in == nil {
			m[s] = &info{1, i}
		} else {
			in.count++
		}
		fmt.Printf("%v %v\n", s, m[s])
	}

	var targets []int
	minIndex := len([]rune(str))
	for k, v := range m {
		if v.count == 1 && v.firstIndex < minIndex {
			targets = append(targets, v.firstIndex)
		}
		fmt.Printf("%v : %v", k, v)
	}

	fmt.Printf("%v", targets)
	if len(targets) == 0 {
		return ""
	} else {
		sort.Ints(targets)
		return string(str[targets[0]])
	}
}
