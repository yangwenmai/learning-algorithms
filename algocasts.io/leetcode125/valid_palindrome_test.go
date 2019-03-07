package leetcode125

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	testcases := []struct {
		str string
		ret bool
	}{
		{
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"race a car",
			false,
		},
		{
			"",
			true,
		},
		{
			" ",
			true,
		},
		{
			"  ",
			true,
		},
	}

	for _, v := range testcases {
		ret := isPalindrome(v.str)
		assert.True(t, ret == v.ret)
	}
}
