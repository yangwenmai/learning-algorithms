package problem0009

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	x   int
	ans bool
}{
	{-121, false},
	{121, true},
	// 可以有多个 testcase
}

func Test_isPalindrome(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, isPalindrome(tc.x), "输入:%v", tc)
	}
}

func Benchmark_isPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isPalindrome(tc.x)
		}
	}
}
