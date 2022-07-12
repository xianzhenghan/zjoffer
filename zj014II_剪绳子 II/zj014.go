package zj014II_剪绳子_II

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
