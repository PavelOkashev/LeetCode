package main

import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
	l1 := &ListNode{}
	l1.Val = 9
	l1.Next = &ListNode{}
	l1.Next.Val = 9
	l2 := &ListNode{}
	l2.Val = 9
	l2.Next = &ListNode{}
	l2.Next.Val = 9
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l3.Val, l3.Next.Val, l3.Next.Next.Val)
}
 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := l1
	var dec int = 0
	for rN := result ; rN != nil || l2 != nil || dec > 0; rN = rN.Next {
		if l1 != nil {
			dec += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			dec +=l2.Val
			l2 = l2.Next
		}
		rN.Val = dec%10
		dec /= 10
		if l2 != nil || dec > 0 || l1 != nil {
			rN.Next = &ListNode{}
		}
	}
	return result
}