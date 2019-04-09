# 算法学习(learning-algorithms) 及 [LeetCode](https://leetcode.com) 的题解（Go 版本）

[![LeetCode badge](https://leetcode-badge.chyroc.cn/?name=maiyang)](https://leetcode-badge.chyroc.cn/?name=maiyang)
[![LeetCode Ranking](https://img.shields.io/badge/maiyang-100001-blue.svg)](https://leetcode.com/maiyang/)
[![Build Status](https://travis-ci.org/yangwenmai/learning-algorithms.svg?branch=master)](https://travis-ci.org/yangwenmai/learning-algorithms)
[![Go Report Card](https://goreportcard.com/badge/github.com/yangwenmai/learning-algorithms)](https://goreportcard.com/report/github.com/yangwenmai/learning-algorithms)
[![Documentation](https://godoc.org/github.com/yangwenmai/learning-algorithms?status.svg)](http://godoc.org/github.com/yangwenmai/learning-algorithms)
[![Coverage Status](https://coveralls.io/repos/github/yangwenmai/learning-algorithms/badge.svg?branch=master)](https://coveralls.io/github/yangwenmai/learning-algorithms?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/yangwenmai/learning-algorithms.svg)](https://github.com/yangwenmai/learning-algorithms/issues)
[![license](https://img.shields.io/github/license/yangwenmai/learning-algorithms.svg?maxAge=2592000)](https://github.com/yangwenmai/learning-algorithms/LICENSE)
[![codecov](https://codecov.io/gh/yangwenmai/learning-algorithms/branch/master/graph/badge.svg)](https://codecov.io/gh/yangwenmai/learning-algorithms)

如果你正在算法之门外彷徨，如果你正在发愁不知道该怎么进行算法学习，或许本项目能够适合你。

本项目是我准备的算法学习历程，包括我对算法书籍的整理，算法实践（Go 实现），以及各种算法的理解，也还会有一些有关方面的经验总结。

>希望有更多的人参与进来，咱们一起前行，一起精进。

Inspire by [aQuaYi/LeetCode-in-Go](https://github.com/aQuaYi/LeetCode-in-Go), [WindomZ/leetcode-graphql](https://github.com/WindomZ/leetcode-graphql)

## LeetCode 生成器

> 可以生成 LeetCode 模板，还可以更新 README ，还可以添加待办任务等部分琐碎的工作。

### 配置方法

1. 在 `.gitignore` 中，添加 `*.toml`（避免用户名密码暴露）
1. 在 `learning-algorithms` 目录下，添加文本文件`config.toml`。
1. 把以下内容复制到 `config.toml` 中。

```toml
Username="your leetcode username"
Password="your leetcode password"
```

## 进度

>统计范围：能提交 Go 解答的免费算法题。

|     |Easy|Medium|Hard|Total|
|:---:|:---:|:---:|:---:|:---:|
|**Accepted**|4|1|1|6|
|**Total**|237|403|173|813|

### 开发环境

- macOS 10.14.3
- go version go1.12.3 darwin/amd64
- GoLand

## 完成清单

|题号|题目|通过率|难度|收藏|
|:-:|:-|:-: | :-: | :-: |
|[0001](https://leetcode.com/problems/two-sum/)|✅[Two Sum](./practive/leetcode/0001.two-sum)|43%|Easy||
|[0002](https://leetcode.com/problems/add-two-numbers/)|✅[Add Two Numbers](./practive/leetcode/0002.add-two-numbers)|30%|Medium||
|[0003](https://leetcode.com/problems/longest-substring-without-repeating-characters/)| * Longest Substring Without Repeating Characters|28%|Medium||
|[0004](https://leetcode.com/problems/median-of-two-sorted-arrays/)| * Median of Two Sorted Arrays|25%|Hard||
|[0005](https://leetcode.com/problems/longest-palindromic-substring/)| * Longest Palindromic Substring|26%|Medium||
|[0006](https://leetcode.com/problems/zigzag-conversion/)| * ZigZag Conversion|31%|Medium||
|[0007](https://leetcode.com/problems/reverse-integer/)|✅[Reverse Integer](./practive/leetcode/0007.reverse-integer)|25%|Easy||
|[0008](https://leetcode.com/problems/string-to-integer-atoi/)| * String to Integer (atoi)|14%|Medium||
|[0009](https://leetcode.com/problems/palindrome-number/)|✅[Palindrome Number](./practive/leetcode/0009.palindrome-number)|42%|Easy||
|[0010](https://leetcode.com/problems/regular-expression-matching/)| * Regular Expression Matching|25%|Hard||
|[0011](https://leetcode.com/problems/container-with-most-water/)| * Container With Most Water|43%|Medium||
|[0012](https://leetcode.com/problems/integer-to-roman/)| * Integer to Roman|50%|Medium||
|[0013](https://leetcode.com/problems/roman-to-integer/)| * Roman to Integer|51%|Easy||
|[0014](https://leetcode.com/problems/longest-common-prefix/)| * Longest Common Prefix|33%|Easy||
|[0015](https://leetcode.com/problems/3sum/)| * 3Sum|23%|Medium||
|[0016](https://leetcode.com/problems/3sum-closest/)| * 3Sum Closest|42%|Medium||
|[0017](https://leetcode.com/problems/letter-combinations-of-a-phone-number/)| * Letter Combinations of a Phone Number|40%|Medium||
|[0018](https://leetcode.com/problems/4sum/)| * 4Sum|30%|Medium||
|[0019](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)| * Remove Nth Node From End of List|34%|Medium||
|[0020](https://leetcode.com/problems/valid-parentheses/)| * Valid Parentheses|36%|Easy||
|[0021](https://leetcode.com/problems/merge-two-sorted-lists/)| * Merge Two Sorted Lists|46%|Easy||
|[0022](https://leetcode.com/problems/generate-parentheses/)| * Generate Parentheses|53%|Medium||
|[0023](https://leetcode.com/problems/merge-k-sorted-lists/)| * Merge k Sorted Lists|33%|Hard||
|[0024](https://leetcode.com/problems/swap-nodes-in-pairs/)| * Swap Nodes in Pairs|43%|Medium||
|[0025](https://leetcode.com/problems/reverse-nodes-in-k-group/)| * Reverse Nodes in k-Group|35%|Hard||
|[0026](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)| * Remove Duplicates from Sorted Array|40%|Easy||
|[0027](https://leetcode.com/problems/remove-element/)| * Remove Element|44%|Easy||
|[0028](https://leetcode.com/problems/implement-strstr/)| * Implement strStr()|31%|Easy||
|[0029](https://leetcode.com/problems/divide-two-integers/)| * Divide Two Integers|16%|Medium||
|[0030](https://leetcode.com/problems/substring-with-concatenation-of-all-words/)| * Substring with Concatenation of All Words|23%|Hard||
|[0031](https://leetcode.com/problems/next-permutation/)| * Next Permutation|30%|Medium||
|[0032](https://leetcode.com/problems/longest-valid-parentheses/)| * Longest Valid Parentheses|25%|Hard||
|[0033](https://leetcode.com/problems/search-in-rotated-sorted-array/)| * Search in Rotated Sorted Array|32%|Medium||
|[0034](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)| * Find First and Last Position of Element in Sorted Array|33%|Medium||
|[0035](https://leetcode.com/problems/search-insert-position/)| * Search Insert Position|40%|Easy||
|[0036](https://leetcode.com/problems/valid-sudoku/)| * Valid Sudoku|42%|Medium||
|[0037](https://leetcode.com/problems/sudoku-solver/)| * Sudoku Solver|36%|Hard||
|[0038](https://leetcode.com/problems/count-and-say/)| * Count and Say|39%|Easy||
|[0039](https://leetcode.com/problems/combination-sum/)| * Combination Sum|47%|Medium||
|[0040](https://leetcode.com/problems/combination-sum-ii/)| * Combination Sum II|40%|Medium||
|[0041](https://leetcode.com/problems/first-missing-positive/)| * First Missing Positive|28%|Hard||
|[0042](https://leetcode.com/problems/trapping-rain-water/)|✅[Trapping Rain Water](./practive/leetcode/0042.trapping-rain-water)|42%|Hard||
|[0043](https://leetcode.com/problems/multiply-strings/)| * Multiply Strings|30%|Medium||
|[0044](https://leetcode.com/problems/wildcard-matching/)| * Wildcard Matching|22%|Hard||
|[0045](https://leetcode.com/problems/jump-game-ii/)| * Jump Game II|27%|Hard||
|[0046](https://leetcode.com/problems/permutations/)| * Permutations|54%|Medium||
|[0047](https://leetcode.com/problems/permutations-ii/)| * Permutations II|39%|Medium||
|[0048](https://leetcode.com/problems/rotate-image/)| * Rotate Image|47%|Medium||
|[0049](https://leetcode.com/problems/group-anagrams/)| * Group Anagrams|45%|Medium||
|[0050](https://leetcode.com/problems/powx-n/)| * Pow(x, n)|27%|Medium||
|[0051](https://leetcode.com/problems/n-queens/)| * N-Queens|38%|Hard||
|[0052](https://leetcode.com/problems/n-queens-ii/)| * N-Queens II|51%|Hard||
|[0053](https://leetcode.com/problems/maximum-subarray/)| * Maximum Subarray|43%|Easy||
|[0054](https://leetcode.com/problems/spiral-matrix/)| * Spiral Matrix|29%|Medium||
|[0055](https://leetcode.com/problems/jump-game/)| * Jump Game|31%|Medium||
|[0056](https://leetcode.com/problems/merge-intervals/)| * Merge Intervals|35%|Medium||
|[0057](https://leetcode.com/problems/insert-interval/)| * Insert Interval|30%|Hard||
|[0058](https://leetcode.com/problems/length-of-last-word/)| * Length of Last Word|32%|Easy||
|[0059](https://leetcode.com/problems/spiral-matrix-ii/)| * Spiral Matrix II|45%|Medium||
|[0060](https://leetcode.com/problems/permutation-sequence/)| * Permutation Sequence|32%|Medium||
|[0061](https://leetcode.com/problems/rotate-list/)| * Rotate List|26%|Medium||
|[0062](https://leetcode.com/problems/unique-paths/)| * Unique Paths|46%|Medium||
|[0063](https://leetcode.com/problems/unique-paths-ii/)| * Unique Paths II|33%|Medium||
|[0064](https://leetcode.com/problems/minimum-path-sum/)| * Minimum Path Sum|46%|Medium||
|[0065](https://leetcode.com/problems/valid-number/)| * Valid Number|13%|Hard||
|[0066](https://leetcode.com/problems/plus-one/)| * Plus One|40%|Easy||
|[0067](https://leetcode.com/problems/add-binary/)| * Add Binary|38%|Easy||
|[0068](https://leetcode.com/problems/text-justification/)| * Text Justification|22%|Hard||
|[0069](https://leetcode.com/problems/sqrtx/)| * Sqrt(x)|30%|Easy||
|[0070](https://leetcode.com/problems/climbing-stairs/)| * Climbing Stairs|43%|Easy||
|[0071](https://leetcode.com/problems/simplify-path/)| * Simplify Path|28%|Medium||
|[0072](https://leetcode.com/problems/edit-distance/)| * Edit Distance|36%|Hard||
|[0073](https://leetcode.com/problems/set-matrix-zeroes/)| * Set Matrix Zeroes|39%|Medium||
|[0074](https://leetcode.com/problems/search-a-2d-matrix/)| * Search a 2D Matrix|34%|Medium||
|[0075](https://leetcode.com/problems/sort-colors/)| * Sort Colors|41%|Medium||
|[0076](https://leetcode.com/problems/minimum-window-substring/)| * Minimum Window Substring|30%|Hard||
|[0077](https://leetcode.com/problems/combinations/)| * Combinations|46%|Medium||
|[0078](https://leetcode.com/problems/subsets/)| * Subsets|51%|Medium||
|[0079](https://leetcode.com/problems/word-search/)| * Word Search|30%|Medium||
|[0080](https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/)| * Remove Duplicates from Sorted Array II|39%|Medium||
|[0081](https://leetcode.com/problems/search-in-rotated-sorted-array-ii/)| * Search in Rotated Sorted Array II|32%|Medium||
|[0082](https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/)| * Remove Duplicates from Sorted List II|32%|Medium||
|[0083](https://leetcode.com/problems/remove-duplicates-from-sorted-list/)| * Remove Duplicates from Sorted List|42%|Easy||
|[0084](https://leetcode.com/problems/largest-rectangle-in-histogram/)| * Largest Rectangle in Histogram|30%|Hard||
|[0085](https://leetcode.com/problems/maximal-rectangle/)| * Maximal Rectangle|32%|Hard||
|[0086](https://leetcode.com/problems/partition-list/)| * Partition List|36%|Medium||
|[0087](https://leetcode.com/problems/scramble-string/)| * Scramble String|31%|Hard||
|[0088](https://leetcode.com/problems/merge-sorted-array/)| * Merge Sorted Array|35%|Easy||
|[0089](https://leetcode.com/problems/gray-code/)| * Gray Code|45%|Medium||
|[0090](https://leetcode.com/problems/subsets-ii/)| * Subsets II|41%|Medium||
|[0091](https://leetcode.com/problems/decode-ways/)| * Decode Ways|22%|Medium||
|[0092](https://leetcode.com/problems/reverse-linked-list-ii/)| * Reverse Linked List II|34%|Medium||
|[0093](https://leetcode.com/problems/restore-ip-addresses/)| * Restore IP Addresses|31%|Medium||
|[0094](https://leetcode.com/problems/binary-tree-inorder-traversal/)| * Binary Tree Inorder Traversal|55%|Medium||
|[0095](https://leetcode.com/problems/unique-binary-search-trees-ii/)| * Unique Binary Search Trees II|35%|Medium||
|[0096](https://leetcode.com/problems/unique-binary-search-trees/)| * Unique Binary Search Trees|45%|Medium||
|[0097](https://leetcode.com/problems/interleaving-string/)| * Interleaving String|27%|Hard||
|[0098](https://leetcode.com/problems/validate-binary-search-tree/)| * Validate Binary Search Tree|25%|Medium||
|[0099](https://leetcode.com/problems/recover-binary-search-tree/)| * Recover Binary Search Tree|34%|Hard||
|[0100](https://leetcode.com/problems/same-tree/)| * Same Tree|49%|Easy||
|[0101](https://leetcode.com/problems/symmetric-tree/)| * Symmetric Tree|43%|Easy||
|[0102](https://leetcode.com/problems/binary-tree-level-order-traversal/)| * Binary Tree Level Order Traversal|47%|Medium||
|[0103](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/)| * Binary Tree Zigzag Level Order Traversal|41%|Medium||
|[0104](https://leetcode.com/problems/maximum-depth-of-binary-tree/)| * Maximum Depth of Binary Tree|59%|Easy||
|[0105](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)| * Construct Binary Tree from Preorder and Inorder Traversal|40%|Medium||
|[0106](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)| * Construct Binary Tree from Inorder and Postorder Traversal|38%|Medium||
|[0107](https://leetcode.com/problems/binary-tree-level-order-traversal-ii/)| * Binary Tree Level Order Traversal II|46%|Easy||
|[0108](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)| * Convert Sorted Array to Binary Search Tree|49%|Easy||
|[0109](https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/)| * Convert Sorted List to Binary Search Tree|40%|Medium||
|[0110](https://leetcode.com/problems/balanced-binary-tree/)| * Balanced Binary Tree|40%|Easy||
|[0111](https://leetcode.com/problems/minimum-depth-of-binary-tree/)| * Minimum Depth of Binary Tree|35%|Easy||
|[0112](https://leetcode.com/problems/path-sum/)| * Path Sum|37%|Easy||
|[0113](https://leetcode.com/problems/path-sum-ii/)| * Path Sum II|39%|Medium||
|[0114](https://leetcode.com/problems/flatten-binary-tree-to-linked-list/)| * Flatten Binary Tree to Linked List|41%|Medium||
|[0115](https://leetcode.com/problems/distinct-subsequences/)| * Distinct Subsequences|34%|Hard||
|[0118](https://leetcode.com/problems/pascals-triangle/)| * Pascal&#39;s Triangle|45%|Easy||
|[0119](https://leetcode.com/problems/pascals-triangle-ii/)| * Pascal&#39;s Triangle II|42%|Easy||
|[0120](https://leetcode.com/problems/triangle/)| * Triangle|38%|Medium||
|[0121](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)| * Best Time to Buy and Sell Stock|46%|Easy||
|[0122](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/)| * Best Time to Buy and Sell Stock II|51%|Easy||
|[0123](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/)| * Best Time to Buy and Sell Stock III|33%|Hard||
|[0124](https://leetcode.com/problems/binary-tree-maximum-path-sum/)| * Binary Tree Maximum Path Sum|29%|Hard||
|[0125](https://leetcode.com/problems/valid-palindrome/)|✅[Valid Palindrome](./practive/leetcode/0125.valid-palindrome)|30%|Easy||
|[0126](https://leetcode.com/problems/word-ladder-ii/)| * Word Ladder II|17%|Hard||
|[0127](https://leetcode.com/problems/word-ladder/)| * Word Ladder|23%|Medium||
|[0128](https://leetcode.com/problems/longest-consecutive-sequence/)| * Longest Consecutive Sequence|41%|Hard||
|[0129](https://leetcode.com/problems/sum-root-to-leaf-numbers/)| * Sum Root to Leaf Numbers|41%|Medium||
|[0130](https://leetcode.com/problems/surrounded-regions/)| * Surrounded Regions|22%|Medium||
|[0131](https://leetcode.com/problems/palindrome-partitioning/)| * Palindrome Partitioning|40%|Medium||
|[0132](https://leetcode.com/problems/palindrome-partitioning-ii/)| * Palindrome Partitioning II|26%|Hard||
|[0134](https://leetcode.com/problems/gas-station/)| * Gas Station|33%|Medium||
|[0135](https://leetcode.com/problems/candy/)| * Candy|28%|Hard||
|[0136](https://leetcode.com/problems/single-number/)| * Single Number|59%|Easy||
|[0137](https://leetcode.com/problems/single-number-ii/)| * Single Number II|45%|Medium||
|[0139](https://leetcode.com/problems/word-break/)| * Word Break|34%|Medium||
|[0140](https://leetcode.com/problems/word-break-ii/)| * Word Break II|26%|Hard||
|[0141](https://leetcode.com/problems/linked-list-cycle/)| * Linked List Cycle|36%|Easy||
|[0142](https://leetcode.com/problems/linked-list-cycle-ii/)| * Linked List Cycle II|31%|Medium||
|[0143](https://leetcode.com/problems/reorder-list/)| * Reorder List|30%|Medium||
|[0144](https://leetcode.com/problems/binary-tree-preorder-traversal/)| * Binary Tree Preorder Traversal|50%|Medium||
|[0145](https://leetcode.com/problems/binary-tree-postorder-traversal/)| * Binary Tree Postorder Traversal|47%|Hard||
|[0146](https://leetcode.com/problems/lru-cache/)| * LRU Cache|24%|Hard||
|[0147](https://leetcode.com/problems/insertion-sort-list/)| * Insertion Sort List|36%|Medium||
|[0148](https://leetcode.com/problems/sort-list/)| * Sort List|34%|Medium||
|[0149](https://leetcode.com/problems/max-points-on-a-line/)| * Max Points on a Line|15%|Hard||
|[0150](https://leetcode.com/problems/evaluate-reverse-polish-notation/)| * Evaluate Reverse Polish Notation|31%|Medium||
|[0152](https://leetcode.com/problems/maximum-product-subarray/)| * Maximum Product Subarray|28%|Medium||
|[0153](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)| * Find Minimum in Rotated Sorted Array|42%|Medium||
|[0154](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/)| * Find Minimum in Rotated Sorted Array II|39%|Hard||
|[0155](https://leetcode.com/problems/min-stack/)| * Min Stack|36%|Easy||
|[0160](https://leetcode.com/problems/intersection-of-two-linked-lists/)| * Intersection of Two Linked Lists|33%|Easy||
|[0162](https://leetcode.com/problems/find-peak-element/)| * Find Peak Element|40%|Medium||
|[0164](https://leetcode.com/problems/maximum-gap/)| * Maximum Gap|32%|Hard||
|[0165](https://leetcode.com/problems/compare-version-numbers/)| * Compare Version Numbers|23%|Medium||
|[0166](https://leetcode.com/problems/fraction-to-recurring-decimal/)| * Fraction to Recurring Decimal|19%|Medium||
|[0167](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)| * Two Sum II - Input array is sorted|49%|Easy||
|[0168](https://leetcode.com/problems/excel-sheet-column-title/)| * Excel Sheet Column Title|28%|Easy||
|[0169](https://leetcode.com/problems/majority-element/)| * Majority Element|52%|Easy||
|[0171](https://leetcode.com/problems/excel-sheet-column-number/)| * Excel Sheet Column Number|51%|Easy||
|[0172](https://leetcode.com/problems/factorial-trailing-zeroes/)| * Factorial Trailing Zeroes|37%|Easy||
|[0173](https://leetcode.com/problems/binary-search-tree-iterator/)| * Binary Search Tree Iterator|47%|Medium||
|[0174](https://leetcode.com/problems/dungeon-game/)| * Dungeon Game|26%|Hard||
|[0179](https://leetcode.com/problems/largest-number/)| * Largest Number|25%|Medium||
|[0187](https://leetcode.com/problems/repeated-dna-sequences/)| * Repeated DNA Sequences|35%|Medium||
|[0188](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/)| * Best Time to Buy and Sell Stock IV|26%|Hard||
|[0189](https://leetcode.com/problems/rotate-array/)| * Rotate Array|29%|Easy||
|[0190](https://leetcode.com/problems/reverse-bits/)| * Reverse Bits|30%|Easy||
|[0191](https://leetcode.com/problems/number-of-1-bits/)| * Number of 1 Bits|42%|Easy||
|[0198](https://leetcode.com/problems/house-robber/)| * House Robber|40%|Easy||
|[0199](https://leetcode.com/problems/binary-tree-right-side-view/)| * Binary Tree Right Side View|47%|Medium||
|[0200](https://leetcode.com/problems/number-of-islands/)| * Number of Islands|40%|Medium||
|[0201](https://leetcode.com/problems/bitwise-and-of-numbers-range/)| * Bitwise AND of Numbers Range|35%|Medium||
|[0202](https://leetcode.com/problems/happy-number/)| * Happy Number|44%|Easy||
|[0203](https://leetcode.com/problems/remove-linked-list-elements/)| * Remove Linked List Elements|35%|Easy||
|[0204](https://leetcode.com/problems/count-primes/)| * Count Primes|28%|Easy||
|[0205](https://leetcode.com/problems/isomorphic-strings/)| * Isomorphic Strings|36%|Easy||
|[0206](https://leetcode.com/problems/reverse-linked-list/)| * Reverse Linked List|53%|Easy||
|[0207](https://leetcode.com/problems/course-schedule/)| * Course Schedule|37%|Medium||
|[0208](https://leetcode.com/problems/implement-trie-prefix-tree/)| * Implement Trie (Prefix Tree)|37%|Medium||
|[0209](https://leetcode.com/problems/minimum-size-subarray-sum/)| * Minimum Size Subarray Sum|34%|Medium||
|[0210](https://leetcode.com/problems/course-schedule-ii/)| * Course Schedule II|34%|Medium||
|[0211](https://leetcode.com/problems/add-and-search-word-data-structure-design/)| * Add and Search Word - Data structure design|29%|Medium||
|[0212](https://leetcode.com/problems/word-search-ii/)| * Word Search II|28%|Hard||
|[0213](https://leetcode.com/problems/house-robber-ii/)| * House Robber II|35%|Medium||
|[0214](https://leetcode.com/problems/shortest-palindrome/)| * Shortest Palindrome|27%|Hard||
|[0215](https://leetcode.com/problems/kth-largest-element-in-an-array/)| * Kth Largest Element in an Array|46%|Medium||
|[0216](https://leetcode.com/problems/combination-sum-iii/)| * Combination Sum III|50%|Medium||
|[0217](https://leetcode.com/problems/contains-duplicate/)| * Contains Duplicate|51%|Easy||
|[0218](https://leetcode.com/problems/the-skyline-problem/)| * The Skyline Problem|31%|Hard||
|[0219](https://leetcode.com/problems/contains-duplicate-ii/)| * Contains Duplicate II|34%|Easy||
|[0220](https://leetcode.com/problems/contains-duplicate-iii/)| * Contains Duplicate III|19%|Medium||
|[0221](https://leetcode.com/problems/maximal-square/)| * Maximal Square|32%|Medium||
|[0222](https://leetcode.com/problems/count-complete-tree-nodes/)| * Count Complete Tree Nodes|32%|Medium||
|[0223](https://leetcode.com/problems/rectangle-area/)| * Rectangle Area|35%|Medium||
|[0224](https://leetcode.com/problems/basic-calculator/)| * Basic Calculator|32%|Hard||
|[0225](https://leetcode.com/problems/implement-stack-using-queues/)| * Implement Stack using Queues|38%|Easy||
|[0226](https://leetcode.com/problems/invert-binary-tree/)| * Invert Binary Tree|57%|Easy||
|[0227](https://leetcode.com/problems/basic-calculator-ii/)| * Basic Calculator II|32%|Medium||
|[0228](https://leetcode.com/problems/summary-ranges/)| * Summary Ranges|35%|Medium||
|[0229](https://leetcode.com/problems/majority-element-ii/)| * Majority Element II|31%|Medium||
|[0230](https://leetcode.com/problems/kth-smallest-element-in-a-bst/)| * Kth Smallest Element in a BST|50%|Medium||
|[0231](https://leetcode.com/problems/power-of-two/)| * Power of Two|41%|Easy||
|[0232](https://leetcode.com/problems/implement-queue-using-stacks/)| * Implement Queue using Stacks|42%|Easy||
|[0233](https://leetcode.com/problems/number-of-digit-one/)| * Number of Digit One|30%|Hard||
|[0234](https://leetcode.com/problems/palindrome-linked-list/)| * Palindrome Linked List|35%|Easy||
|[0235](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)| * Lowest Common Ancestor of a Binary Search Tree|43%|Easy||
|[0236](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)| * Lowest Common Ancestor of a Binary Tree|36%|Medium||
|[0237](https://leetcode.com/problems/delete-node-in-a-linked-list/)| * Delete Node in a Linked List|52%|Easy||
|[0238](https://leetcode.com/problems/product-of-array-except-self/)| * Product of Array Except Self|54%|Medium||
|[0239](https://leetcode.com/problems/sliding-window-maximum/)| * Sliding Window Maximum|37%|Hard||
|[0240](https://leetcode.com/problems/search-a-2d-matrix-ii/)| * Search a 2D Matrix II|40%|Medium||
|[0241](https://leetcode.com/problems/different-ways-to-add-parentheses/)| * Different Ways to Add Parentheses|49%|Medium||
|[0242](https://leetcode.com/problems/valid-anagram/)| * Valid Anagram|51%|Easy||
|[0257](https://leetcode.com/problems/binary-tree-paths/)| * Binary Tree Paths|45%|Easy||
|[0258](https://leetcode.com/problems/add-digits/)| * Add Digits|53%|Easy||
|[0260](https://leetcode.com/problems/single-number-iii/)| * Single Number III|56%|Medium||
|[0263](https://leetcode.com/problems/ugly-number/)| * Ugly Number|40%|Easy||
|[0264](https://leetcode.com/problems/ugly-number-ii/)| * Ugly Number II|35%|Medium||
|[0268](https://leetcode.com/problems/missing-number/)| * Missing Number|47%|Easy||
|[0273](https://leetcode.com/problems/integer-to-english-words/)| * Integer to English Words|24%|Hard||
|[0274](https://leetcode.com/problems/h-index/)| * H-Index|34%|Medium||
|[0275](https://leetcode.com/problems/h-index-ii/)| * H-Index II|35%|Medium||
|[0279](https://leetcode.com/problems/perfect-squares/)| * Perfect Squares|41%|Medium||
|[0282](https://leetcode.com/problems/expression-add-operators/)| * Expression Add Operators|32%|Hard||
|[0283](https://leetcode.com/problems/move-zeroes/)| * Move Zeroes|53%|Easy||
|[0287](https://leetcode.com/problems/find-the-duplicate-number/)| * Find the Duplicate Number|48%|Medium||
|[0289](https://leetcode.com/problems/game-of-life/)| * Game of Life|44%|Medium||
|[0290](https://leetcode.com/problems/word-pattern/)| * Word Pattern|34%|Easy||
|[0292](https://leetcode.com/problems/nim-game/)| * Nim Game|55%|Easy||
|[0295](https://leetcode.com/problems/find-median-from-data-stream/)| * Find Median from Data Stream|35%|Hard||
|[0299](https://leetcode.com/problems/bulls-and-cows/)| * Bulls and Cows|38%|Medium||
|[0300](https://leetcode.com/problems/longest-increasing-subsequence/)| * Longest Increasing Subsequence|40%|Medium||
|[0301](https://leetcode.com/problems/remove-invalid-parentheses/)| * Remove Invalid Parentheses|38%|Hard||
|[0303](https://leetcode.com/problems/range-sum-query-immutable/)| * Range Sum Query - Immutable|37%|Easy||
|[0304](https://leetcode.com/problems/range-sum-query-2d-immutable/)| * Range Sum Query 2D - Immutable|31%|Medium||
|[0306](https://leetcode.com/problems/additive-number/)| * Additive Number|28%|Medium||
|[0307](https://leetcode.com/problems/range-sum-query-mutable/)| * Range Sum Query - Mutable|27%|Medium||
|[0309](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)| * Best Time to Buy and Sell Stock with Cooldown|43%|Medium||
|[0310](https://leetcode.com/problems/minimum-height-trees/)| * Minimum Height Trees|29%|Medium||
|[0312](https://leetcode.com/problems/burst-balloons/)| * Burst Balloons|46%|Hard||
|[0313](https://leetcode.com/problems/super-ugly-number/)| * Super Ugly Number|40%|Medium||
|[0315](https://leetcode.com/problems/count-of-smaller-numbers-after-self/)| * Count of Smaller Numbers After Self|37%|Hard||
|[0316](https://leetcode.com/problems/remove-duplicate-letters/)| * Remove Duplicate Letters|32%|Hard||
|[0318](https://leetcode.com/problems/maximum-product-of-word-lengths/)| * Maximum Product of Word Lengths|48%|Medium||
|[0319](https://leetcode.com/problems/bulb-switcher/)| * Bulb Switcher|43%|Medium||
|[0321](https://leetcode.com/problems/create-maximum-number/)| * Create Maximum Number|25%|Hard||
|[0322](https://leetcode.com/problems/coin-change/)| * Coin Change|29%|Medium||
|[0324](https://leetcode.com/problems/wiggle-sort-ii/)| * Wiggle Sort II|27%|Medium||
|[0326](https://leetcode.com/problems/power-of-three/)| * Power of Three|41%|Easy||
|[0327](https://leetcode.com/problems/count-of-range-sum/)| * Count of Range Sum|32%|Hard||
|[0328](https://leetcode.com/problems/odd-even-linked-list/)| * Odd Even Linked List|48%|Medium||
|[0329](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/)| * Longest Increasing Path in a Matrix|39%|Hard||
|[0330](https://leetcode.com/problems/patching-array/)| * Patching Array|33%|Hard||
|[0331](https://leetcode.com/problems/verify-preorder-serialization-of-a-binary-tree/)| * Verify Preorder Serialization of a Binary Tree|38%|Medium||
|[0332](https://leetcode.com/problems/reconstruct-itinerary/)| * Reconstruct Itinerary|30%|Medium||
|[0334](https://leetcode.com/problems/increasing-triplet-subsequence/)| * Increasing Triplet Subsequence|39%|Medium||
|[0335](https://leetcode.com/problems/self-crossing/)| * Self Crossing|26%|Hard||
|[0336](https://leetcode.com/problems/palindrome-pairs/)| * Palindrome Pairs|30%|Hard||
|[0337](https://leetcode.com/problems/house-robber-iii/)| * House Robber III|47%|Medium||
|[0338](https://leetcode.com/problems/counting-bits/)| * Counting Bits|64%|Medium||
|[0342](https://leetcode.com/problems/power-of-four/)| * Power of Four|40%|Easy||
|[0343](https://leetcode.com/problems/integer-break/)| * Integer Break|47%|Medium||
|[0344](https://leetcode.com/problems/reverse-string/)| * Reverse String|63%|Easy||
|[0345](https://leetcode.com/problems/reverse-vowels-of-a-string/)| * Reverse Vowels of a String|41%|Easy||
|[0347](https://leetcode.com/problems/top-k-frequent-elements/)| * Top K Frequent Elements|54%|Medium||
|[0349](https://leetcode.com/problems/intersection-of-two-arrays/)| * Intersection of Two Arrays|53%|Easy||
|[0350](https://leetcode.com/problems/intersection-of-two-arrays-ii/)| * Intersection of Two Arrays II|47%|Easy||
|[0352](https://leetcode.com/problems/data-stream-as-disjoint-intervals/)| * Data Stream as Disjoint Intervals|43%|Hard||
|[0354](https://leetcode.com/problems/russian-doll-envelopes/)| * Russian Doll Envelopes|33%|Hard||
|[0355](https://leetcode.com/problems/design-twitter/)| * Design Twitter|27%|Medium||
|[0357](https://leetcode.com/problems/count-numbers-with-unique-digits/)| * Count Numbers with Unique Digits|46%|Medium||
|[0363](https://leetcode.com/problems/max-sum-of-rectangle-no-larger-than-k/)| * Max Sum of Rectangle No Larger Than K|34%|Hard||
|[0365](https://leetcode.com/problems/water-and-jug-problem/)| * Water and Jug Problem|28%|Medium||
|[0367](https://leetcode.com/problems/valid-perfect-square/)| * Valid Perfect Square|39%|Easy||
|[0368](https://leetcode.com/problems/largest-divisible-subset/)| * Largest Divisible Subset|34%|Medium||
|[0371](https://leetcode.com/problems/sum-of-two-integers/)| * Sum of Two Integers|51%|Easy||
|[0372](https://leetcode.com/problems/super-pow/)| * Super Pow|35%|Medium||
|[0373](https://leetcode.com/problems/find-k-pairs-with-smallest-sums/)| * Find K Pairs with Smallest Sums|33%|Medium||
|[0375](https://leetcode.com/problems/guess-number-higher-or-lower-ii/)| * Guess Number Higher or Lower II|37%|Medium||
|[0376](https://leetcode.com/problems/wiggle-subsequence/)| * Wiggle Subsequence|37%|Medium||
|[0377](https://leetcode.com/problems/combination-sum-iv/)| * Combination Sum IV|43%|Medium||
|[0378](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/)| * Kth Smallest Element in a Sorted Matrix|48%|Medium||
|[0380](https://leetcode.com/problems/insert-delete-getrandom-o1/)| * Insert Delete GetRandom O(1)|42%|Medium||
|[0381](https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/)| * Insert Delete GetRandom O(1) - Duplicates allowed|31%|Hard||
|[0382](https://leetcode.com/problems/linked-list-random-node/)| * Linked List Random Node|48%|Medium||
|[0383](https://leetcode.com/problems/ransom-note/)| * Ransom Note|49%|Easy||
|[0384](https://leetcode.com/problems/shuffle-an-array/)| * Shuffle an Array|49%|Medium||
|[0385](https://leetcode.com/problems/mini-parser/)| * Mini Parser|31%|Medium||
|[0386](https://leetcode.com/problems/lexicographical-numbers/)| * Lexicographical Numbers|45%|Medium||
|[0387](https://leetcode.com/problems/first-unique-character-in-a-string/)| * First Unique Character in a String|49%|Easy||
|[0388](https://leetcode.com/problems/longest-absolute-file-path/)| * Longest Absolute File Path|38%|Medium||
|[0389](https://leetcode.com/problems/find-the-difference/)| * Find the Difference|52%|Easy||
|[0390](https://leetcode.com/problems/elimination-game/)| * Elimination Game|43%|Medium||
|[0391](https://leetcode.com/problems/perfect-rectangle/)| * Perfect Rectangle|28%|Hard||
|[0392](https://leetcode.com/problems/is-subsequence/)| * Is Subsequence|46%|Medium||
|[0393](https://leetcode.com/problems/utf-8-validation/)| * UTF-8 Validation|35%|Medium||
|[0394](https://leetcode.com/problems/decode-string/)| * Decode String|44%|Medium||
|[0395](https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/)| * Longest Substring with At Least K Repeating Characters|38%|Medium||
|[0396](https://leetcode.com/problems/rotate-function/)| * Rotate Function|34%|Medium||
|[0397](https://leetcode.com/problems/integer-replacement/)| * Integer Replacement|31%|Medium||
|[0398](https://leetcode.com/problems/random-pick-index/)| * Random Pick Index|49%|Medium||
|[0399](https://leetcode.com/problems/evaluate-division/)| * Evaluate Division|47%|Medium||
|[0400](https://leetcode.com/problems/nth-digit/)| * Nth Digit|30%|Easy||
|[0401](https://leetcode.com/problems/binary-watch/)| * Binary Watch|45%|Easy||
|[0402](https://leetcode.com/problems/remove-k-digits/)| * Remove K Digits|26%|Medium||
|[0403](https://leetcode.com/problems/frog-jump/)| * Frog Jump|35%|Hard||
|[0404](https://leetcode.com/problems/sum-of-left-leaves/)| * Sum of Left Leaves|48%|Easy||
|[0405](https://leetcode.com/problems/convert-a-number-to-hexadecimal/)| * Convert a Number to Hexadecimal|41%|Easy||
|[0406](https://leetcode.com/problems/queue-reconstruction-by-height/)| * Queue Reconstruction by Height|59%|Medium||
|[0407](https://leetcode.com/problems/trapping-rain-water-ii/)| * Trapping Rain Water II|38%|Hard||
|[0409](https://leetcode.com/problems/longest-palindrome/)| * Longest Palindrome|47%|Easy||
|[0410](https://leetcode.com/problems/split-array-largest-sum/)| * Split Array Largest Sum|42%|Hard||
|[0412](https://leetcode.com/problems/fizz-buzz/)| * Fizz Buzz|59%|Easy||
|[0413](https://leetcode.com/problems/arithmetic-slices/)| * Arithmetic Slices|55%|Medium||
|[0414](https://leetcode.com/problems/third-maximum-number/)| * Third Maximum Number|28%|Easy||
|[0415](https://leetcode.com/problems/add-strings/)| * Add Strings|43%|Easy||
|[0416](https://leetcode.com/problems/partition-equal-subset-sum/)| * Partition Equal Subset Sum|40%|Medium||
|[0417](https://leetcode.com/problems/pacific-atlantic-water-flow/)| * Pacific Atlantic Water Flow|37%|Medium||
|[0419](https://leetcode.com/problems/battleships-in-a-board/)| * Battleships in a Board|65%|Medium||
|[0420](https://leetcode.com/problems/strong-password-checker/)| * Strong Password Checker|17%|Hard||
|[0421](https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/)| * Maximum XOR of Two Numbers in an Array|50%|Medium||
|[0423](https://leetcode.com/problems/reconstruct-original-digits-from-english/)| * Reconstruct Original Digits from English|45%|Medium||
|[0424](https://leetcode.com/problems/longest-repeating-character-replacement/)| * Longest Repeating Character Replacement|43%|Medium||
|[0432](https://leetcode.com/problems/all-oone-data-structure/)| * All O`one Data Structure|29%|Hard||
|[0433](https://leetcode.com/problems/minimum-genetic-mutation/)| * Minimum Genetic Mutation|37%|Medium||
|[0434](https://leetcode.com/problems/number-of-segments-in-a-string/)| * Number of Segments in a String|36%|Easy||
|[0435](https://leetcode.com/problems/non-overlapping-intervals/)| * Non-overlapping Intervals|41%|Medium||
|[0436](https://leetcode.com/problems/find-right-interval/)| * Find Right Interval|42%|Medium||
|[0437](https://leetcode.com/problems/path-sum-iii/)| * Path Sum III|42%|Easy||
|[0438](https://leetcode.com/problems/find-all-anagrams-in-a-string/)| * Find All Anagrams in a String|36%|Easy||
|[0440](https://leetcode.com/problems/k-th-smallest-in-lexicographical-order/)| * K-th Smallest in Lexicographical Order|26%|Hard||
|[0441](https://leetcode.com/problems/arranging-coins/)| * Arranging Coins|37%|Easy||
|[0442](https://leetcode.com/problems/find-all-duplicates-in-an-array/)| * Find All Duplicates in an Array|60%|Medium||
|[0443](https://leetcode.com/problems/string-compression/)| * String Compression|37%|Easy||
|[0445](https://leetcode.com/problems/add-two-numbers-ii/)| * Add Two Numbers II|49%|Medium||
|[0446](https://leetcode.com/problems/arithmetic-slices-ii-subsequence/)| * Arithmetic Slices II - Subsequence|29%|Hard||
|[0447](https://leetcode.com/problems/number-of-boomerangs/)| * Number of Boomerangs|49%|Easy||
|[0448](https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/)| * Find All Numbers Disappeared in an Array|52%|Easy||
|[0450](https://leetcode.com/problems/delete-node-in-a-bst/)| * Delete Node in a BST|39%|Medium||
|[0451](https://leetcode.com/problems/sort-characters-by-frequency/)| * Sort Characters By Frequency|55%|Medium||
|[0452](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)| * Minimum Number of Arrows to Burst Balloons|45%|Medium||
|[0453](https://leetcode.com/problems/minimum-moves-to-equal-array-elements/)| * Minimum Moves to Equal Array Elements|49%|Easy||
|[0454](https://leetcode.com/problems/4sum-ii/)| * 4Sum II|50%|Medium||
|[0455](https://leetcode.com/problems/assign-cookies/)| * Assign Cookies|48%|Easy||
|[0456](https://leetcode.com/problems/132-pattern/)| * 132 Pattern|27%|Medium||
|[0457](https://leetcode.com/problems/circular-array-loop/)| * Circular Array Loop|27%|Medium||
|[0458](https://leetcode.com/problems/poor-pigs/)| * Poor Pigs|45%|Hard||
|[0459](https://leetcode.com/problems/repeated-substring-pattern/)| * Repeated Substring Pattern|39%|Easy||
|[0460](https://leetcode.com/problems/lfu-cache/)| * LFU Cache|28%|Hard||
|[0461](https://leetcode.com/problems/hamming-distance/)| * Hamming Distance|70%|Easy||
|[0462](https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/)| * Minimum Moves to Equal Array Elements II|52%|Medium||
|[0463](https://leetcode.com/problems/island-perimeter/)| * Island Perimeter|60%|Easy||
|[0464](https://leetcode.com/problems/can-i-win/)| * Can I Win|27%|Medium||
|[0466](https://leetcode.com/problems/count-the-repetitions/)| * Count The Repetitions|27%|Hard||
|[0467](https://leetcode.com/problems/unique-substrings-in-wraparound-string/)| * Unique Substrings in Wraparound String|33%|Medium||
|[0468](https://leetcode.com/problems/validate-ip-address/)| * Validate IP Address|21%|Medium||
|[0470](https://leetcode.com/problems/implement-rand10-using-rand7/)| * Implement Rand10() Using Rand7()|44%|Medium||
|[0472](https://leetcode.com/problems/concatenated-words/)| * Concatenated Words|33%|Hard||
|[0473](https://leetcode.com/problems/matchsticks-to-square/)| * Matchsticks to Square|35%|Medium||
|[0474](https://leetcode.com/problems/ones-and-zeroes/)| * Ones and Zeroes|39%|Medium||
|[0475](https://leetcode.com/problems/heaters/)| * Heaters|31%|Easy||
|[0476](https://leetcode.com/problems/number-complement/)| * Number Complement|62%|Easy||
|[0477](https://leetcode.com/problems/total-hamming-distance/)| * Total Hamming Distance|48%|Medium||
|[0478](https://leetcode.com/problems/generate-random-point-in-a-circle/)| * Generate Random Point in a Circle|36%|Medium||
|[0479](https://leetcode.com/problems/largest-palindrome-product/)| * Largest Palindrome Product|27%|Hard||
|[0480](https://leetcode.com/problems/sliding-window-median/)| * Sliding Window Median|31%|Hard||
|[0481](https://leetcode.com/problems/magical-string/)| * Magical String|45%|Medium||
|[0482](https://leetcode.com/problems/license-key-formatting/)| * License Key Formatting|40%|Easy||
|[0483](https://leetcode.com/problems/smallest-good-base/)| * Smallest Good Base|34%|Hard||
|[0485](https://leetcode.com/problems/max-consecutive-ones/)| * Max Consecutive Ones|54%|Easy||
|[0486](https://leetcode.com/problems/predict-the-winner/)| * Predict the Winner|46%|Medium||
|[0488](https://leetcode.com/problems/zuma-game/)| * Zuma Game|38%|Hard||
|[0491](https://leetcode.com/problems/increasing-subsequences/)| * Increasing Subsequences|41%|Medium||
|[0492](https://leetcode.com/problems/construct-the-rectangle/)| * Construct the Rectangle|48%|Easy||
|[0493](https://leetcode.com/problems/reverse-pairs/)| * Reverse Pairs|22%|Hard||
|[0494](https://leetcode.com/problems/target-sum/)| * Target Sum|44%|Medium||
|[0495](https://leetcode.com/problems/teemo-attacking/)| * Teemo Attacking|52%|Medium||
|[0496](https://leetcode.com/problems/next-greater-element-i/)| * Next Greater Element I|59%|Easy||
|[0497](https://leetcode.com/problems/random-point-in-non-overlapping-rectangles/)| * Random Point in Non-overlapping Rectangles|35%|Medium||
|[0498](https://leetcode.com/problems/diagonal-traverse/)| * Diagonal Traverse|45%|Medium||
|[0500](https://leetcode.com/problems/keyboard-row/)| * Keyboard Row|61%|Easy||
|[0501](https://leetcode.com/problems/find-mode-in-binary-search-tree/)| * Find Mode in Binary Search Tree|39%|Easy||
|[0502](https://leetcode.com/problems/ipo/)| * IPO|37%|Hard||
|[0503](https://leetcode.com/problems/next-greater-element-ii/)| * Next Greater Element II|50%|Medium||
|[0504](https://leetcode.com/problems/base-7/)| * Base 7|44%|Easy||
|[0506](https://leetcode.com/problems/relative-ranks/)| * Relative Ranks|48%|Easy||
|[0507](https://leetcode.com/problems/perfect-number/)| * Perfect Number|33%|Easy||
|[0508](https://leetcode.com/problems/most-frequent-subtree-sum/)| * Most Frequent Subtree Sum|54%|Medium||
|[0509](https://leetcode.com/problems/fibonacci-number/)| * Fibonacci Number|66%|Easy||
|[0513](https://leetcode.com/problems/find-bottom-left-tree-value/)| * Find Bottom Left Tree Value|58%|Medium||
|[0514](https://leetcode.com/problems/freedom-trail/)| * Freedom Trail|40%|Hard||
|[0515](https://leetcode.com/problems/find-largest-value-in-each-tree-row/)| * Find Largest Value in Each Tree Row|57%|Medium||
|[0516](https://leetcode.com/problems/longest-palindromic-subsequence/)| * Longest Palindromic Subsequence|45%|Medium||
|[0517](https://leetcode.com/problems/super-washing-machines/)| * Super Washing Machines|36%|Hard||
|[0518](https://leetcode.com/problems/coin-change-2/)| * Coin Change 2|42%|Medium||
|[0519](https://leetcode.com/problems/random-flip-matrix/)| * Random Flip Matrix|32%|Medium||
|[0520](https://leetcode.com/problems/detect-capital/)| * Detect Capital|52%|Easy||
|[0521](https://leetcode.com/problems/longest-uncommon-subsequence-i/)| * Longest Uncommon Subsequence I |56%|Easy||
|[0522](https://leetcode.com/problems/longest-uncommon-subsequence-ii/)| * Longest Uncommon Subsequence II|32%|Medium||
|[0523](https://leetcode.com/problems/continuous-subarray-sum/)| * Continuous Subarray Sum|24%|Medium||
|[0524](https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/)| * Longest Word in Dictionary through Deleting|45%|Medium||
|[0525](https://leetcode.com/problems/contiguous-array/)| * Contiguous Array|42%|Medium||
|[0526](https://leetcode.com/problems/beautiful-arrangement/)| * Beautiful Arrangement|54%|Medium||
|[0528](https://leetcode.com/problems/random-pick-with-weight/)| * Random Pick with Weight|42%|Medium||
|[0529](https://leetcode.com/problems/minesweeper/)| * Minesweeper|52%|Medium||
|[0530](https://leetcode.com/problems/minimum-absolute-difference-in-bst/)| * Minimum Absolute Difference in BST|50%|Easy||
|[0532](https://leetcode.com/problems/k-diff-pairs-in-an-array/)| * K-diff Pairs in an Array|29%|Easy||
|[0537](https://leetcode.com/problems/complex-number-multiplication/)| * Complex Number Multiplication|65%|Medium||
|[0538](https://leetcode.com/problems/convert-bst-to-greater-tree/)| * Convert BST to Greater Tree|50%|Easy||
|[0539](https://leetcode.com/problems/minimum-time-difference/)| * Minimum Time Difference|47%|Medium||
|[0540](https://leetcode.com/problems/single-element-in-a-sorted-array/)| * Single Element in a Sorted Array|57%|Medium||
|[0541](https://leetcode.com/problems/reverse-string-ii/)| * Reverse String II|45%|Easy||
|[0542](https://leetcode.com/problems/01-matrix/)| * 01 Matrix|35%|Medium||
|[0543](https://leetcode.com/problems/diameter-of-binary-tree/)| * Diameter of Binary Tree|46%|Easy||
|[0546](https://leetcode.com/problems/remove-boxes/)| * Remove Boxes|37%|Hard||
|[0547](https://leetcode.com/problems/friend-circles/)| * Friend Circles|53%|Medium||
|[0551](https://leetcode.com/problems/student-attendance-record-i/)| * Student Attendance Record I|45%|Easy||
|[0552](https://leetcode.com/problems/student-attendance-record-ii/)| * Student Attendance Record II|32%|Hard||
|[0553](https://leetcode.com/problems/optimal-division/)| * Optimal Division|55%|Medium||
|[0554](https://leetcode.com/problems/brick-wall/)| * Brick Wall|47%|Medium||
|[0556](https://leetcode.com/problems/next-greater-element-iii/)| * Next Greater Element III|29%|Medium||
|[0557](https://leetcode.com/problems/reverse-words-in-a-string-iii/)| * Reverse Words in a String III|63%|Easy||
|[0560](https://leetcode.com/problems/subarray-sum-equals-k/)| * Subarray Sum Equals K|41%|Medium||
|[0561](https://leetcode.com/problems/array-partition-i/)| * Array Partition I|68%|Easy||
|[0563](https://leetcode.com/problems/binary-tree-tilt/)| * Binary Tree Tilt|46%|Easy||
|[0564](https://leetcode.com/problems/find-the-closest-palindrome/)| * Find the Closest Palindrome|18%|Hard||
|[0565](https://leetcode.com/problems/array-nesting/)| * Array Nesting|52%|Medium||
|[0566](https://leetcode.com/problems/reshape-the-matrix/)| * Reshape the Matrix|58%|Easy||
|[0567](https://leetcode.com/problems/permutation-in-string/)| * Permutation in String|38%|Medium||
|[0572](https://leetcode.com/problems/subtree-of-another-tree/)| * Subtree of Another Tree|41%|Easy||
|[0575](https://leetcode.com/problems/distribute-candies/)| * Distribute Candies|59%|Easy||
|[0576](https://leetcode.com/problems/out-of-boundary-paths/)| * Out of Boundary Paths|31%|Medium||
|[0581](https://leetcode.com/problems/shortest-unsorted-continuous-subarray/)| * Shortest Unsorted Continuous Subarray|29%|Easy||
|[0583](https://leetcode.com/problems/delete-operation-for-two-strings/)| * Delete Operation for Two Strings|44%|Medium||
|[0587](https://leetcode.com/problems/erect-the-fence/)| * Erect the Fence|34%|Hard||
|[0591](https://leetcode.com/problems/tag-validator/)| * Tag Validator|31%|Hard||
|[0592](https://leetcode.com/problems/fraction-addition-and-subtraction/)| * Fraction Addition and Subtraction|46%|Medium||
|[0593](https://leetcode.com/problems/valid-square/)| * Valid Square|40%|Medium||
|[0594](https://leetcode.com/problems/longest-harmonious-subsequence/)| * Longest Harmonious Subsequence|43%|Easy||
|[0598](https://leetcode.com/problems/range-addition-ii/)| * Range Addition II|48%|Easy||
|[0599](https://leetcode.com/problems/minimum-index-sum-of-two-lists/)| * Minimum Index Sum of Two Lists|47%|Easy||
|[0600](https://leetcode.com/problems/non-negative-integers-without-consecutive-ones/)| * Non-negative Integers without Consecutive Ones|32%|Hard||
|[0605](https://leetcode.com/problems/can-place-flowers/)| * Can Place Flowers|30%|Easy||
|[0606](https://leetcode.com/problems/construct-string-from-binary-tree/)| * Construct String from Binary Tree|51%|Easy||
|[0609](https://leetcode.com/problems/find-duplicate-file-in-system/)| * Find Duplicate File in System|54%|Medium||
|[0611](https://leetcode.com/problems/valid-triangle-number/)| * Valid Triangle Number|44%|Medium||
|[0617](https://leetcode.com/problems/merge-two-binary-trees/)| * Merge Two Binary Trees|69%|Easy||
|[0621](https://leetcode.com/problems/task-scheduler/)| * Task Scheduler|44%|Medium||
|[0622](https://leetcode.com/problems/design-circular-queue/)| * Design Circular Queue|38%|Medium||
|[0623](https://leetcode.com/problems/add-one-row-to-tree/)| * Add One Row to Tree|47%|Medium||
|[0628](https://leetcode.com/problems/maximum-product-of-three-numbers/)| * Maximum Product of Three Numbers|45%|Easy||
|[0629](https://leetcode.com/problems/k-inverse-pairs-array/)| * K Inverse Pairs Array|29%|Hard||
|[0630](https://leetcode.com/problems/course-schedule-iii/)| * Course Schedule III|31%|Hard||
|[0632](https://leetcode.com/problems/smallest-range/)| * Smallest Range|46%|Hard||
|[0633](https://leetcode.com/problems/sum-of-square-numbers/)| * Sum of Square Numbers|32%|Easy||
|[0636](https://leetcode.com/problems/exclusive-time-of-functions/)| * Exclusive Time of Functions|48%|Medium||
|[0637](https://leetcode.com/problems/average-of-levels-in-binary-tree/)| * Average of Levels in Binary Tree|58%|Easy||
|[0638](https://leetcode.com/problems/shopping-offers/)| * Shopping Offers|48%|Medium||
|[0639](https://leetcode.com/problems/decode-ways-ii/)| * Decode Ways II|24%|Hard||
|[0640](https://leetcode.com/problems/solve-the-equation/)| * Solve the Equation|40%|Medium||
|[0641](https://leetcode.com/problems/design-circular-deque/)| * Design Circular Deque|48%|Medium||
|[0643](https://leetcode.com/problems/maximum-average-subarray-i/)| * Maximum Average Subarray I|39%|Easy||
|[0645](https://leetcode.com/problems/set-mismatch/)| * Set Mismatch|40%|Easy||
|[0646](https://leetcode.com/problems/maximum-length-of-pair-chain/)| * Maximum Length of Pair Chain|48%|Medium||
|[0647](https://leetcode.com/problems/palindromic-substrings/)| * Palindromic Substrings|56%|Medium||
|[0648](https://leetcode.com/problems/replace-words/)| * Replace Words|51%|Medium||
|[0649](https://leetcode.com/problems/dota2-senate/)| * Dota2 Senate|37%|Medium||
|[0650](https://leetcode.com/problems/2-keys-keyboard/)| * 2 Keys Keyboard|46%|Medium||
|[0652](https://leetcode.com/problems/find-duplicate-subtrees/)| * Find Duplicate Subtrees|44%|Medium||
|[0653](https://leetcode.com/problems/two-sum-iv-input-is-a-bst/)| * Two Sum IV - Input is a BST|52%|Easy||
|[0654](https://leetcode.com/problems/maximum-binary-tree/)| * Maximum Binary Tree|75%|Medium||
|[0655](https://leetcode.com/problems/print-binary-tree/)| * Print Binary Tree|51%|Medium||
|[0657](https://leetcode.com/problems/robot-return-to-origin/)| * Robot Return to Origin|71%|Easy||
|[0658](https://leetcode.com/problems/find-k-closest-elements/)| * Find K Closest Elements|37%|Medium||
|[0659](https://leetcode.com/problems/split-array-into-consecutive-subsequences/)| * Split Array into Consecutive Subsequences|40%|Medium||
|[0661](https://leetcode.com/problems/image-smoother/)| * Image Smoother|48%|Easy||
|[0662](https://leetcode.com/problems/maximum-width-of-binary-tree/)| * Maximum Width of Binary Tree|39%|Medium||
|[0664](https://leetcode.com/problems/strange-printer/)| * Strange Printer|35%|Hard||
|[0665](https://leetcode.com/problems/non-decreasing-array/)| * Non-decreasing Array|19%|Easy||
|[0667](https://leetcode.com/problems/beautiful-arrangement-ii/)| * Beautiful Arrangement II|51%|Medium||
|[0668](https://leetcode.com/problems/kth-smallest-number-in-multiplication-table/)| * Kth Smallest Number in Multiplication Table|41%|Hard||
|[0669](https://leetcode.com/problems/trim-a-binary-search-tree/)| * Trim a Binary Search Tree|60%|Easy||
|[0670](https://leetcode.com/problems/maximum-swap/)| * Maximum Swap|39%|Medium||
|[0671](https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/)| * Second Minimum Node In a Binary Tree|43%|Easy||
|[0672](https://leetcode.com/problems/bulb-switcher-ii/)| * Bulb Switcher II|49%|Medium||
|[0673](https://leetcode.com/problems/number-of-longest-increasing-subsequence/)| * Number of Longest Increasing Subsequence|33%|Medium||
|[0674](https://leetcode.com/problems/longest-continuous-increasing-subsequence/)| * Longest Continuous Increasing Subsequence|44%|Easy||
|[0675](https://leetcode.com/problems/cut-off-trees-for-golf-event/)| * Cut Off Trees for Golf Event|30%|Hard||
|[0676](https://leetcode.com/problems/implement-magic-dictionary/)| * Implement Magic Dictionary|51%|Medium||
|[0677](https://leetcode.com/problems/map-sum-pairs/)| * Map Sum Pairs|51%|Medium||
|[0678](https://leetcode.com/problems/valid-parenthesis-string/)| * Valid Parenthesis String|32%|Medium||
|[0679](https://leetcode.com/problems/24-game/)| * 24 Game|42%|Hard||
|[0680](https://leetcode.com/problems/valid-palindrome-ii/)| * Valid Palindrome II|33%|Easy||
|[0682](https://leetcode.com/problems/baseball-game/)| * Baseball Game|60%|Easy||
|[0684](https://leetcode.com/problems/redundant-connection/)| * Redundant Connection|50%|Medium||
|[0685](https://leetcode.com/problems/redundant-connection-ii/)| * Redundant Connection II|30%|Hard||
|[0686](https://leetcode.com/problems/repeated-string-match/)| * Repeated String Match|31%|Easy||
|[0687](https://leetcode.com/problems/longest-univalue-path/)| * Longest Univalue Path|33%|Easy||
|[0688](https://leetcode.com/problems/knight-probability-in-chessboard/)| * Knight Probability in Chessboard|43%|Medium||
|[0689](https://leetcode.com/problems/maximum-sum-of-3-non-overlapping-subarrays/)| * Maximum Sum of 3 Non-Overlapping Subarrays|43%|Hard||
|[0691](https://leetcode.com/problems/stickers-to-spell-word/)| * Stickers to Spell Word|37%|Hard||
|[0692](https://leetcode.com/problems/top-k-frequent-words/)| * Top K Frequent Words|45%|Medium||
|[0693](https://leetcode.com/problems/binary-number-with-alternating-bits/)| * Binary Number with Alternating Bits|57%|Easy||
|[0695](https://leetcode.com/problems/max-area-of-island/)| * Max Area of Island|56%|Medium||
|[0696](https://leetcode.com/problems/count-binary-substrings/)| * Count Binary Substrings|52%|Easy||
|[0697](https://leetcode.com/problems/degree-of-an-array/)| * Degree of an Array|49%|Easy||
|[0698](https://leetcode.com/problems/partition-to-k-equal-sum-subsets/)| * Partition to K Equal Sum Subsets|41%|Medium||
|[0699](https://leetcode.com/problems/falling-squares/)| * Falling Squares|39%|Hard||
|[0700](https://leetcode.com/problems/search-in-a-binary-search-tree/)| * Search in a Binary Search Tree|67%|Easy||
|[0701](https://leetcode.com/problems/insert-into-a-binary-search-tree/)| * Insert into a Binary Search Tree|74%|Medium||
|[0703](https://leetcode.com/problems/kth-largest-element-in-a-stream/)| * Kth Largest Element in a Stream|45%|Easy||
|[0704](https://leetcode.com/problems/binary-search/)| * Binary Search|46%|Easy||
|[0705](https://leetcode.com/problems/design-hashset/)| * Design HashSet|51%|Easy||
|[0706](https://leetcode.com/problems/design-hashmap/)| * Design HashMap|55%|Easy||
|[0707](https://leetcode.com/problems/design-linked-list/)| * Design Linked List|23%|Easy||
|[0709](https://leetcode.com/problems/to-lower-case/)| * To Lower Case|76%|Easy||
|[0710](https://leetcode.com/problems/random-pick-with-blacklist/)| * Random Pick with Blacklist|30%|Hard||
|[0712](https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/)| * Minimum ASCII Delete Sum for Two Strings|53%|Medium||
|[0713](https://leetcode.com/problems/subarray-product-less-than-k/)| * Subarray Product Less Than K|36%|Medium||
|[0714](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)| * Best Time to Buy and Sell Stock with Transaction Fee|49%|Medium||
|[0715](https://leetcode.com/problems/range-module/)| * Range Module|34%|Hard||
|[0717](https://leetcode.com/problems/1-bit-and-2-bit-characters/)| * 1-bit and 2-bit Characters|49%|Easy||
|[0718](https://leetcode.com/problems/maximum-length-of-repeated-subarray/)| * Maximum Length of Repeated Subarray|45%|Medium||
|[0719](https://leetcode.com/problems/find-k-th-smallest-pair-distance/)| * Find K-th Smallest Pair Distance|28%|Hard||
|[0720](https://leetcode.com/problems/longest-word-in-dictionary/)| * Longest Word in Dictionary|44%|Easy||
|[0721](https://leetcode.com/problems/accounts-merge/)| * Accounts Merge|39%|Medium||
|[0722](https://leetcode.com/problems/remove-comments/)| * Remove Comments|30%|Medium||
|[0724](https://leetcode.com/problems/find-pivot-index/)| * Find Pivot Index|40%|Easy||
|[0725](https://leetcode.com/problems/split-linked-list-in-parts/)| * Split Linked List in Parts|48%|Medium||
|[0726](https://leetcode.com/problems/number-of-atoms/)| * Number of Atoms|44%|Hard||
|[0728](https://leetcode.com/problems/self-dividing-numbers/)| * Self Dividing Numbers|69%|Easy||
|[0729](https://leetcode.com/problems/my-calendar-i/)| * My Calendar I|46%|Medium||
|[0730](https://leetcode.com/problems/count-different-palindromic-subsequences/)| * Count Different Palindromic Subsequences|38%|Hard||
|[0731](https://leetcode.com/problems/my-calendar-ii/)| * My Calendar II|43%|Medium||
|[0732](https://leetcode.com/problems/my-calendar-iii/)| * My Calendar III|53%|Hard||
|[0733](https://leetcode.com/problems/flood-fill/)| * Flood Fill|50%|Easy||
|[0735](https://leetcode.com/problems/asteroid-collision/)| * Asteroid Collision|38%|Medium||
|[0736](https://leetcode.com/problems/parse-lisp-expression/)| * Parse Lisp Expression|43%|Hard||
|[0738](https://leetcode.com/problems/monotone-increasing-digits/)| * Monotone Increasing Digits|41%|Medium||
|[0739](https://leetcode.com/problems/daily-temperatures/)| * Daily Temperatures|59%|Medium||
|[0740](https://leetcode.com/problems/delete-and-earn/)| * Delete and Earn|45%|Medium||
|[0741](https://leetcode.com/problems/cherry-pickup/)| * Cherry Pickup|28%|Hard||
|[0743](https://leetcode.com/problems/network-delay-time/)| * Network Delay Time|40%|Medium||
|[0744](https://leetcode.com/problems/find-smallest-letter-greater-than-target/)| * Find Smallest Letter Greater Than Target|43%|Easy||
|[0745](https://leetcode.com/problems/prefix-and-suffix-search/)| * Prefix and Suffix Search|29%|Hard||
|[0746](https://leetcode.com/problems/min-cost-climbing-stairs/)| * Min Cost Climbing Stairs|46%|Easy||
|[0747](https://leetcode.com/problems/largest-number-at-least-twice-of-others/)| * Largest Number At Least Twice of Others|40%|Easy||
|[0748](https://leetcode.com/problems/shortest-completing-word/)| * Shortest Completing Word|53%|Easy||
|[0749](https://leetcode.com/problems/contain-virus/)| * Contain Virus|40%|Hard||
|[0752](https://leetcode.com/problems/open-the-lock/)| * Open the Lock|45%|Medium||
|[0753](https://leetcode.com/problems/cracking-the-safe/)| * Cracking the Safe|45%|Hard||
|[0754](https://leetcode.com/problems/reach-a-number/)| * Reach a Number|31%|Easy||
|[0756](https://leetcode.com/problems/pyramid-transition-matrix/)| * Pyramid Transition Matrix|50%|Medium||
|[0757](https://leetcode.com/problems/set-intersection-size-at-least-two/)| * Set Intersection Size At Least Two|36%|Hard||
|[0761](https://leetcode.com/problems/special-binary-string/)| * Special Binary String|50%|Hard||
|[0762](https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/)| * Prime Number of Set Bits in Binary Representation|58%|Easy||
|[0763](https://leetcode.com/problems/partition-labels/)| * Partition Labels|69%|Medium||
|[0764](https://leetcode.com/problems/largest-plus-sign/)| * Largest Plus Sign|43%|Medium||
|[0765](https://leetcode.com/problems/couples-holding-hands/)| * Couples Holding Hands|51%|Hard||
|[0766](https://leetcode.com/problems/toeplitz-matrix/)| * Toeplitz Matrix|61%|Easy||
|[0767](https://leetcode.com/problems/reorganize-string/)| * Reorganize String|41%|Medium||
|[0768](https://leetcode.com/problems/max-chunks-to-make-sorted-ii/)| * Max Chunks To Make Sorted II|45%|Hard||
|[0769](https://leetcode.com/problems/max-chunks-to-make-sorted/)| * Max Chunks To Make Sorted|51%|Medium||
|[0770](https://leetcode.com/problems/basic-calculator-iv/)| * Basic Calculator IV|44%|Hard||
|[0771](https://leetcode.com/problems/jewels-and-stones/)| * Jewels and Stones|82%|Easy||
|[0773](https://leetcode.com/problems/sliding-puzzle/)| * Sliding Puzzle|51%|Hard||
|[0775](https://leetcode.com/problems/global-and-local-inversions/)| * Global and Local Inversions|38%|Medium||
|[0777](https://leetcode.com/problems/swap-adjacent-in-lr-string/)| * Swap Adjacent in LR String|32%|Medium||
|[0778](https://leetcode.com/problems/swim-in-rising-water/)| * Swim in Rising Water|46%|Hard||
|[0779](https://leetcode.com/problems/k-th-symbol-in-grammar/)| * K-th Symbol in Grammar|37%|Medium||
|[0780](https://leetcode.com/problems/reaching-points/)| * Reaching Points|27%|Hard||
|[0781](https://leetcode.com/problems/rabbits-in-forest/)| * Rabbits in Forest|51%|Medium||
|[0782](https://leetcode.com/problems/transform-to-chessboard/)| * Transform to Chessboard|39%|Hard||
|[0783](https://leetcode.com/problems/minimum-distance-between-bst-nodes/)| * Minimum Distance Between BST Nodes|50%|Easy||
|[0784](https://leetcode.com/problems/letter-case-permutation/)| * Letter Case Permutation|55%|Easy||
|[0785](https://leetcode.com/problems/is-graph-bipartite/)| * Is Graph Bipartite?|42%|Medium||
|[0786](https://leetcode.com/problems/k-th-smallest-prime-fraction/)| * K-th Smallest Prime Fraction|39%|Hard||
|[0787](https://leetcode.com/problems/cheapest-flights-within-k-stops/)| * Cheapest Flights Within K Stops|34%|Medium||
|[0788](https://leetcode.com/problems/rotated-digits/)| * Rotated Digits|53%|Easy||
|[0789](https://leetcode.com/problems/escape-the-ghosts/)| * Escape The Ghosts|54%|Medium||
|[0790](https://leetcode.com/problems/domino-and-tromino-tiling/)| * Domino and Tromino Tiling|35%|Medium||
|[0791](https://leetcode.com/problems/custom-sort-string/)| * Custom Sort String|61%|Medium||
|[0792](https://leetcode.com/problems/number-of-matching-subsequences/)| * Number of Matching Subsequences|42%|Medium||
|[0793](https://leetcode.com/problems/preimage-size-of-factorial-zeroes-function/)| * Preimage Size of Factorial Zeroes Function|39%|Hard||
|[0794](https://leetcode.com/problems/valid-tic-tac-toe-state/)| * Valid Tic-Tac-Toe State|29%|Medium||
|[0795](https://leetcode.com/problems/number-of-subarrays-with-bounded-maximum/)| * Number of Subarrays with Bounded Maximum|42%|Medium||
|[0796](https://leetcode.com/problems/rotate-string/)| * Rotate String|48%|Easy||
|[0797](https://leetcode.com/problems/all-paths-from-source-to-target/)| * All Paths From Source to Target|69%|Medium||
|[0798](https://leetcode.com/problems/smallest-rotation-with-highest-score/)| * Smallest Rotation with Highest Score|39%|Hard||
|[0799](https://leetcode.com/problems/champagne-tower/)| * Champagne Tower|33%|Medium||
|[0801](https://leetcode.com/problems/minimum-swaps-to-make-sequences-increasing/)| * Minimum Swaps To Make Sequences Increasing|34%|Medium||
|[0802](https://leetcode.com/problems/find-eventual-safe-states/)| * Find Eventual Safe States|43%|Medium||
|[0803](https://leetcode.com/problems/bricks-falling-when-hit/)| * Bricks Falling When Hit|28%|Hard||
|[0804](https://leetcode.com/problems/unique-morse-code-words/)| * Unique Morse Code Words|74%|Easy||
|[0805](https://leetcode.com/problems/split-array-with-same-average/)| * Split Array With Same Average|23%|Hard||
|[0806](https://leetcode.com/problems/number-of-lines-to-write-string/)| * Number of Lines To Write String|62%|Easy||
|[0807](https://leetcode.com/problems/max-increase-to-keep-city-skyline/)| * Max Increase to Keep City Skyline|81%|Medium||
|[0808](https://leetcode.com/problems/soup-servings/)| * Soup Servings|36%|Medium||
|[0809](https://leetcode.com/problems/expressive-words/)| * Expressive Words|42%|Medium||
|[0810](https://leetcode.com/problems/chalkboard-xor-game/)| * Chalkboard XOR Game|44%|Hard||
|[0811](https://leetcode.com/problems/subdomain-visit-count/)| * Subdomain Visit Count|64%|Easy||
|[0812](https://leetcode.com/problems/largest-triangle-area/)| * Largest Triangle Area|55%|Easy||
|[0813](https://leetcode.com/problems/largest-sum-of-averages/)| * Largest Sum of Averages|44%|Medium||
|[0814](https://leetcode.com/problems/binary-tree-pruning/)| * Binary Tree Pruning|70%|Medium||
|[0815](https://leetcode.com/problems/bus-routes/)| * Bus Routes|39%|Hard||
|[0816](https://leetcode.com/problems/ambiguous-coordinates/)| * Ambiguous Coordinates|43%|Medium||
|[0817](https://leetcode.com/problems/linked-list-components/)| * Linked List Components|54%|Medium||
|[0818](https://leetcode.com/problems/race-car/)| * Race Car|34%|Hard||
|[0819](https://leetcode.com/problems/most-common-word/)| * Most Common Word|42%|Easy||
|[0820](https://leetcode.com/problems/short-encoding-of-words/)| * Short Encoding of Words|46%|Medium||
|[0821](https://leetcode.com/problems/shortest-distance-to-a-character/)| * Shortest Distance to a Character|62%|Easy||
|[0822](https://leetcode.com/problems/card-flipping-game/)| * Card Flipping Game|40%|Medium||
|[0823](https://leetcode.com/problems/binary-trees-with-factors/)| * Binary Trees With Factors|32%|Medium||
|[0824](https://leetcode.com/problems/goat-latin/)| * Goat Latin|57%|Easy||
|[0825](https://leetcode.com/problems/friends-of-appropriate-ages/)| * Friends Of Appropriate Ages|35%|Medium||
|[0826](https://leetcode.com/problems/most-profit-assigning-work/)| * Most Profit Assigning Work|35%|Medium||
|[0827](https://leetcode.com/problems/making-a-large-island/)| * Making A Large Island|42%|Hard||
|[0828](https://leetcode.com/problems/unique-letter-string/)| * Unique Letter String|38%|Hard||
|[0829](https://leetcode.com/problems/consecutive-numbers-sum/)| * Consecutive Numbers Sum|32%|Hard||
|[0830](https://leetcode.com/problems/positions-of-large-groups/)| * Positions of Large Groups|47%|Easy||
|[0831](https://leetcode.com/problems/masking-personal-information/)| * Masking Personal Information|41%|Medium||
|[0832](https://leetcode.com/problems/flipping-an-image/)| * Flipping an Image|71%|Easy||
|[0833](https://leetcode.com/problems/find-and-replace-in-string/)| * Find And Replace in String|45%|Medium||
|[0834](https://leetcode.com/problems/sum-of-distances-in-tree/)| * Sum of Distances in Tree|39%|Hard||
|[0835](https://leetcode.com/problems/image-overlap/)| * Image Overlap|51%|Medium||
|[0836](https://leetcode.com/problems/rectangle-overlap/)| * Rectangle Overlap|46%|Easy||
|[0837](https://leetcode.com/problems/new-21-game/)| * New 21 Game|31%|Medium||
|[0838](https://leetcode.com/problems/push-dominoes/)| * Push Dominoes|43%|Medium||
|[0839](https://leetcode.com/problems/similar-string-groups/)| * Similar String Groups|34%|Hard||
|[0840](https://leetcode.com/problems/magic-squares-in-grid/)| * Magic Squares In Grid|35%|Easy||
|[0841](https://leetcode.com/problems/keys-and-rooms/)| * Keys and Rooms|59%|Medium||
|[0842](https://leetcode.com/problems/split-array-into-fibonacci-sequence/)| * Split Array into Fibonacci Sequence|34%|Medium||
|[0843](https://leetcode.com/problems/guess-the-word/)| * Guess the Word|42%|Hard||
|[0844](https://leetcode.com/problems/backspace-string-compare/)| * Backspace String Compare|45%|Easy||
|[0845](https://leetcode.com/problems/longest-mountain-in-array/)| * Longest Mountain in Array|34%|Medium||
|[0846](https://leetcode.com/problems/hand-of-straights/)| * Hand of Straights|48%|Medium||
|[0847](https://leetcode.com/problems/shortest-path-visiting-all-nodes/)| * Shortest Path Visiting All Nodes|46%|Hard||
|[0848](https://leetcode.com/problems/shifting-letters/)| * Shifting Letters|40%|Medium||
|[0849](https://leetcode.com/problems/maximize-distance-to-closest-person/)| * Maximize Distance to Closest Person|40%|Easy||
|[0850](https://leetcode.com/problems/rectangle-area-ii/)| * Rectangle Area II|44%|Hard||
|[0851](https://leetcode.com/problems/loud-and-rich/)| * Loud and Rich|47%|Medium||
|[0852](https://leetcode.com/problems/peak-index-in-a-mountain-array/)| * Peak Index in a Mountain Array|69%|Easy||
|[0853](https://leetcode.com/problems/car-fleet/)| * Car Fleet|38%|Medium||
|[0854](https://leetcode.com/problems/k-similar-strings/)| * K-Similar Strings|33%|Hard||
|[0855](https://leetcode.com/problems/exam-room/)| * Exam Room|37%|Medium||
|[0856](https://leetcode.com/problems/score-of-parentheses/)| * Score of Parentheses|55%|Medium||
|[0857](https://leetcode.com/problems/minimum-cost-to-hire-k-workers/)| * Minimum Cost to Hire K Workers|47%|Hard||
|[0858](https://leetcode.com/problems/mirror-reflection/)| * Mirror Reflection|51%|Medium||
|[0859](https://leetcode.com/problems/buddy-strings/)| * Buddy Strings|27%|Easy||
|[0860](https://leetcode.com/problems/lemonade-change/)| * Lemonade Change|50%|Easy||
|[0861](https://leetcode.com/problems/score-after-flipping-matrix/)| * Score After Flipping Matrix|68%|Medium||
|[0862](https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/)| * Shortest Subarray with Sum at Least K|21%|Hard||
|[0863](https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/)| * All Nodes Distance K in Binary Tree|46%|Medium||
|[0864](https://leetcode.com/problems/shortest-path-to-get-all-keys/)| * Shortest Path to Get All Keys|35%|Hard||
|[0865](https://leetcode.com/problems/smallest-subtree-with-all-the-deepest-nodes/)| * Smallest Subtree with all the Deepest Nodes|54%|Medium||
|[0866](https://leetcode.com/problems/prime-palindrome/)| * Prime Palindrome|19%|Medium||
|[0867](https://leetcode.com/problems/transpose-matrix/)| * Transpose Matrix|63%|Easy||
|[0868](https://leetcode.com/problems/binary-gap/)| * Binary Gap|59%|Easy||
|[0869](https://leetcode.com/problems/reordered-power-of-2/)| * Reordered Power of 2|50%|Medium||
|[0870](https://leetcode.com/problems/advantage-shuffle/)| * Advantage Shuffle|41%|Medium||
|[0871](https://leetcode.com/problems/minimum-number-of-refueling-stops/)| * Minimum Number of Refueling Stops|28%|Hard||
|[0872](https://leetcode.com/problems/leaf-similar-trees/)| * Leaf-Similar Trees|62%|Easy||
|[0873](https://leetcode.com/problems/length-of-longest-fibonacci-subsequence/)| * Length of Longest Fibonacci Subsequence|45%|Medium||
|[0874](https://leetcode.com/problems/walking-robot-simulation/)| * Walking Robot Simulation|31%|Easy||
|[0875](https://leetcode.com/problems/koko-eating-bananas/)| * Koko Eating Bananas|45%|Medium||
|[0876](https://leetcode.com/problems/middle-of-the-linked-list/)| * Middle of the Linked List|63%|Easy||
|[0877](https://leetcode.com/problems/stone-game/)| * Stone Game|60%|Medium||
|[0878](https://leetcode.com/problems/nth-magical-number/)| * Nth Magical Number|24%|Hard||
|[0879](https://leetcode.com/problems/profitable-schemes/)| * Profitable Schemes|35%|Hard||
|[0880](https://leetcode.com/problems/decoded-string-at-index/)| * Decoded String at Index|22%|Medium||
|[0881](https://leetcode.com/problems/boats-to-save-people/)| * Boats to Save People|43%|Medium||
|[0882](https://leetcode.com/problems/reachable-nodes-in-subdivided-graph/)| * Reachable Nodes In Subdivided Graph|37%|Hard||
|[0883](https://leetcode.com/problems/projection-area-of-3d-shapes/)| * Projection Area of 3D Shapes|65%|Easy||
|[0884](https://leetcode.com/problems/uncommon-words-from-two-sentences/)| * Uncommon Words from Two Sentences|60%|Easy||
|[0885](https://leetcode.com/problems/spiral-matrix-iii/)| * Spiral Matrix III|63%|Medium||
|[0886](https://leetcode.com/problems/possible-bipartition/)| * Possible Bipartition|40%|Medium||
|[0887](https://leetcode.com/problems/super-egg-drop/)| * Super Egg Drop|24%|Hard||
|[0888](https://leetcode.com/problems/fair-candy-swap/)| * Fair Candy Swap|56%|Easy||
|[0889](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/)| * Construct Binary Tree from Preorder and Postorder Traversal|59%|Medium||
|[0890](https://leetcode.com/problems/find-and-replace-pattern/)| * Find and Replace Pattern|70%|Medium||
|[0891](https://leetcode.com/problems/sum-of-subsequence-widths/)| * Sum of Subsequence Widths|28%|Hard||
|[0892](https://leetcode.com/problems/surface-area-of-3d-shapes/)| * Surface Area of 3D Shapes|55%|Easy||
|[0893](https://leetcode.com/problems/groups-of-special-equivalent-strings/)| * Groups of Special-Equivalent Strings|61%|Easy||
|[0894](https://leetcode.com/problems/all-possible-full-binary-trees/)| * All Possible Full Binary Trees|70%|Medium||
|[0895](https://leetcode.com/problems/maximum-frequency-stack/)| * Maximum Frequency Stack|53%|Hard||
|[0896](https://leetcode.com/problems/monotonic-array/)| * Monotonic Array|54%|Easy||
|[0897](https://leetcode.com/problems/increasing-order-search-tree/)| * Increasing Order Search Tree|63%|Easy||
|[0898](https://leetcode.com/problems/bitwise-ors-of-subarrays/)| * Bitwise ORs of Subarrays|33%|Medium||
|[0899](https://leetcode.com/problems/orderly-queue/)| * Orderly Queue|46%|Hard||
|[0900](https://leetcode.com/problems/rle-iterator/)| * RLE Iterator|49%|Medium||
|[0901](https://leetcode.com/problems/online-stock-span/)| * Online Stock Span|47%|Medium||
|[0902](https://leetcode.com/problems/numbers-at-most-n-given-digit-set/)| * Numbers At Most N Given Digit Set|28%|Hard||
|[0903](https://leetcode.com/problems/valid-permutations-for-di-sequence/)| * Valid Permutations for DI Sequence|42%|Hard||
|[0904](https://leetcode.com/problems/fruit-into-baskets/)| * Fruit Into Baskets|41%|Medium||
|[0905](https://leetcode.com/problems/sort-array-by-parity/)| * Sort Array By Parity|72%|Easy||
|[0906](https://leetcode.com/problems/super-palindromes/)| * Super Palindromes|30%|Hard||
|[0907](https://leetcode.com/problems/sum-of-subarray-minimums/)| * Sum of Subarray Minimums|26%|Medium||
|[0908](https://leetcode.com/problems/smallest-range-i/)| * Smallest Range I|64%|Easy||
|[0909](https://leetcode.com/problems/snakes-and-ladders/)| * Snakes and Ladders|31%|Medium||
|[0910](https://leetcode.com/problems/smallest-range-ii/)| * Smallest Range II|23%|Medium||
|[0911](https://leetcode.com/problems/online-election/)| * Online Election|45%|Medium||
|[0913](https://leetcode.com/problems/cat-and-mouse/)| * Cat and Mouse|27%|Hard||
|[0914](https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/)| * X of a Kind in a Deck of Cards|34%|Easy||
|[0915](https://leetcode.com/problems/partition-array-into-disjoint-intervals/)| * Partition Array into Disjoint Intervals|42%|Medium||
|[0916](https://leetcode.com/problems/word-subsets/)| * Word Subsets|44%|Medium||
|[0917](https://leetcode.com/problems/reverse-only-letters/)| * Reverse Only Letters|55%|Easy||
|[0918](https://leetcode.com/problems/maximum-sum-circular-subarray/)| * Maximum Sum Circular Subarray|31%|Medium||
|[0919](https://leetcode.com/problems/complete-binary-tree-inserter/)| * Complete Binary Tree Inserter|55%|Medium||
|[0920](https://leetcode.com/problems/number-of-music-playlists/)| * Number of Music Playlists|43%|Hard||
|[0921](https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/)| * Minimum Add to Make Parentheses Valid|69%|Medium||
|[0922](https://leetcode.com/problems/sort-array-by-parity-ii/)| * Sort Array By Parity II|66%|Easy||
|[0923](https://leetcode.com/problems/3sum-with-multiplicity/)| * 3Sum With Multiplicity|32%|Medium||
|[0924](https://leetcode.com/problems/minimize-malware-spread/)| * Minimize Malware Spread|39%|Hard||
|[0925](https://leetcode.com/problems/long-pressed-name/)| * Long Pressed Name|44%|Easy||
|[0926](https://leetcode.com/problems/flip-string-to-monotone-increasing/)| * Flip String to Monotone Increasing|48%|Medium||
|[0927](https://leetcode.com/problems/three-equal-parts/)| * Three Equal Parts|30%|Hard||
|[0928](https://leetcode.com/problems/minimize-malware-spread-ii/)| * Minimize Malware Spread II|38%|Hard||
|[0929](https://leetcode.com/problems/unique-email-addresses/)| * Unique Email Addresses|73%|Easy||
|[0930](https://leetcode.com/problems/binary-subarrays-with-sum/)| * Binary Subarrays With Sum|37%|Medium||
|[0931](https://leetcode.com/problems/minimum-falling-path-sum/)| * Minimum Falling Path Sum|58%|Medium||
|[0932](https://leetcode.com/problems/beautiful-array/)| * Beautiful Array|52%|Medium||
|[0933](https://leetcode.com/problems/number-of-recent-calls/)| * Number of Recent Calls|69%|Easy||
|[0934](https://leetcode.com/problems/shortest-bridge/)| * Shortest Bridge|43%|Medium||
|[0935](https://leetcode.com/problems/knight-dialer/)| * Knight Dialer|39%|Medium||
|[0936](https://leetcode.com/problems/stamping-the-sequence/)| * Stamping The Sequence|35%|Hard||
|[0937](https://leetcode.com/problems/reorder-log-files/)| * Reorder Log Files|59%|Easy||
|[0938](https://leetcode.com/problems/range-sum-of-bst/)| * Range Sum of BST|80%|Medium||
|[0939](https://leetcode.com/problems/minimum-area-rectangle/)| * Minimum Area Rectangle|49%|Medium||
|[0940](https://leetcode.com/problems/distinct-subsequences-ii/)| * Distinct Subsequences II|38%|Hard||
|[0941](https://leetcode.com/problems/valid-mountain-array/)| * Valid Mountain Array|34%|Easy||
|[0942](https://leetcode.com/problems/di-string-match/)| * DI String Match|70%|Easy||
|[0943](https://leetcode.com/problems/find-the-shortest-superstring/)| * Find the Shortest Superstring|36%|Hard||
|[0944](https://leetcode.com/problems/delete-columns-to-make-sorted/)| * Delete Columns to Make Sorted|69%|Easy||
|[0945](https://leetcode.com/problems/minimum-increment-to-make-array-unique/)| * Minimum Increment to Make Array Unique|42%|Medium||
|[0946](https://leetcode.com/problems/validate-stack-sequences/)| * Validate Stack Sequences|57%|Medium||
|[0947](https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/)| * Most Stones Removed with Same Row or Column|53%|Medium||
|[0948](https://leetcode.com/problems/bag-of-tokens/)| * Bag of Tokens|39%|Medium||
|[0949](https://leetcode.com/problems/largest-time-for-given-digits/)| * Largest Time for Given Digits|33%|Easy||
|[0950](https://leetcode.com/problems/reveal-cards-in-increasing-order/)| * Reveal Cards In Increasing Order|71%|Medium||
|[0951](https://leetcode.com/problems/flip-equivalent-binary-trees/)| * Flip Equivalent Binary Trees|65%|Medium||
|[0952](https://leetcode.com/problems/largest-component-size-by-common-factor/)| * Largest Component Size by Common Factor|25%|Hard||
|[0953](https://leetcode.com/problems/verifying-an-alien-dictionary/)| * Verifying an Alien Dictionary|55%|Easy||
|[0954](https://leetcode.com/problems/array-of-doubled-pairs/)| * Array of Doubled Pairs|34%|Medium||
|[0955](https://leetcode.com/problems/delete-columns-to-make-sorted-ii/)| * Delete Columns to Make Sorted II|31%|Medium||
|[0956](https://leetcode.com/problems/tallest-billboard/)| * Tallest Billboard|37%|Hard||
|[0957](https://leetcode.com/problems/prison-cells-after-n-days/)| * Prison Cells After N Days|37%|Medium||
|[0958](https://leetcode.com/problems/check-completeness-of-a-binary-tree/)| * Check Completeness of a Binary Tree|46%|Medium||
|[0959](https://leetcode.com/problems/regions-cut-by-slashes/)| * Regions Cut By Slashes|61%|Medium||
|[0960](https://leetcode.com/problems/delete-columns-to-make-sorted-iii/)| * Delete Columns to Make Sorted III|51%|Hard||
|[0961](https://leetcode.com/problems/n-repeated-element-in-size-2n-array/)| * N-Repeated Element in Size 2N Array|72%|Easy||
|[0962](https://leetcode.com/problems/maximum-width-ramp/)| * Maximum Width Ramp|40%|Medium||
|[0963](https://leetcode.com/problems/minimum-area-rectangle-ii/)| * Minimum Area Rectangle II|43%|Medium||
|[0964](https://leetcode.com/problems/least-operators-to-express-number/)| * Least Operators to Express Number|40%|Hard||
|[0965](https://leetcode.com/problems/univalued-binary-tree/)| * Univalued Binary Tree|67%|Easy||
|[0966](https://leetcode.com/problems/vowel-spellchecker/)| * Vowel Spellchecker|40%|Medium||
|[0967](https://leetcode.com/problems/numbers-with-same-consecutive-differences/)| * Numbers With Same Consecutive Differences|36%|Medium||
|[0968](https://leetcode.com/problems/binary-tree-cameras/)| * Binary Tree Cameras|34%|Hard||
|[0969](https://leetcode.com/problems/pancake-sorting/)| * Pancake Sorting|61%|Medium||
|[0970](https://leetcode.com/problems/powerful-integers/)| * Powerful Integers|39%|Easy||
|[0971](https://leetcode.com/problems/flip-binary-tree-to-match-preorder-traversal/)| * Flip Binary Tree To Match Preorder Traversal|42%|Medium||
|[0972](https://leetcode.com/problems/equal-rational-numbers/)| * Equal Rational Numbers|40%|Hard||
|[0973](https://leetcode.com/problems/k-closest-points-to-origin/)| * K Closest Points to Origin|64%|Medium||
|[0974](https://leetcode.com/problems/subarray-sums-divisible-by-k/)| * Subarray Sums Divisible by K|43%|Medium||
|[0975](https://leetcode.com/problems/odd-even-jump/)| * Odd Even Jump|49%|Hard||
|[0976](https://leetcode.com/problems/largest-perimeter-triangle/)| * Largest Perimeter Triangle|57%|Easy||
|[0977](https://leetcode.com/problems/squares-of-a-sorted-array/)| * Squares of a Sorted Array|72%|Easy||
|[0978](https://leetcode.com/problems/longest-turbulent-subarray/)| * Longest Turbulent Subarray|45%|Medium||
|[0979](https://leetcode.com/problems/distribute-coins-in-binary-tree/)| * Distribute Coins in Binary Tree|66%|Medium||
|[0980](https://leetcode.com/problems/unique-paths-iii/)| * Unique Paths III|71%|Hard||
|[0981](https://leetcode.com/problems/time-based-key-value-store/)| * Time Based Key-Value Store|50%|Medium||
|[0982](https://leetcode.com/problems/triples-with-bitwise-and-equal-to-zero/)| * Triples with Bitwise AND Equal To Zero|53%|Hard||
|[0983](https://leetcode.com/problems/minimum-cost-for-tickets/)| * Minimum Cost For Tickets|57%|Medium||
|[0984](https://leetcode.com/problems/string-without-aaa-or-bbb/)| * String Without AAA or BBB|32%|Medium||
|[0985](https://leetcode.com/problems/sum-of-even-numbers-after-queries/)| * Sum of Even Numbers After Queries|65%|Easy||
|[0986](https://leetcode.com/problems/interval-list-intersections/)| * Interval List Intersections|62%|Medium||
|[0987](https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/)| * Vertical Order Traversal of a Binary Tree|32%|Medium||
|[0988](https://leetcode.com/problems/smallest-string-starting-from-leaf/)| * Smallest String Starting From Leaf|48%|Medium||
|[0989](https://leetcode.com/problems/add-to-array-form-of-integer/)| * Add to Array-Form of Integer|44%|Easy||
|[0990](https://leetcode.com/problems/satisfiability-of-equality-equations/)| * Satisfiability of Equality Equations|39%|Medium||
|[0991](https://leetcode.com/problems/broken-calculator/)| * Broken Calculator|39%|Medium||
|[0992](https://leetcode.com/problems/subarrays-with-k-different-integers/)| * Subarrays with K Different Integers|44%|Hard||
|[0993](https://leetcode.com/problems/cousins-in-binary-tree/)| * Cousins in Binary Tree|52%|Easy||
|[0994](https://leetcode.com/problems/rotting-oranges/)| * Rotting Oranges|46%|Easy||
|[0995](https://leetcode.com/problems/minimum-number-of-k-consecutive-bit-flips/)| * Minimum Number of K Consecutive Bit Flips|49%|Hard||
|[0996](https://leetcode.com/problems/number-of-squareful-arrays/)| * Number of Squareful Arrays|47%|Hard||
|[0997](https://leetcode.com/problems/find-the-town-judge/)| * Find the Town Judge|48%|Easy||
|[0998](https://leetcode.com/problems/maximum-binary-tree-ii/)| * Maximum Binary Tree II|62%|Medium||
|[0999](https://leetcode.com/problems/available-captures-for-rook/)| * Available Captures for Rook|67%|Easy||
|[1000](https://leetcode.com/problems/minimum-cost-to-merge-stones/)| * Minimum Cost to Merge Stones|27%|Hard||
|[1001](https://leetcode.com/problems/grid-illumination/)| * Grid Illumination|34%|Hard||
|[1002](https://leetcode.com/problems/find-common-characters/)| * Find Common Characters|66%|Easy||
|[1003](https://leetcode.com/problems/check-if-word-is-valid-after-substitutions/)| * Check If Word Is Valid After Substitutions|51%|Medium||
|[1004](https://leetcode.com/problems/max-consecutive-ones-iii/)| * Max Consecutive Ones III|52%|Medium||
|[1005](https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/)| * Maximize Sum Of Array After K Negations|49%|Easy||
|[1006](https://leetcode.com/problems/clumsy-factorial/)| * Clumsy Factorial|54%|Medium||
|[1007](https://leetcode.com/problems/minimum-domino-rotations-for-equal-row/)| * Minimum Domino Rotations For Equal Row|45%|Medium||
|[1008](https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/)| * Construct Binary Search Tree from Preorder Traversal|72%|Medium||
|[1009](https://leetcode.com/problems/complement-of-base-10-integer/)| * Complement of Base 10 Integer|58%|Easy||
|[1010](https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/)| * Pairs of Songs With Total Durations Divisible by 60|44%|Easy||
|[1011](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/)| * Capacity To Ship Packages Within D Days|50%|Medium||
|[1012](https://leetcode.com/problems/numbers-with-repeated-digits/)| * Numbers With Repeated Digits|33%|Hard||
|[1013](https://leetcode.com/problems/partition-array-into-three-parts-with-equal-sum/)| * Partition Array Into Three Parts With Equal Sum|53%|Easy||
|[1014](https://leetcode.com/problems/best-sightseeing-pair/)| * Best Sightseeing Pair|46%|Medium||
|[1015](https://leetcode.com/problems/smallest-integer-divisible-by-k/)| * Smallest Integer Divisible by K|26%|Medium||
|[1016](https://leetcode.com/problems/binary-string-with-substrings-representing-1-to-n/)| * Binary String With Substrings Representing 1 To N|65%|Medium||
|[1017](https://leetcode.com/problems/convert-to-base-2/)| * Convert to Base -2|54%|Medium||
|[1018](https://leetcode.com/problems/binary-prefix-divisible-by-5/)| * Binary Prefix Divisible By 5|44%|Easy||
|[1019](https://leetcode.com/problems/next-greater-node-in-linked-list/)| * Next Greater Node In Linked List|56%|Medium||
|[1020](https://leetcode.com/problems/number-of-enclaves/)| * Number of Enclaves|54%|Medium||
|[1021](https://leetcode.com/problems/remove-outermost-parentheses/)| * Remove Outermost Parentheses :new: |84%|Easy||
|[1022](https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/)| * Sum of Root To Leaf Binary Numbers :new: |39%|Easy||
|[1023](https://leetcode.com/problems/camelcase-matching/)| * Camelcase Matching :new: |58%|Medium||
|[1024](https://leetcode.com/problems/video-stitching/)| * Video Stitching :new: |47%|Medium||


以下免费的算法题，暂时不能提交 Go 解答

- [116.Populating Next Right Pointers in Each Node](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/)
- [117.Populating Next Right Pointers in Each Node II](https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/)
- [133.Clone Graph](https://leetcode.com/problems/clone-graph/)
- [138.Copy List with Random Pointer](https://leetcode.com/problems/copy-list-with-random-pointer/)
- [151.Reverse Words in a String](https://leetcode.com/problems/reverse-words-in-a-string/)
- [278.First Bad Version](https://leetcode.com/problems/first-bad-version/)
- [284.Peeking Iterator](https://leetcode.com/problems/peeking-iterator/)
- [297.Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/)
- [341.Flatten Nested List Iterator](https://leetcode.com/problems/flatten-nested-list-iterator/)
- [374.Guess Number Higher or Lower](https://leetcode.com/problems/guess-number-higher-or-lower/)
- [426.Convert Binary Search Tree to Sorted Doubly Linked List](https://leetcode.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/)
- [427.Construct Quad Tree](https://leetcode.com/problems/construct-quad-tree/)
- [429.N-ary Tree Level Order Traversal](https://leetcode.com/problems/n-ary-tree-level-order-traversal/)
- [430.Flatten a Multilevel Doubly Linked List](https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/)
- [449.Serialize and Deserialize BST](https://leetcode.com/problems/serialize-and-deserialize-bst/)
- [535.Encode and Decode TinyURL](https://leetcode.com/problems/encode-and-decode-tinyurl/)
- [558.Quad Tree Intersection](https://leetcode.com/problems/quad-tree-intersection/)
- [559.Maximum Depth of N-ary Tree](https://leetcode.com/problems/maximum-depth-of-n-ary-tree/)
- [589.N-ary Tree Preorder Traversal](https://leetcode.com/problems/n-ary-tree-preorder-traversal/)
- [590.N-ary Tree Postorder Traversal](https://leetcode.com/problems/n-ary-tree-postorder-traversal/)
- [690.Employee Importance](https://leetcode.com/problems/employee-importance/)
- [708.Insert into a Cyclic Sorted List](https://leetcode.com/problems/insert-into-a-cyclic-sorted-list/)

## 题解

[notes](practice/notes) LeetCode 题解，也会记录我答题过程中对知识点的总结。

----

## 学习方法

1. 仔细研读经典算法书籍（详细书单见后文）
2. 把所有经典数据结构和算法都写一遍
3. 刷题（LeetCode等）
4. 找志同道合的人一起交流学习
5. 参加比赛
6. 。。。

----

## 推荐阅读

### 基础（<b style="color:red">进行中...</b>）

- ☆☆☆☆☆《算法基础-打开算法之门》科尔曼[算法导论作者之一]（<b style="color:blue">阅读进度20%</b>）
- ☆☆☆☆☆《算法导论》 是一本对算法介绍比较全面的经典书籍
- ☆☆☆☆《编程珠玑I》
- ☆☆☆☆《编程珠玑II》
- ☆☆☆《算法-Algorithms》第四版，普林斯顿大学，这本近千页的书只有6章,其中四章分别是排序，查找，图，字符串，足见介绍的有多么的细致。（**暂时不阅读**）

### 算法演示

- [David Galles 的可视化的数据结构和算法](http://www.cs.usfca.edu/~galles/visualization/)
- [酷壳-可视化的数据结构和算法](https://coolshell.cn/articles/4671.html)

### 编程网站

1. [LeetCode](http://leetcode.com/) 全球最大的在线程序评测系统。
2. [OpenJudge](http://openjudge.cn/) OpenJudge是开放的在线程序评测系统 您可以创建自己的OJ小组。

### 笔试刷题必备

- 《剑指offer》
- 《编程之法:面试和算法心得》
- 《算法谜题》
- 《编程之美》

### 延伸阅读

- 《推荐系统-技术、评估及高效算法》
- 《深入理解计算机系统》
- 《TCP/IP详解一二三卷》

## 其他

### 1. [awesome-algorithms](https://github.com/tayllan/awesome-algorithms)

### 2. [July 博客](http://blog.csdn.net/v_july_v)

>比方说：《数据挖掘十大算法系列》、《经典算法研究》、《BAT机器学习面试1000题系列》、《程序员编程艺术》等系列。

- [The-Art-Of-Programming-By-July](https://github.com/julycoding/The-Art-Of-Programming-By-July)
- [程序员编程艺术](http://blog.csdn.net/v_JULY_v/article/details/6460494)

## 如何参与/贡献？

1. Fork 此项目
2. 克隆你自己的项目到你本地 （git clone https://github.com/your_github_name/learning-algorithms.git）
2. 创建你新的 feature 分支 (git checkout -b my_feature)
3. 添加并提交你的修改内容 (git commit -am 'Add some feature')
4. 推送到你自己项目的远端 feature 分支 (git push origin my_feature)
5. 创建一个新的 PR（Pull Request）

## 参考资料

- [阻碍你使用 GraphQL 的十个问题](https://leetcode-cn.com/articles/%E9%98%BB%E7%A2%8D%E4%BD%A0%E4%BD%BF%E7%94%A8-graphql-%E7%9A%84%E5%8D%81%E4%B8%AA%E9%97%AE%E9%A2%98/)

## License

License: BSD 3-Clause License
