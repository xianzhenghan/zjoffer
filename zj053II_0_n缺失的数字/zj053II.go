package zj053II_0_n缺失的数字

//输入: [0,1,2,3,4,5,6,7,9]
//输出: 8
func MissingNumber(nums []int) int {
	l, r := -1, len(nums)
	for l+1 != r {
		mid := (l + r) / 2
		if isbule(nums, mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return l + 1
}

func isbule(nums []int, index int) bool {
	return nums[index] == index
}
