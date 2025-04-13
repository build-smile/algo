package service

import (
	"reflect"
	"testing"
)

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, v := range values[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name   string
		l1     []int
		l2     []int
		output []int
	}{
		{"same_length_no_carry", []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{"different_length_no_carry", []int{1}, []int{9, 9, 9}, []int{0, 0, 0, 1}},
		{"same_length_with_carry", []int{9, 9, 9}, []int{1, 1, 1}, []int{0, 1, 1, 1}},
		{"different_length_with_carry", []int{1}, []int{9, 9}, []int{0, 0, 1}},
		{"one_empty_list", nil, []int{2, 3}, []int{2, 3}},
		{"both_empty_lists", nil, nil, nil},
		{"multiple_carries", []int{9, 9, 9}, []int{9, 9, 9}, []int{8, 9, 9, 1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l1 := createList(tc.l1)
			l2 := createList(tc.l2)
			expected := createList(tc.output)
			result := AddTwoNumbers(l1, l2)

			if !reflect.DeepEqual(listToSlice(result), listToSlice(expected)) {
				t.Errorf("expected %v, got %v", listToSlice(expected), listToSlice(result))
			}
		})
	}
}
