package main

import (
	"fmt"
)

// ListNode defines the structure for the linked list nodes
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd removes the nth node from the end of the list
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first, second := dummy, dummy

	for i := 0; i <= n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummy.Next
}

// createLinkedList creates a linked list from a slice of integers
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, val := range nums[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// printLinkedList prints the linked list values
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Create a linked list
	nums := []int{1, 2, 3, 4, 5}
	head := createLinkedList(nums)
	fmt.Println("Original linked list:")
	printLinkedList(head)

	// Remove the 2nd node from the end
	n := 2
	head = removeNthFromEnd(head, n)
	fmt.Println("Linked list after removing the 2nd node from the end:")
	printLinkedList(head)
}
