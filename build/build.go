package build

import "github.com/yangwenmai/learning-algorithms/leetcode"

type Builder struct {
	LC *leetcode.Leetcode
}

type Build interface {
	Build()
}

func (b Builder) Build() {
	b.Build()
}
