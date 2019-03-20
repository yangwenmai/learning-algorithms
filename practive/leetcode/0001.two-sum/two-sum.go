package problem0001

/*
	[1] Two Sum

	https://leetcode.com/problems/two-sum/description/

	algorithms Easy (42.5%)
	Total Accepted:    1.6M(1575246)
	Total Submissions: 3.7M(3705940)

Given an array of integers, return **indices** of the two numbers such that they add up to a specific target.

You may assume that each input would have **_exactly_** one solution, and you may not use the _same_ element twice.

**Example:**

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```
*/

func twoSum(nums []int, target int) []int {
	var idx []int
	for i, num1 := range nums {
		for j, num2 := range nums {
			if i >= j {
				continue
			}
			if target-num1 == num2 {
				idx = append(idx, i)
				idx = append(idx, j)
				return idx
			}
		}
	}
	return idx
}

// 优化版本
func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

// hash
func twoSum3(nums []int, target int) []int {
	hv := map[int]int{}
	for i := 0; i < len(nums); i++ {
		need := target - nums[i]
		idx, ok := hv[need]
		if ok {
			return []int{idx, i}
		} else {
			hv[nums[i]] = i
		}
	}
	return []int{-1, -1}
}
