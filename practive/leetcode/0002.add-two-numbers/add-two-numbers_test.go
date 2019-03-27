package problem0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var l1 = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}

// tcs is testcase slice
var tcs = []struct {
	l1  *ListNode
	l2  *ListNode
	ans *ListNode
}{
	{
		&ListNode{5, nil},
		&ListNode{5, nil},
		&ListNode{0, &ListNode{1, nil}},
	},
	{
		&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
		&ListNode{7, &ListNode{0, &ListNode{8, nil}}},
	},
	{
		&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		&ListNode{5, &ListNode{6, &ListNode{4, &ListNode{3, nil}}}},
		&ListNode{7, &ListNode{0, &ListNode{8, &ListNode{3, nil}}}},
	},
	{
		&ListNode{1, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}},
		&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
		&ListNode{6, &ListNode{6, &ListNode{4, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}},
	},
	// 可以有多个 testcase
}

func TestDigui(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	ret := []int{}
	ret = recursive(l1, ret)
	t.Log(ret)
}

func TestTonum(t *testing.T) {
	ret := []int{2, 4, 3}
	t.Log(reverse(ret))
	ret = []int{5, 6, 4, 3}
	t.Log(reverse(ret))
}

func TestToListNode(t *testing.T) {
	l := &ListNode{}
	toListNode(l, 342)
	t.Log(l.Next.Next.Next)
}

func TestToToListNode(t *testing.T) {
	l := &ListNode{7, &ListNode{0, &ListNode{8, nil}}}
	r := reverseListNode(l)
	t.Log(r.Val)
	t.Log(r.Next)
	t.Log(r.Next.Next)
}

func Test_addTwoNumbers(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, addTwoNumbers(tc.l1, tc.l2), "输入:%v", tc)
	}
}

// Benchmark_addTwoNumbers-8   	 2000000	       745 ns/op
// PASS
func Benchmark_addTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			addTwoNumbers(tc.l1, tc.l2)
		}
	}
}
