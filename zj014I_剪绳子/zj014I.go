package zj014I_剪绳子

import "math"

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	y := n % 3
	times := n / 3
	if y == 0 {
		return int(math.Pow(3, float64(times)))
	} else if y == 1 {
		return int(math.Pow(3, float64(times-1))) * 4
	}
	return int(math.Pow(3, float64(times))) * 2
}

func cuttingRope2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	dp := make([]int, n+1)
	for i := 3; i < n+1; i++ {
		for j := 0; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
