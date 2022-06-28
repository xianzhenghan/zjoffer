package zj021_调整数组顺序使奇数位于偶数前面

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left]%2 == 1 {
			left++
		} else if nums[right]%2 == 0 {
			right--
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return nums
}
