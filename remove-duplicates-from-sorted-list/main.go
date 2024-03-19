package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}
//
//func deleteDuplicates(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	prev := head
//	current := head.Next
//
//	for current != nil {
//		if current.Val != prev.Val {
//			prev.Next, prev = current, current
//		} else {
//			prev.Next = nil
//		}
//		current = current.Next
//	}
//
//	return head
//}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head

	for current.Next != nil {
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}
