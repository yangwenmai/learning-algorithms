package problems0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	testcases := []struct {
		nums   []int
		target int
		ret    []int
	}{
		{[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
	}
	for _, tc := range testcases {
		ret := twoSum(tc.nums, tc.target)
		assert.Equal(t, tc.ret[0], ret[0], tc.ret[1], ret[1])
	}
}
