package easy

import "strings"

/*
   14. Longest Common Prefix

   Write a function to find the longest common prefix string amongst an array of strings.

   If there is no common prefix, return an empty string "".

   Example 1:

   Input: strs = ["flower","flow","flight"]
   Output: "fl"
   Example 2:

   Input: strs = ["dog","racecar","car"]
   Output: ""
   Explanation: There is no common prefix among the input strings.
*/

func LongestCommonPrefix(strs []string) string {
	i := 1
	pre := strs[0]

	for i < len(strs) {
		if !strings.HasPrefix(strs[i], pre) {
			pre = pre[:len(pre)-1]
		} else {
			i++
		}
	}

	return pre
}
