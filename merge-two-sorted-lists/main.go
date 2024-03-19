package merge_two_sorted_lists

// https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	retList, curListNode1, curListNode2 := new(ListNode), new(ListNode), new(ListNode)
	if list1.Val > list2.Val {
		retList = list2
		curListNode1 = list1
		curListNode2 = list2.Next
	} else {
		retList = list1
		curListNode1 = list1.Next
		curListNode2 = list2
	}
	firstNode := retList

	for {
		if curListNode2 == nil || (curListNode1 != nil || curListNode2 != nil) && curListNode1.Val < curListNode2.Val {
			retList.Next = curListNode1
			retList = retList.Next
			if curListNode1.Next == nil {
				break
			}
			curListNode1 = curListNode1.Next
		} else {
			retList.Next = curListNode2
			retList = retList.Next
			if curListNode2.Next == nil {
				break
			}
			curListNode2 = curListNode2.Next
		}
	}

	return firstNode
}
