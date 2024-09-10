package zj016_数值的整数次方

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		fn := quick(x, n)
		return 1 / fn
	}
	return quick(x, n)
}

func quick(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quick(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func quickPow(x int64, n int32) int64 {
	if n == 0 {
		return 1
	}
	y := quickPow(x, n/2)
	if n%2 == 0 {
		return y * y
	} else {
		return y * y * x
	}
}

func myPow2(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if x == -1 {
		if n%2 == 0 {
			return 1
		} else {
			return -1
		}
	}

	flag := n > 0
	absn := n
	if n < 0 {
		absn = -absn
	}
	rn := float64(1)
	for i := 0; i < absn; i++ {
		if rn == 0 {
			return 0
		}
		if flag {
			rn = rn * x
		} else {
			rn = (1 / x) * rn
		}
	}
	return rn
}
