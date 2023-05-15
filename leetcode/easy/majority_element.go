package easy

/*
   169. Majority Element

   Given an array nums of size n, return the majority element.

   The majority element is the element that appears more than ⌊n / 2⌋ times.
   You may assume that the majority element always exists in the array.



   Example 1:

   Input: nums = [3,2,3]
   Output: 3

   Example 2:

   Input: nums = [2,2,1,1,1,2,2]
   Output: 2
*/

func MajorityElement(nums []int) int {
	cache := make(map[int]int)
	max := 0
	var majority int

	for _, element := range nums {
		if val, ok := cache[element]; ok {
			cache[element] = val + 1
		} else {
			cache[element] = 1
		}
	}

	for key, val := range cache {
		if val > max {
			max = val
			majority = key
		}
	}

	return majority
}
