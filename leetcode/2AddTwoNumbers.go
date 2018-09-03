package leetcode

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	var addition int
	for l1 != nil || l2 != nil {
		var n1, n2 int
		if l1 != nil {
			n1 = l1.Val
		}
		if l2 != nil {
			n2 = l2.Val
		}
		sum := n1 + n2 + addition
		addition = sum / 10
		val := sum % 10
		cur.Next = &ListNode{val, nil}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if addition > 0 {
		cur.Next = &ListNode{addition, nil}
	}
	return dummy.Next
}
