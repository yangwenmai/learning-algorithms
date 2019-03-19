package problem0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	l1 *ListNode
 l2 *ListNode
	ans *ListNode
}{



	// 可以有多个 testcase
}

func Test_addTwoNumbers(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, addTwoNumbers(tc.l1, tc.l2), "输入:%v", tc)
	}
}

func Benchmark_addTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			addTwoNumbers(tc.l1, tc.l2)
		}
	}
}
