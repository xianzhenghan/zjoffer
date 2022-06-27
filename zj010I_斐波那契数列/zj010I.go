package zj010I_斐波那契数列

func fib(n int) int {
	mod := 1000000007
	var sli [101]int
	sli[0] = 0
	sli[1] = 1
	for idx := 2; idx <= n; idx++ {
		sli[idx] = (sli[idx-1] + sli[idx-2]) % mod
	}
	return sli[n]
}
