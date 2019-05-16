package problem0146

import (
	"testing"
)

// tcs is testcase slice
var tcs = []struct {
	capacity int
	ans      LRUCache
}{
	// 可以有多个 testcase

}

func Test_Constructor(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if cache.Get(1) != 1 {
		t.Error("case1 : Get(1) -> 1")
	}
	cache.Put(3, 3)
	if cache.Get(2) != -1 {
		t.Error("case2 : Get(2) -> -1")
	}
	cache.Put(4, 4)
	if cache.Get(1) != -1 {
		t.Error("case3 : Get(1) -> -1")
	}
	if cache.Get(3) != 3 {
		t.Error("case4 : Get(3) -> 3")
	}
	if cache.Get(4) != 4 {
		t.Error("case5 : Get(4) -> 4")
	}
}

func Benchmark_Constructor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			Constructor(tc.capacity)
		}
	}
}
