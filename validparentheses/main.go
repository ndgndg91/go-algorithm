package main

import "errors"

type stack struct {
	values []rune
}

func newStack() *stack {
	var values []rune
	return &stack{values: values}
}

func (s *stack) push(value rune) {
	s.values = append(s.values, value)
}

func (s *stack) pop() (rune, error) {
	l := len(s.values)
	if l == 0 {
		return 0, errors.New("empty stack")
	}

	res := s.values[l-1]
	s.values = s.values[:l-1]
	return res, nil
}

func (s *stack) isEmpty() bool {
	return len(s.values) == 0
}

func main() {
	valid := ValidParentheses("()")
	println(valid)
}

func ValidParentheses(parens string) bool {
	stack := newStack()
	for _, c := range parens {
		if c == '(' {
			stack.push(c)
		} else {
			var pop rune
			var err error
			if pop, err = stack.pop(); err != nil {
				return false
			}

			if string(pop) != "(" {
				return false
			}
		}
	}

	if !stack.isEmpty() {
		return false
	}

	return true
}

func ValidParentheses2(parens string) bool {
	b := 0
	for _, c := range parens {
		if c == '(' {
			b++
		} else {
			b--
			if b < 0 {
				return false
			}
		}
	}
	return b == 0
}
