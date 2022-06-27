package zj063_股票的最大利润

//输入: [7,1,5,3,6,4]
//输出: 5
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minprice := 1 << 32
	maxprofit := 0
	for idx := 0; idx < len(prices); idx++ {
		if prices[idx] < minprice {
			minprice = prices[idx]
		} else {
			if prices[idx]-minprice > maxprofit {
				maxprofit = prices[idx] - minprice
			}
		}
	}
	return maxprofit
}

func maxProfit2(prices []int) int {
	// 时间复杂度O(n) 空间复杂度O(1)
	var min, max, maxProfit = 0x7fffffff, 0, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] > max {
			max = prices[i]
		}
		if prices[i] < min {
			//买入必须在买出前
			min = prices[i]
			max = 0
		}
		if max-min > maxProfit {
			maxProfit = max - min
		}
	}
	return maxProfit
}
