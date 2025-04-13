/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	node := head
	nodeList := []*ListNode{}
	for node != nil {
		nodeList = append(nodeList, node)
		node = node.Next
	}
	if len(nodeList) != 0 {
		head = nodeList[len(nodeList)-1]
		node = head
		for i := len(nodeList) - 2; i > -1; i-- {
			node.Next = nodeList[i]
			node = node.Next
		}
		node.Next = nil // This is critical
	}

	return head
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	reversed := reverseList(head)

	current := reversed
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}
