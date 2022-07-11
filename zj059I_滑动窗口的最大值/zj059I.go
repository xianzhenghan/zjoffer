package zj059I_滑动窗口的最大值

import (
	"fmt"
)

//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7]
func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	length := len(nums)
	dp := make([]int, (length - k + 1))

	for i := 0; i < (length - k + 1); i++ {
		dp[i] = maxNums(nums[i:i+k], k)
	}
	return dp
}

//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7]
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	length := len(nums)
	dp := make([]int, (length - k + 1))
	dp[0] = maxNums(nums[0:k], k)
	for i := 1; i < (length - k + 1); i++ {
		//nums[i-1]
		//nums[i+k-1]
		if nums[i+k-1] >= (dp[i-1]) {
			dp[i] = nums[i+k-1]
		} else {
			if nums[i-1] < dp[i-1] {
				dp[i] = dp[i-1]
			} else {
				dp[i] = maxNums(nums[i:i+k], k)
			}
		}
	}
	return dp
}

func maxNums(nums []int, k int) int {
	fmt.Printf("nums=%v  ", nums)
	max := ^int(^uint32((0)) >> 1)
	for idx, num := range nums {
		if num > max && idx < k {
			max = num
		}
	}
	fmt.Printf("max=%d  \n", max)
	return max
}
