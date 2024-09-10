package zj042_连续子数组的最大值

import (
	"math"
	"slices"
)

// 输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出: 6
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1], 0) + nums[i]
	}
	return slices.Max(dp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArrayV2(nums []int) int {
	ans := math.MinInt32
	f := 0
	for i := 0; i < len(nums); i++ {
		f = max(f, 0) + nums[i]
		ans = max(ans, f)
	}
	return ans
}
