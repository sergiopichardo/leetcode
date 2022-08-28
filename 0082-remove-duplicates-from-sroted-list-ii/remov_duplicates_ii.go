package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := new(ListNode)
	dummy.Next = head

	prev := dummy
	current := head

	for current != nil && current.Next != nil {
		if current.Next != nil && current.Val == current.Next.Val {
			for current.Next != nil && current.Val == current.Next.Val {
				current = current.Next
			}
			prev.Next = current.Next
		} else {
			prev = prev.Next
		}
		current = current.Next
	}

	return dummy.Next
}

// SrdjanLinked List template
// template for 3 pointer slide for LL

// func something(head *ListNode) *ListNode {
//   if head == nil || head.Next == nil {
//     return head
//   }
//   dummy := new(ListNode)
//   dummy.Next = head
//   prev, curr, next := dummy, head, head.Next

//   for next != nil {
//     // do some comparison
//     // adjust pointers
//   }
//   // return new head for in place solution
// }
