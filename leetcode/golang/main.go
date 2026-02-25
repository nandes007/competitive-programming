package main

import (
	"fmt"

	"github.com/nandes007/leetcode/golang/exercise"
	"github.com/nandes007/leetcode/golang/model"
)

func printLinkedList(node *model.ListNode) {
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println("nil")
}

func main() {
	s1 := "babad"
	s2 := "abbd"

	exercise.LongestPalindrome(s1)
	exercise.LongestPalindrome(s2)
}
