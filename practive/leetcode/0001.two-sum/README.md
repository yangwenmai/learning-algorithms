# [1. Two Sum](https://leetcode.com/problems/two-sum/)

Given an array of integers, return **indices** of the two numbers such that they add up to a specific target.

You may assume that each input would have **_exactly_** one solution, and you may not use the _same_ element twice.

**Example:**

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## 题解

- 暴力破解，两层 for 循环。
    >时间复杂度：O(n^2), 空间复杂度：O(1)
- hash 表
    >时间复杂度：O(n), 空间复杂度：O(n)
    
Benchmark 的结果来看，在数量比较小的时候，暴力破解会更快。

延伸出来的一个问题是？

在实际使用过程中，我们应该怎么权衡何时使用这两种算法呢？
