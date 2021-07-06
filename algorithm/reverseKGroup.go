package main

import "fmt"

type ListNode struct {
	data int
	Next *ListNode
}

func main() {
	root := &ListNode{
		data: 1,
		Next: &ListNode{
			data: 2,
			Next: &ListNode{
				data: 3,
				Next: &ListNode{
					data: 4,
					Next: nil,
				},
			},
		},
	}
	PrintList(reverseKGroup(root, 3))
}
func PrintList(root *ListNode) {
	for root != nil {
		fmt.Printf("%d->", root.data)
		root = root.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	stack := make([]*ListNode, 0)
	result := &ListNode{}
	p := result
	for {
		tmp := head
		for tmp != nil && len(stack) < k {
			stack = append(stack, tmp)
			tmp = tmp.Next
		}
		if len(stack) != k {
			p.Next = head
			break
		}
		for len(stack) > 0 {
			p.Next = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = p.Next
		}
		head = tmp
	}
	return result.Next
}
