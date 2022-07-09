package zj062_圆圈中最后剩下的数字

func lastRemaining(n int, m int) int {
	return rec(n, m)
}

func rec(n, m int) int {
	if n == 1 {
		return 0
	}
	x := rec(n-1, m)
	return (m + x) % n
}
