package leetcode002

type ListNode struct {
	Val  int       `json:"val"`
	Next *ListNode `json:"next"`
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0
	var fakeNode *ListNode = &ListNode{Val: -1, Next: nil}
	var curNode *ListNode = fakeNode

	for carry != 0 || l1 != nil || l2 != nil {
		l1Val := 0
		if l1 != nil {
			l1Val = l1.Val
		}

		l2Val := 0
		if l2 != nil {
			l2Val = l2.Val
		}

		tmpSum := l1Val + l2Val + carry
		carry = tmpSum / 10
		tmpSum %= 10

		curNode.Next = &ListNode{Val: tmpSum, Next: nil}
		curNode = curNode.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return fakeNode.Next
}
