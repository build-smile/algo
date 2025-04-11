package main

import (
	"algo/service"
	"fmt"
)

func main() {

	// 1
	fmt.Printf("%v", service.AddTwoNumbers(
		//111
		&service.ListNode{
			Val: 1,
			Next: &service.ListNode{
				Val: 1,
				Next: &service.ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
		// 4321
		&service.ListNode{
			Val: 1,
			Next: &service.ListNode{
				Val: 2,
				Next: &service.ListNode{
					Val: 3,
					Next: &service.ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
	))

	//2
	fmt.Println(service.LongestSubstringWithoutDuplicate("abcabcbbgit "))

}
