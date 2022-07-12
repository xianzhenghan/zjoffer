package zj044_数字序列中某一位的数字

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	if n <= 9 {
		return n
	}
	nums := make([]int, 12+1)
	nums[0] = 0
	nums[1] = 10
	for i := 2; i < len(nums); i++ {
		nums[i] = 9*Pow(10, i-1)*i + nums[i-1]
	}
	countnum := 0
	for i := 2; i < len(nums); i++ {
		if nums[i-1] <= n && n < nums[i] {
			countnum = i
		}
	}
	//n 是个 countnum 位数的 第bitnnum 位
	bitnnum := n - nums[countnum-1]
	//数个数字
	digits := Pow(10, countnum-1) + bitnnum - 1/countnum
	fmt.Printf("digits=%v\n", digits)
	strnum := bitnnum % countnum
	fmt.Printf("strnum=%d\n", strnum)
	str := strconv.Itoa(digits)
	return int(str[strnum] - '0')
}

func Pow(e int, n int) int {
	s := 1
	for i := 0; i < n; i++ {
		s = s * e
	}
	return s
}
