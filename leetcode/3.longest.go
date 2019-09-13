package main

import (
	"fmt"
)

//Given a string, find the length of the longest substring without repeating characters.
//
//
// Example 1:
//
//
//Input: "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//
//
//
// Example 2:
//
//
//Input: "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//
//
//
// Example 3:
//
//
//Input: "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
//
//j
//
//
//

func lengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		start := s[i]
		subStr := []uint8{start}

		for j := i + 1; j < len(s); j++ {
			contains := false
			for _, subChar := range subStr {
				if s[j] == subChar {
					contains = true
					break
				}
			}
			if contains {
				break
			}
			subStr = append(subStr, s[j])
		}
		if max < len(subStr) {
			max = len(subStr)
		}
	}

	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
