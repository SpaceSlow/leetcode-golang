package merge_two_sorted_lists

// https://leetcode.com/problems/merge-two-sorted-lists/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	if list1 == nil && list2 == nil {
//		return nil
//	} else if list1 == nil && list2 != nil {
//		return list2
//	} else if list1 != nil && list2 == nil {
//		return list1
//	}
//	if list1.Val <= list2.Val {
//		list1.Next = mergeTwoLists(list1.Next, list2)
//		return list1
//	} else {
//		list2.Next = mergeTwoLists(list1, list2.Next)
//		return list2
//	}
//}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	head := &ListNode{}
	current := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next, list1 = list1, list1.Next
		} else {
			current.Next, list2 = list2, list2.Next
		}
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	return head.Next
}
