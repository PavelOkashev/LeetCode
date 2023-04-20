package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
	l1 := &ListNode{}
	l1.Val = 3
	l1.Next = &ListNode{}
	l1.Next.Val = 9
	l2 := &ListNode{}
	l2.Val = 5
	l2.Next = &ListNode{}
	l2.Next.Val = 2
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l3.Val, l3.Next.Val)
}
 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	resultNext := result
	var count int = 0 
	for l1 != nil && l2 != nil {
		var sum int = l1.Val + l2.Val
		if count > 0 {
			sum+=1
			count-=1
		}
		if sum > 10 {
			sum-= 10
			count+=1
		}
		resultNext.Val = sum
		l1 = l1.Next
		l2 = l2.Next
		if l1 != nil || l2 != nil {
			resultNext.Next = &ListNode{}
			resultNext = resultNext.Next 
		}
	}
	
	return result
}