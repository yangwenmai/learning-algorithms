# [2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order** and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Example:**

```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```

## 耗时

7:40~8:40
约 1 小时 实现了 反转，递归
还剩余一个递归生成 listnode 结果（在这里卡住了1个小时）

23:00~00:00

反转整数后再转 listnode

终于做出来了。。。

耗时：2个小时，性能非常差 
`// Benchmark_addTwoNumbers-8   	 2000000	       745 ns/op`

提交 LeetCode 后发现，我的测试用例未覆盖到比较多的 listnode 的情况，也就是无法转为整型的时候，不能用此方法来转换和实现。

宣告本题解答失败。。。（耗时3小时左右）

错误的原因是因为我未考虑到进位溢出和整数最大值的问题。

>参考伪代码实现：07:37~07:50
