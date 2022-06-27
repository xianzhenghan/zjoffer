package zj042_连续子数组的最大值

//输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	dp := make([]int, length)
	dp[0] = nums[0]
	result := nums[0]
	for i := 1; i < length; i++ {
		dp[i] = max((dp[i-1] + nums[i]), nums[i])
		if dp[i] > result {
			result = dp[i]
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
