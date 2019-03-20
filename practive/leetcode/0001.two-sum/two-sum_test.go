package problem0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums   []int
	target int
	ans    []int
}{
	// 可以有多个 testcase
	{[]int{2, 7, 11, 15},
		9,
		[]int{0, 1},
	},
	{[]int{8, 9, 11, 25},
		33,
		[]int{0, 3},
	},
	{[]int{1, 2, 3, 6, 8, 11},
		10,
		[]int{1, 4},
	},
	{[]int{1, 2, 3, 6, 8, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		21,
		[]int{0, 14},
	},
}

func Test_twoSum(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, twoSum(tc.nums, tc.target), "输入:%v", tc)
	}
}

func Test_twoSum2(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, twoSum2(tc.nums, tc.target), "输入:%v", tc)
	}
}

func Test_twoSum3(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, twoSum3(tc.nums, tc.target), "输入:%v", tc)
	}
}

// Benchmark_twoSum-8   	20000000	       105 ns/op
func Benchmark_twoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			twoSum(tc.nums, tc.target)
		}
	}
}

// Benchmark_twoSum2-8   	20000000	        83.6 ns/op
func Benchmark_twoSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			twoSum2(tc.nums, tc.target)
		}
	}
}

// Benchmark_twoSum3-8   	 5000000	       271 ns/op
func Benchmark_twoSum3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			twoSum3(tc.nums, tc.target)
		}
	}
}
