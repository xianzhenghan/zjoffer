package zj045_把数组排成最小的数

import (
	"fmt"
	"strconv"
)

/*
输入: [3,30,34,5,9]
输出: "3033459"
*/
func MinNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	leng := len(nums)
	for i := 0; i < leng; i++ {
		for j := 0; j < leng; j++ {
			bo := !compare(nums[i], nums[j])
			fmt.Printf("numsi = %d; > numsj = %d bool=%t\n", nums[i], nums[j], bo)
			if bo {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	str := ""
	for _, v := range nums {
		str = str + strconv.Itoa(v)
	}
	return str
}

func compare(numA, numB int) bool {
	stra := strconv.Itoa(numA)
	strb := strconv.Itoa(numB)
	xa := stra + strb
	xb := strb + stra
	return xa > xb
}
