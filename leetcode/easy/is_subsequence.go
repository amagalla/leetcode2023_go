package easy

/*
   isSubsequence


   Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

   A subsequence of a string is a new string that is formed from the original string by deleting
   some (can be none) of the characters without disturbing the relative positions of the remaining
   characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

   Example 1:

   Input: s = "abc", t = "ahbgdc"
   Output: true
   Example 2:

   Input: s = "axc", t = "ahbgdc"
   Output: false
*/

func IsSubsequence(s string, t string) bool {
	slow := 0

	for i := 0; i < len(t); i++ {
		if t[i] == s[slow] {
			slow++
		}
	}

	return slow == len(s)
}
