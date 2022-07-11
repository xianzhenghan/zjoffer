package zj060_n个骰子的点数

func dicesProbability(n int) []float64 {
	dp := make([][]float64, n+1)
	dp[1] = []float64{1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0}

	for i := 2; i < n+1; i++ {
		sli := make([]float64, (i*6 - i + 1))
		for j := 0; j < len(dp[i-1]); j++ {
			for m := 0; m < 6; m++ {
				sli[j+m] = sli[j+m] + float64(dp[i-1][j])*float64(dp[1][m])
			}
		}
		dp[i] = sli
	}
	return dp[n]
}
