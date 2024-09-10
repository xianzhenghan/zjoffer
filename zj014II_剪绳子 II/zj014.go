package zj014II_剪绳子_II

import "math"

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	y := n % 3
	times := n / 3
	if y == 0 {
		return Pow(3, times)
	} else if y == 1 {
		return (Pow(3, times-1) * 4) % 1000000007
	}
	return (Pow(3, times) * 2) % 1000000007
}

// 2 <= bamboo_len <= 58
func cuttingBamboo(bamboo_len int) int {
	if bamboo_len == 0 {
		return 0
	}
	if bamboo_len == 1 || bamboo_len == 2 {
		return 1
	}
	dp := make([]int, bamboo_len+1)
	for i := 3; i < bamboo_len+1; i++ {
		for j := 0; j < i; j++ {
			dp[i] = int(math.Max(float64(dp[i]), math.Max(float64(j*(i-j)), float64(j*dp[i-j]))))
		}
	}
	return dp[bamboo_len]
}

func Pow(e int, n int) int {
	s := 1
	max := 1000000007
	for i := 0; i < n; i++ {
		s = s * e
		if s >= max {
			s = s % max
		}
	}
	return s
}
