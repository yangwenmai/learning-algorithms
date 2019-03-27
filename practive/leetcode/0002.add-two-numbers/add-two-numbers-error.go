package problem0002

/*
[2] Add Two Numbers

https://leetcode.com/problems/add-two-numbers/description/

algorithms Medium (30.7%)
Total Accepted:    793.5K(793487)
Total Submissions: 2.6M(2582883)
Testcase Example:
`
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
`

You are given two **non-empty** linked lists representing two non-negative integers.
The digits are stored in **reverse order** and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Example:**

```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```
*/
func addTwoNumbersError(l1 *ListNode, l2 *ListNode) *ListNode {
	r := []int{}
	r1 := recursive(l1, r)
	r11 := reverse(r1)
	r2 := recursive(l2, r)
	r21 := reverse(r2)
	i := r11 + r21
	l := toListNode(nil, i)
	rl := reverseListNode(l)
	return rl
}

func reverseListNode(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	} else {
		newHead := reverseListNode(l.Next)
		l.Next.Next = l
		l.Next = nil
		return newHead
	}
}

func recursive(l *ListNode, vals []int) []int {
	if l != nil {
		val := l.Val
		vals = append(vals, val)
		if l.Next != nil {
			vals = recursive(l.Next, vals)
		}
	}
	return vals
}

func reverse(ret []int) int {
	r := 0
	// 2 4 3
	for k, v := range ret {
		pow := 1
		for i := 0; i < k; i++ {
			pow *= 10
		}
		r += v * pow
	}
	return r
}

func order(ret []int) int {
	r := 0
	// 2 4 3
	for k, v := range ret {
		pow := 1
		for i := len(ret) - 1 - k; i > 0; i-- {
			pow *= 10
		}
		r += v * pow
	}
	return r
}

func toListNode(l *ListNode, r int) *ListNode {
	ni := r % 10
	lr := &ListNode{}
	lr.Val = ni
	lr.Next = l
	rf := r / 10
	if rf == 0 {
		return lr
	}
	return toListNode(lr, rf)
}
