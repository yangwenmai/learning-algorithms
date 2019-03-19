package problems0001

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (42.35%)
 * Total Accepted:    1.6M
 * Total Submissions: 3.7M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 *
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 *
 * Example:
 *
 *
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 *
 *
 *
 *
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
