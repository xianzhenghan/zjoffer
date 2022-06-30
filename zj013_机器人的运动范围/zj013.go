package zj013_机器人的运动范围

import "fmt"

func movingCount(m int, n int, k int) int {
	if m == 0 || n == 0 {
		return 0
	}
	flag := make([][]int, m)
	for idx, _ := range flag {
		flag[idx] = make([]int, n)
	}
	fmt.Printf("m=%d,n=%d,k=%d\n", m, n, k)
	fmt.Printf("flag=%v\n", flag)
	return dsp(0, 0, m, n, k, flag)
}

func byteSum(num int) int {
	sum := 0
	temp := num
	for temp > 0 {
		sum = sum + temp%10
		temp = temp / 10
	}
	fmt.Printf("num=%d,sum=%d\n", num, sum)
	return sum
}
func dsp(i, j int, m, n int, k int, flag [][]int) int {
	if i < 0 || i >= m || j < 0 || j >= n || flag[i][j] == 1 || bigk(i, j, k) {
		return 0
	}
	flag[i][j] = 1
	return 1 + dsp(i, j+1, m, n, k, flag) + dsp(i, j-1, m, n, k, flag) +
		dsp(i+1, j, m, n, k, flag) + dsp(i-1, j, m, n, k, flag)
}

func bigk(m, n int, k int) bool {
	return byteSum(m)+byteSum(n) > k
}
