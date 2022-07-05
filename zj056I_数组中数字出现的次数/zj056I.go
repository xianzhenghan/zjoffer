package zj065I_不用加减乘除做加法

import "fmt"

func singleNumbers(nums []int) []int {
	if len(nums) < 3 {
		return nums
	}
	xor := 0
	// 1 xor
	for i := 0; i < len(nums); i++ {
		xor = xor ^ nums[i]
	}
	// get xor ==1
	leftshit := 0
	for leftshit < 32 {
		if (xor>>leftshit)&1 == 1 {
			break
		}
		leftshit++
	}
	fmt.Printf("leftshit=%d\n", leftshit)
	// get num
	numa, numb := 0, 0
	for i := 0; i < len(nums); i++ {
		if (nums[i]>>leftshit)&1 == 1 {
			numa = numa ^ nums[i]
		} else {
			numb = numb ^ nums[i]
		}
	}
	return []int{numa, numb}
}
