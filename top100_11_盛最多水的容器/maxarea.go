package top100_11_盛最多水的容器

import "math"

//n == height.length
//2 <= n <= 105
//0 <= height[i] <= 104
func maxArea(height []int) int {
	dp := make([]int, len(height))
	dp[0] = 0
	dp[1] = int(math.Min(float64(height[0]), float64(height[1])))
	for i := 2; i < len(height); i++ {
		tmp := 0
		for j := 0; j < i; j++ {
			if (i-j)*int(math.Min(float64(height[i]), float64(height[j]))) > tmp {
				tmp = (i - j) * int(math.Min(float64(height[i]), float64(height[j])))
			}
		}
		dp[i] = int(math.Max(float64(tmp), float64(dp[i-1])))
	}
	return dp[len(height)-1]
}

func maxArea2(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1
	for left < right {
		if (right-left)*int(math.Min(float64(height[left]), float64(height[right]))) > max {
			max = (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}
