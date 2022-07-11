package zj049_丑数

import "fmt"

func NthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	index2 := 1
	index3 := 1
	index5 := 1
	for i := 2; i < n+1; i++ {
		dp[i] = min(dp[index2]*2, dp[index3]*3, dp[index5]*5)
		if dp[i] == dp[index2]*2 {
			index2++
		}
		if dp[i] == dp[index3]*3 {
			index3++
		}
		if dp[i] == dp[index5] {
			index5++
		}
	}
	fmt.Printf("dp=%v", dp)
	return dp[n]
}

func min(a, b, c int) int {
	tmp := a
	if a > b {
		tmp = b
	}
	if tmp > c {
		tmp = c
	}
	return tmp
}
func NthUglyNumber2(n int) int {
	dp := make([]int, 0)
	m := map[int]int{}
	for i := 1; i < 7; i++ {
		dp = append(dp, i)
		m[i] = i - 1
	}
	if n < 7 {
		return dp[n-1]
	}
	n1, n2, n3 := 2, 3, 5

	fmt.Printf("dp=%v", dp)
	for i := 7; i <= n; i++ {
		for j := dp[i-2] + 1; j <= n1*(dp[i-2]+1); j++ {
			if j%n1 == 0 && m[j/n1] > 0 {
				dp = append(dp, j)
				m[j] = i
				break
			} else if j%n2 == 0 && m[j/n2] > 0 {
				dp = append(dp, j)
				m[j] = i
				break
			} else if j%n3 == 0 && m[j/n3] > 0 {
				dp = append(dp, j)
				m[j] = i
				break
			}
		}
	}
	fmt.Printf("dp=%v", dp)
	return dp[n-1]
}

func binarySearch(dp []int, num int) bool {
	left := 0
	right := len(dp) - 1
	for left <= right {
		mid := (left + right) >> 1
		if dp[mid] == num {
			return true
		} else if dp[mid] > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
