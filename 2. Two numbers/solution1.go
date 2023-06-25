package addTwoNumbers

type ListNode struct {
     Val int
     Next *ListNode
}

func lenOfList(l *ListNode) int {
	count := 0
	for curr := l; curr != nil; curr = curr.Next {
		count++
	}

	return count
}

func appendToList(head *ListNode, elem int) *ListNode {
	if head == nil {
		head = &ListNode{elem, nil}
		return head
	}
	temp := head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = &ListNode{elem, nil}
	return head
}

func fillWithZeroes(l *ListNode, n int) *ListNode {
	for i := 0; i < n; i++ {
		l = appendToList(l, 0)
	}

	return l
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curr1, curr2 := l1, l2
	if lenOfList(l1) > lenOfList(l2) {
		curr2 = fillWithZeroes(curr2, lenOfList(l1)-lenOfList(l2))
	} else {
		curr1 = fillWithZeroes(curr1, lenOfList(l2)-lenOfList(l1))
	}

	var res *ListNode = nil
	tmp := false
	for curr1 != nil {
		sum := curr1.Val + curr2.Val
		if tmp {
			sum++
			tmp = false
		}
		if sum > 9 {
			tmp = true
			sum = sum % 10
		}

		res = appendToList(res, sum)

		curr2 = curr2.Next
		curr1 = curr1.Next
	}

	if tmp {
		res = appendToList(res, 1)
	}

	return res
}