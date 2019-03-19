package problem0007

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	x int
	ans int
}{



	// 可以有多个 testcase
}

func Test_reverse(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, reverse(tc.x), "输入:%v", tc)
	}
}

func Benchmark_reverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reverse(tc.x)
		}
	}
}
