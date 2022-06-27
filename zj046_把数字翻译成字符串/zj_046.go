package zj046_把数字翻译成字符串

import (
	"fmt"
	"strconv"
)

//给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

func TranslateNum(num int) int {
	sli := getSlice(num)
	fmt.Printf("sli =%v \n", sli)
	if len(sli) == 1 {
		return 1
	}
	length := len(sli)
	dp := make([]int, length)
	dp[0] = 1
	if (sli[0]*10 + sli[1]) > 25 {
		dp[1] = 1
	} else {
		dp[1] = 2
	}
	for i := 2; i < length; i++ {
		digit := sli[i-1]*10 + sli[i]
		fmt.Printf("digit =%v \n", digit)
		if digit > 25 || digit < 10 {
			dp[i] = dp[i-1]
		}
		if digit <= 25 && digit > 9 {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	fmt.Printf("dp =%v \n", dp)
	return dp[length-1]
}

func getSlice(num int) []int {
	if num == 0 {
		return []int{0}
	}
	sli := make([]int, 0)
	tmp := num
	for tmp > 0 {
		sli = append(sli, tmp%10)
		tmp = tmp / 10
	}
	left := 0
	right := len(sli) - 1
	for left < right {
		sli[left], sli[right] = sli[right], sli[left]
		left++
		right--
	}
	return sli
}

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	if num >= 10 && num <= 25 {
		return 2
	}
	str := strconv.Itoa(num)
	dp := make([]int, len(str))
	dp[0] = 1                               //边界
	if str[:2] >= "10" && str[:2] <= "25" { //边界
		dp[1] = 2
	} else {
		dp[1] = 1
	}
	for i := 2; i < len(str); i++ {
		newnum := str[i-1 : i+1]
		if newnum >= "10" && newnum <= "25" {
			dp[i] = dp[i-1] + dp[i-2] //递归关系式
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(str)-1]
}
