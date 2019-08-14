/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package leetcode

func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}