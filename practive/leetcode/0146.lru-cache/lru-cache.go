package problem0146

/*
	[146] LRU Cache

	https://leetcode.com/problems/lru-cache/description/

	algorithms Hard (25.2%)
	Total Accepted:    291.9K(291854)
	Total Submissions: 1.2M(1155924)

Design and implement a data structure for [Least Recently Used (LRU) cache](https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU). It should support the following operations: `get` and `put`.

`get(key)` - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.


`put(key, value)` - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

**Follow up:**

Could you do both operations in **O(1)** time complexity?

**Example:**


```
LRUCache cache = new LRUCache( 2 \/* capacity *\/ );
cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
```
*/

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

type LRUCache struct {
	capacity int
	mmap     map[int]*node
	head     *node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{capacity: capacity}
	lru.mmap = make(map[int]*node, capacity)
	lru.head = &node{key: -1, value: -1}
	cur := lru.head
	for i := 1; i < capacity; i++ {
		tmp := &node{key: -1, value: -1}
		cur.next = tmp
		tmp.prev = cur
		cur = tmp
	}
	cur.next = lru.head
	lru.head.prev = cur
	return lru
}

func (lru *LRUCache) move2head(cur *node) {
	if cur == lru.head {
		lru.head = lru.head.prev
		return
	}
	// 有点绕。。。
	cur.prev.next = cur.next
	cur.next.prev = cur.prev

	head := lru.head
	cur.next = head.next
	cur.next.prev = cur
	head.next = cur
	cur.prev = head
}

func (lru *LRUCache) Get(key int) int {
	if r, ok := lru.mmap[key]; ok {
		lru.move2head(r)
		return r.value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if r, ok := lru.mmap[key]; ok {
		r.value = value
		lru.move2head(r)
	} else {
		delete(lru.mmap, lru.head.key)
		lru.head.key = key
		lru.head.value = value
		lru.mmap[key] = lru.head
		lru.head = lru.head.prev
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
