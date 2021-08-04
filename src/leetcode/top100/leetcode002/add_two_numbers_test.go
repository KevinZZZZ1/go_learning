package leetcode002

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	var list1 ListNode = ListNode{Val: 9, Next: nil}
	list1.Next = &ListNode{Val: 9, Next: nil}
	list1.Next.Next = &ListNode{Val: 9, Next: nil}
	list1.Next.Next.Next = &ListNode{Val: 9, Next: nil}

	var list2 ListNode = ListNode{Val: 9, Next: nil}
	list2.Next = &ListNode{Val: 9, Next: nil}
	list2.Next.Next = &ListNode{Val: 9, Next: nil}
	list2.Next.Next.Next = &ListNode{Val: 9, Next: nil}
	list2.Next.Next.Next.Next = &ListNode{Val: 9, Next: nil}
	list2.Next.Next.Next.Next.Next = &ListNode{Val: 9, Next: nil}
	list2.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9, Next: nil}
	list2.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9, Next: nil}

	AddTwoNumbers(&list1, &list2)

}
