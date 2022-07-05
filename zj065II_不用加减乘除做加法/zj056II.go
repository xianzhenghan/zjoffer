package zj065II_不用加减乘除做加法

/*输入：nums = [9,1,7,9,7,9,7]
输出：1*/
func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		bit := 1 << i
		countbit1 := 0
		for _, v := range nums {
			if v&bit == 1<<i {
				countbit1++
			}
		}
		if countbit1%3 == 1 {
			res = res | bit
		}
	}
	return res
}
