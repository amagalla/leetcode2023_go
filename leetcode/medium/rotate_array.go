package medium

/*
   189. Rotate Array

   Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

   modify nums in place

   Example 1:

   Input: nums = [1,2,3,4,5,6,7], k = 3
   Output: [5,6,7,1,2,3,4]
   Explanation:
   rotate 1 steps to the right: [7,1,2,3,4,5,6]
   rotate 2 steps to the right: [6,7,1,2,3,4,5]
   rotate 3 steps to the right: [5,6,7,1,2,3,4]
   Example 2:

   Input: nums = [-1,-100,3,99], k = 2
   Output: [3,99,-1,-100]
   Explanation:
   rotate 1 steps to the right: [99,-1,-100,3]
   rotate 2 steps to the right: [3,99,-1,-100]
*/

func Rotate(nums []int, k int) []int {
	k = k % len(nums)

	reverse(nums, 0, len(nums)-1)

	reverse(nums, 0, k-1)

	reverse(nums, k, len(nums)-1)

	return nums

}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
