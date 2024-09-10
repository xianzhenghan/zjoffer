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

func fib2(n int) int {
	if n < 2 {
		return n
	}
	prepre, pre := 0, 1
	nums := 2
	for nums <= n {
		prepre, pre = pre, pre+prepre
	}
	return pre
}
