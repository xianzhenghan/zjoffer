package main

import (
	"fmt"
	"math"
)

/*
*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组
是数组中的一个连续部分。

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23
*/
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1], 0) + nums[i]
	}
	fmt.Println(dp)
	return maxSlice(dp)
}
func maxSlice(nums []int) (max int) {
	max = math.MinInt
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(maxSubArray([]int{4, -1, 2, 1, -9}))
}
