package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func newNode(data int) *Node {
	return &Node{Val: data, Next: nil}
}

func reverse(head *Node) *Node {
	curr := head
	var prev *Node = nil
	var next *Node = nil

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func isPalindrome(head *Node) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tail := reverse(slow)

	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}

	return true
}

func insertAtEnd(head *Node, data int) *Node {
	node := newNode(data)
	if head == nil {
		return node
	}

	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = node
	return head
}

func printList(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	var head *Node = nil

	arr := []int{}

	for i := 0; i < len(arr); i++ {
		head = insertAtEnd(head, arr[i])
	}

	printList(head)

	res := isPalindrome(head)
	fmt.Println(res)
}
