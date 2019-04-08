package problem0009

/*
	[9] Palindrome Number

	https://leetcode.com/problems/palindrome-number/description/

	algorithms Easy (42.6%)
	Total Accepted:    545.7K(545702)
	Total Submissions: 1.3M(1279847)

	Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

**Example 1:**

```
Input: 121
Output: true
```

**Example 2:**

```
Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```

**Example 3:**

```
Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```

**Follow up:**

Coud you solve it without converting the integer to a string?
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	y := 0
	for tmp > 0 {
		num := tmp % 10
		y = y*10 + num
		tmp = tmp / 10
	}
	return y == x
}
