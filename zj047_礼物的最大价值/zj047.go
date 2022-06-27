package zj047_礼物的最大价值

import "fmt"

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	fmt.Printf("m=%d;n=%d\n", m, n)
	dp := make([][]int, m) // 二维切片，3行
	for i := range dp {
		dp[i] = make([]int, n) // 每一行4列
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 && j > 0 {
				dp[i][j] = (dp[i][j-1] + grid[i][j])
			} else if i > 0 && j == 0 {
				dp[i][j] = (dp[i-1][j] + grid[i][j])
			} else {
				dp[i][j] = max((dp[i-1][j] + grid[i][j]), (dp[i][j-1] + grid[i][j]))
			}
		}
	}
	fmt.Printf("dp=%v\n", dp)
	return dp[m-1][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
