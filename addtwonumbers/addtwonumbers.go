package addtwonumbers

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func TestCase1() {
	l1Node1 := &ListNode{2, nil}
	l1Node2 := &ListNode{4, nil}
	l1Node3 := &ListNode{3, nil}
	l1Node1.Next = l1Node2
	l1Node2.Next = l1Node3
	l1 := l1Node1

	l2Node1 := &ListNode{5, nil}
	l2Node2 := &ListNode{6, nil}
	l2Node3 := &ListNode{ 4, nil}
	l2Node1.Next = l2Node2
	l2Node2.Next = l2Node3
	l2 := l2Node1
	numbers := addTwoNumbers(l1, l2)
	for numbers != nil {
		val := numbers.Val
		fmt.Println(val)
		numbers = numbers.Next
	}
}

func TestCase2() {
	l1Node1 := &ListNode{9, nil}
	l1Node2 := &ListNode{9, nil}
	l1Node3 := &ListNode{9, nil}
	l1Node4 := &ListNode{9, nil}
	l1Node5 := &ListNode{9, nil}
	l1Node6 := &ListNode{9, nil}
	l1Node7 := &ListNode{9, nil}
	l1Node1.Next = l1Node2
	l1Node2.Next = l1Node3
	l1Node3.Next = l1Node4
	l1Node4.Next = l1Node5
	l1Node5.Next = l1Node6
	l1Node6.Next = l1Node7

	l1 := l1Node1

	l2Node1 := &ListNode{9, nil}
	l2Node2 := &ListNode{9, nil}
	l2Node3 := &ListNode{9, nil}
	l2Node4 := &ListNode{9, nil}
	l2Node1.Next = l2Node2
	l2Node2.Next = l2Node3
	l2Node3.Next = l2Node4

	l2 := l2Node1

	numbers := addTwoNumbers(l1, l2)
	for numbers != nil {
		val := numbers.Val
		fmt.Println(val)
		numbers = numbers.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode  {
	dummyHead := &ListNode{0, nil}
	curr := dummyHead

	carry := 0
	for l1 != nil || l2 != nil {
		x:=0
		y:=0
		if l1 != nil {x = l1.Val}
		if l2 != nil {y = l2.Val}
		sum := carry + x + y
		carry = sum / 10

		val := sum % 10
		curr.Next = &ListNode{val, nil}
		curr = curr.Next

		if l1 != nil {l1 = l1.Next}
		if l2 != nil {l2 = l2.Next}
	}

	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}

	return dummyHead.Next
}