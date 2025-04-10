package main

import "algo/service"

func main() {

	service.AddTwoNumbers(
		&service.ListNode{
			Val: 9,
			Next: &service.ListNode{
				Val: 9,
				Next: &service.ListNode{
					Val: 9,
					Next: &service.ListNode{
						Val: 9,
						Next: &service.ListNode{
							Val: 9,
							Next: &service.ListNode{
								Val: 9,
								Next: &service.ListNode{
									Val:  9,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		&service.ListNode{
			Val: 9,
			Next: &service.ListNode{
				Val: 9,
				Next: &service.ListNode{
					Val: 9,
					Next: &service.ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
		},
	)
}
