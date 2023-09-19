# [761. Special Binary String](https://leetcode.com/problems/special-binary-string/)

## 题目

Special binary strings are binary strings with the following two properties:

1. The number of 0's is equal to the number of 1's.
1. Every prefix of the binary string has at least as many 1's as 0's.

Given a special string S, a move consists of choosing two consecutive, non-empty, special substrings of S, and swapping them.  (Two strings are consecutive if the last character of the first string is exactly one index before the first character of the second string.)

At the end of any number of moves, what is the lexicographically largest resulting string possible?

Example 1:

```text
Input: S = "11011000"
Output: "11100100"
Explanation:
The strings "10" [occuring at S[1]] and "1100" [at S[3]] are swapped.
This is the lexicographically largest string possible after some number of swaps.
```

Note:

1. S has length at most 50.
1. S is guaranteed to be a special binary string as defined above.

## 解题思路

见程序注释
